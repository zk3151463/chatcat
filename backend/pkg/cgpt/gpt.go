package cgpt

import (
	"chatcat/backend/pkg/clog"
	"chatcat/backend/service"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
)

/*
Model endpoint compatibility
ENDPOINT					MODEL NAME
/v1/chat/completions		gpt-4, gpt-4-0314, gpt-4-32k, gpt-4-32k-0314, gpt-3.5-turbo, gpt-3.5-turbo-0301
/v1/completions				text-davinci-003, text-davinci-002, text-curie-001, text-babbage-001, text-ada-001, davinci, curie, babbage, ada
/v1/edits					text-davinci-edit-001, code-davinci-edit-001
/v1/audio/transcriptions	whisper-1
/v1/audio/translations		whisper-1
/v1/fine-tunes				davinci, curie, babbage, ada
/v1/embeddings				text-embedding-ada-002, text-search-ada-doc-001
/v1/moderations				text-moderation-stable, text-moderation-latest
*/
var (
	once             sync.Once
	gpt              *GPT
	Model                           = openai.GPT3Dot5Turbo
	Messages                        = make([]openai.ChatCompletionMessage, 0) // messages
	Temperature      float32        = 0.7                                     // 适用于/completions和/chat/completions接口 介于 0 和 2 之间。较高的值（如 0.8）将使输出更加随机，而较低的值（如 0.2）将使输出更加集中和确定。较高的temperature将使模型更有可能生成新颖、独特的文本，建议更改这个值或者top_p，两者不能同时改。
	TopP             float32        = 0.1                                     // 适用于/completions和/chat/completions接口  核采样，op_p参数用于控制模型生成文本时，选择下一个单词的概率分布的范围。其中模型考虑具有 top_p 概率质量的标记的结果。所以 0.1 意味着只考虑构成前 10% 概率质量的标记。//
	N                               = 1                                       // 适用于/completions和/chat/completions接口  n这个参数控制了API返回的候选文本的数量，即API会生成多少个可能的文本选项供用户选择。 因为这个参数会产生很多完成，它会很快消耗你的令牌配额。请谨慎使用并确保您对 max_tokens 和 stop 进行了合理的设置。
	Stream                          = true                                    // 适用于/completions和/chat/completions接口  是否回流部分进度。如果设置，令牌将在可用时作为纯数据服务器发送事件发送，流由 data: [DONE] 消息终止。
	Stop             []string                                                 // 适用于/completions和/chat/completions接口  API 将停止生成更多令牌的最多 4 个序列。返回的文本将不包含停止序列。
	MaxTokens                       = 20                                      // 适用于/completions和/chat/completions接口  返回最大token数量
	PresencePenalty  float32        = 0.6                                     // 适用于/completions和/chat/completions接口  表示惩罚那些在生成文本中频繁出现的单词或短语。
	FrequencyPenalty float32        = 0.6                                     // 适用于/completions和/chat/completions接口  参数frequency_penalty可以在生成文本时控制模型是否应该生成高频词汇。
	LogitBias        map[string]int                                           // 适用于/completions和/chat/completions接口  适用于/chat/completions接口
	User             string                                                   // 适用于/completions和/chat/completions接口  代表您的最终用户的唯一标识符，可以帮助 OpenAI 监控和检测滥用行为。了解更多 。
	Prompt           any                                                      // 适用于/completions接口
	Suffix           string                                                   // 适用于/completions接口 插入文本完成后出现的后缀。 Todo
	LogProbs         int                                                      // 适用于/completions接口 logprobs参数用于返回每个生成token的概率值（log-softmax）和其对应的token。logprobs 的最大值为 5
	Echo             bool                                                     // 适用于/completions接口 除了完成之外回显提示
	BestOf           int                                                      // 适用于/completions接口 在服务器端生成 best_of 完成并返回“最佳”（每个标记具有最高对数概率的那个）。无法流式传输结果。 因为这个参数会产生很多完成，它会很快消耗你的令牌配额。请谨慎使用并确保您对 max_tokens 和 stop 进行了合理的设置。
)

// GPT top_p适用于需要控制模型生成文本多样性的场景，可以让生成的文本保持一定的可控性，适用于需要保留一些局部结构的场景。而temperature适用于需要生成更加多样、创新的文本，可以让生成的文本更具随机性，适用于需要在新领域进行探索的场景。
type GPT struct {
	App                   *service.App
	Client                *openai.Client
	ChatCompletionRequest openai.ChatCompletionRequest
	CompletionRequest     openai.CompletionRequest
	Token                 string // apiKey
	TikToken              int    // 消耗的token数量
	Done                  chan bool
	ChatCompStream        *openai.ChatCompletionStream
	err                   error
}

func New(token string, app *service.App) *GPT {
	fmt.Println("token", token)
	once.Do(func() {
		gpt = &GPT{
			App:    app,
			Client: openai.NewClient(token),
			Token:  token,
			Done:   make(chan bool),
		}
	})
	return gpt
}

func (g *GPT) WithModel(model string) *GPT {
	Model = model
	return g
}

func (g *GPT) WithMessages(messages []openai.ChatCompletionMessage) *GPT {
	Messages = messages
	return g
}

func (g *GPT) WithPrompt(prompt string) *GPT {
	Prompt = prompt
	return g
}

func (g *GPT) WithStream(stream bool) *GPT {
	Stream = stream
	return g
}

func (g *GPT) WithTemperature(temperature float32) *GPT {
	Temperature = temperature
	return g
}

func (g *GPT) WithPresencePenalty(presencePenalty float32) *GPT {
	PresencePenalty = presencePenalty
	return g
}
func (g *GPT) WithFrequencyPenalty(frequencyPenalty float32) *GPT {
	PresencePenalty = frequencyPenalty
	return g
}
func (g *GPT) WithN(n int) *GPT {
	N = n
	return g
}

func (g *GPT) Error() error {
	return g.err
}

func (g *GPT) WithMaxTokens(tokens int) *GPT {
	var tikToken int
	var err error
	if tokens == 0 {
		modelTokens := getMaxTokensByModel()
		prompt, _ := Prompt.(string)
		if prompt == "" {
			tikToken, err = g.NumTokensFromMessages(Messages)
			if err != nil {
				g.err = errors.New("Chatcat Warm Reminder1: Token calculation failed." + err.Error())
			}
		} else {
			tikToken, err = g.getTikTokenByEncoding(prompt)
			if err != nil {
				g.err = errors.New("Chatcat Warm Reminder: Token calculation failed." + err.Error())
			}
		}
		g.TikToken = tikToken
		MaxTokens = modelTokens - tikToken
		g.App.LogInfof("tikToken: %d modelTokens: %d leftToken: %d model: %s message count: %v", tikToken, modelTokens, MaxTokens, Model, len(Messages))
	} else {
		MaxTokens = tokens
		g.TikToken = tokens
	}
	return g
}

// WithProxy
// @Description: with proxy
// @receiver g
// @param rawUrl http://localhost:{port}
// @return *GPT
// @author cx
func (g *GPT) WithProxy(rawUrl string) *GPT {
	config := openai.DefaultConfig(g.Token)
	// proxy := http.ProxyFromEnvironment
	// if rawUrl != "" {
	// 	proxyUrl, err := url.Parse(rawUrl)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	proxy = http.ProxyURL(proxyUrl)
	// }
	// config.HTTPClient = &http.Client{
	// 	Transport: &http.Transport{
	// 		Proxy: proxy,
	// 	},
	// 	Timeout: time.Second * 30,
	// }
	config.BaseURL = rawUrl
	g.Client = openai.NewClientWithConfig(config)
	return g
}

// WithChatCompletionRequest
// @Description: with chat completion request
// @receiver g
// @return *GPT
// @author cx
func (g *GPT) WithChatCompletionRequest() *GPT {
	g.ChatCompletionRequest = openai.ChatCompletionRequest{
		Model:            Model,
		Messages:         Messages,
		MaxTokens:        MaxTokens,
		Temperature:      Temperature,
		TopP:             TopP,
		N:                N,
		Stream:           Stream,
		Stop:             Stop,
		PresencePenalty:  PresencePenalty,
		FrequencyPenalty: FrequencyPenalty,
		LogitBias:        LogitBias,
		User:             User,
	}
	return g
}

// WithCompletionRequest 是 GPT 结构体的一个方法，用于设置并返回一个包含完成请求的 GPT 实例。
func (g *GPT) WithCompletionRequest() *GPT {
	// 设置 g 的 CompletionRequest 字段，该字段是一个 openai.CompletionRequest 结构体。
	// 该结构体包含了生成文本所需的多个参数。
	g.CompletionRequest = openai.CompletionRequest{
		// Model 指定了使用的模型，例如 "text-davinci-003"。
		Model: Model,
		// Prompt 是生成文本的输入提示。
		Prompt: Prompt,
		// Suffix 是生成文本时添加的后缀。
		Suffix: Suffix,
		// MaxTokens 指定了生成的最大令牌数。
		MaxTokens: MaxTokens,
		// Temperature 控制生成文本的多样性，值越大，生成的文本越随机。
		Temperature: Temperature,
		// TopP 控制生成的多样性，值越大，生成的文本越随机。
		TopP:   TopP,
		N:      N,
		Stream: Stream,
		// Stream 指定是否以流的形式生成文本。
		LogProbs: LogProbs,
		// LogProbs 指定返回每个生成的令牌的对数概率。
		Echo: Echo,
		// Echo 指定是否在返回的文本中包含输入提示。
		Stop: Stop,
		// Stop 是一个字符串列表，指定生成文本时停止的标记。
		PresencePenalty: PresencePenalty,
		// PresencePenalty 控制生成文本时对重复内容的惩罚。
		FrequencyPenalty: FrequencyPenalty,
		// FrequencyPenalty 控制生成文本时对频率高的内容的惩罚。
		BestOf: BestOf,
		// BestOf 指定从多个生成的文本中选择最佳的一个。
		LogitBias: LogitBias,
		// LogitBias 是一个映射，用于控制生成文本时每个令牌的权重。
		User: User,
		// User 是一个字符串，用于标识请求的用户。
	}
	return g
	// 返回设置好的 GPT 实例。
}

// ChatCompletionStream
// @Description: ChatGPT streaming completion https://api.openai.com/v1/chat/completions
// @receiver g
// @param prompt
// @author cx
func (g *GPT) ChatCompletionStream(clientId string) {
	var err error
	if Stream == false {
		panic("ChatCompletionStream should be set stream")
	}
	g.App.ClientId = clientId
	g.ChatCompStream, err = g.Client.CreateChatCompletionStream(g.App.Ctx, g.ChatCompletionRequest)
	if err != nil {
		g.App.WsPushChan <- service.PushResp{
			Code: -1,
			Data: fmt.Sprintf("Chatcat Warm Reminder12: %s", err.Error()),
		}
		return
	}
	defer g.ChatCompStream.Close()
	for {
		select {
		default:
			response, err := g.ChatCompStream.Recv()
			if errors.Is(err, io.EOF) {
				g.App.WsPushChan <- service.PushResp{
					Code: 1,
					Data: "Chatcat Finished",
				}
				return
			}
			if err != nil {
				g.App.WsPushChan <- service.PushResp{
					Code: -1,
					Data: fmt.Sprintf("Chatcat Warm Reminder2: %s", err.Error()),
				}
				return
			}
			// ws 推送
			//fmt.Printf(response.Choices[0].Delta.Content)
			//clog.PrintInfo(fmt.Sprintf("ChatCompletionStream text: %#v", response.Choices[0].Delta.Content))
			g.App.WsPushChan <- service.PushResp{
				Code: 0,
				Data: response.Choices[0].Delta.Content,
			}
		}
	}
}

// CompletionStream
// @Description: gpt3 completion stream https://api.openai.com/v1/completions
// @receiver g
// @param prompt
// @author cx
func (g *GPT) CompletionStream(clientId string) {
	if Stream == false {
		panic("CompletionStream should be set stream")
	}
	g.App.ClientId = clientId
	stream, err := g.Client.CreateCompletionStream(g.App.Ctx, g.CompletionRequest)
	if err != nil {
		g.App.WsPushChan <- service.PushResp{
			Code: -1,
			Data: fmt.Sprintf("Chatcat Warm Reminder: %s", err.Error()),
		}
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			g.App.WsPushChan <- service.PushResp{
				Code: -1,
				Data: fmt.Sprintf("Chatcat Warm Reminder: %s", err.Error()),
			}
			return
		}
		// ws push
		clog.PrintInfo(fmt.Sprintf("CompletionStream text: %#v", response.Choices[0].Text))
		g.App.WsPushChan <- service.PushResp{
			Code: 0,
			Data: response.Choices[0].Text,
		}
	}
	return
}

// ChatCompletionNoStream
// @Description: chat completion no streaming https://api.openai.com/v1/chat/completions
// @receiver g
// @param prompt
// @author cx
func (g *GPT) ChatCompletionNoStream() (*openai.ChatCompletionResponse, error) {
	resp, err := g.Client.CreateChatCompletion(g.App.Ctx, g.ChatCompletionRequest)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CompletionNoStream
// @Description: https://api.openai.com/v1/completions 为提供的提示和参数创建完成
// @receiver g
// @param prompt
// @author cx
func (g *GPT) CompletionNoStream() (openai.CompletionResponse, error) {
	resp, err := g.Client.CreateCompletion(g.App.Ctx, g.CompletionRequest)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ListModels
// @Description: get models
// @receiver g
// @return openai.ModelsList
// @return error
// @author cx
func (g *GPT) ListModels() (openai.ModelsList, error) {
	modelList, err := g.Client.ListModels(g.App.Ctx)
	if err != nil {
		return modelList, err
	}
	return modelList, nil
}

// getMaxTokensByModel
// @Description: https://platform.openai.com/docs/models/gpt-3-5
// @receiver g
// @param model
// @return int
// @author cx
func getMaxTokensByModel() int {
	tokens := map[string]int{
		openai.GPT4:                8192,  // 比任何 GPT-3.5 模型都更强大，能够执行更复杂的任务，并针对聊天进行了优化。将使用我们最新的模型迭代进行更新。截至 2021 年 9 月
		openai.GPT40314:            8192,  // gpt-4 2023 年 3 月 14 日的快照。与 gpt-4 不同，此模型不会收到更新，并且仅在 2023 年 6 月 14 日结束的三个月内提供支持。截至 2021 年 9 月
		openai.GPT432K:             32768, // 与基本 gpt-4 模式相同的功能，但上下文长度是其 4 倍。将使用我们最新的模型迭代进行更新。截至 2021 年 9 月
		openai.GPT432K0314:         32768, // gpt-4-32 2023 年 3 月 14 日的快照。与 gpt-4-32k 不同，此模型不会收到更新，并且仅在 2023 年 6 月 14 日结束的三个月内提供支持。截至 2021 年 9 月
		openai.GPT3Dot5Turbo:       4096,  // 功能最强大的 GPT-3.5 模型并针对聊天进行了优化，成本仅为 text-davinci-003 的 1/10。将使用我们最新的模型迭代进行更新。截至 2021 年 9 月
		openai.GPT3Dot5Turbo0301:   4096,  // gpt-3.5-turbo 的快照，2023 年 3 月 1 日。与 gpt-3.5-turbo 不同，此模型不会收到更新，并且仅在 2023 年 6 月 1 日结束的三个月内提供支持。截至 2021 年 9 月
		openai.GPT3TextDavinci003:  4097,  // 可以以比居里、巴贝奇或 ada 模型更好的质量、更长的输出和一致的指令遵循来完成任何语言任务。还支持在文本中插入补全。 截至 2021 年 6 月
		openai.GPT3TextDavinci002:  4097,  // 与 text-davinci-003 类似的功能，但使用监督微调而不是强化学习进行训练 截至 2021 年 6 月
		openai.CodexCodeDavinci002: 8001,  // 针对代码完成任务进行了优化 截至 2021 年 6 月
		openai.GPT3TextCurie001:    2049,  // 非常有能力，比达芬奇更快，成本更低。截至 2019 年 10 月
		openai.GPT3TextBabbage001:  2049,  // 能够执行简单的任务，速度非常快，成本更低。截至 2019 年 10 月
		openai.GPT3TextAda001:      2049,  // 能够执行非常简单的任务，通常是 GPT-3 系列中最快的型号，而且成本最低。截至 2019 年 10 月
		openai.GPT3Davinci:         2049,  // 功能最强大的 GPT-3 模型。可以完成其他模型可以完成的任何任务，而且通常质量更高。截至 2019 年 10 月
		openai.GPT3Curie:           2049,  // 非常有能力，但比达芬奇更快，成本更低。截至 2019 年 10 月
		openai.GPT3Babbage:         2049,  // 能够执行简单的任务，速度非常快，成本更低。截至 2019 年 10 月
		openai.GPT3Ada:             2049,  // 能够执行非常简单的任务，通常是 GPT-3 系列中最快的型号，而且成本最低。截至 2019 年 10 月
	}
	return tokens[Model]
}

// getTikTokenByEncoding
// @Description: get token by model Todo
// @receiver g
// @param prompt
// @return int
// @return error
// @author cx
func (g *GPT) getTikTokenByEncoding(prompt string) (int, error) {
	encoding := g.getAvailableEncodingModel(Model)
	tkm, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		return 0, err
	}
	token := tkm.Encode(prompt, nil, nil)
	return len(token), nil
}

// NumTokensFromMessages
// @Description: get number of tokens from messages and model
// @receiver g
// @param messages
// @param model
// @return numTokens
// @author cx
func (g *GPT) NumTokensFromMessages(messages []openai.ChatCompletionMessage) (numTokens int, err error) {
	// encoding := g.getAvailableEncodingModel(Model)
	// tkm, err := tiktoken.GetEncoding(encoding)
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("EncodingForModel: %v", err))
	// 	return
	// }
	// var tokensPerMessage int
	// var tokensPerName int
	// if Model == openai.GPT3Dot5Turbo0301 || Model == openai.GPT3Dot5Turbo {
	// 	tokensPerMessage = 4
	// 	tokensPerName = -1
	// } else if Model == openai.GPT40314 || Model == openai.GPT4 {
	// 	tokensPerMessage = 3
	// 	tokensPerName = 1
	// } else {
	// 	fmt.Println("Warning: model not found. Using cl100k_base encoding.")
	// 	tokensPerMessage = 3
	// 	tokensPerName = 1
	// }
	// for _, message := range messages {
	// 	numTokens += tokensPerMessage
	// 	numTokens += len(tkm.Encode(message.Content, nil, nil))
	// 	numTokens += len(tkm.Encode(message.Role, nil, nil))
	// 	if message.Name != "" {
	// 		numTokens += tokensPerName
	// 	}
	// }
	numTokens += 3
	return numTokens, nil
}

// getAvailableEncodingModel
// @Description: get available encoding model
// @receiver g
// @param model
// @return string
// @author cx
func (g *GPT) getAvailableEncodingModel(model string) string {
	encodeMaps := map[string]string{
		openai.GPT4:                "cl100k_base",
		openai.GPT40314:            "cl100k_base",
		openai.GPT432K:             "cl100k_base",
		openai.GPT432K0314:         "cl100k_base",
		openai.GPT3Dot5Turbo:       "cl100k_base",
		openai.GPT3Dot5Turbo0301:   "cl100k_base",
		openai.GPT3TextDavinci003:  "p50k_base",
		openai.GPT3TextDavinci002:  "p50k_base",
		openai.CodexCodeDavinci002: "p50k_base",
		openai.GPT3TextCurie001:    "r50k_base",
		openai.GPT3TextBabbage001:  "r50k_base",
		openai.GPT3TextAda001:      "r50k_base",
		openai.GPT3Davinci:         "r50k_base",
		openai.GPT3Curie:           "r50k_base",
		openai.GPT3Babbage:         "r50k_base",
		openai.GPT3Ada:             "r50k_base",
	}
	return encodeMaps[model]
}
