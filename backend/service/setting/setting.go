package setting

import (
	"chatcat/backend/model"
	"chatcat/backend/pkg/chttp"
	"chatcat/backend/pkg/cresp"
	"chatcat/backend/service"
	"fmt"
	"github.com/google/go-github/v51/github"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net/url"
	rt "runtime"
)

type Service struct {
	App *service.App
}

func New(a *service.App) *Service {
	return &Service{
		App: a,
	}
}

// GetGeneralInfo
// @Description: get general info
// @receiver s
// @return *cresp.Response
// @author cx
func (s *Service) GetGeneralInfo() *cresp.Response {
	var settingInfo model.Setting
	if err := s.App.DB.First(&settingInfo).Error; err != nil {
		return cresp.Fail(err.Error())
	}
	return cresp.Success(settingInfo)
}

// SetGeneralData
// @Description: set general data
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) SetGeneralData(data model.Setting) *cresp.Response {
	if err := s.App.DB.Model(&model.Setting{}).Where("id = 1").Updates(&data).Error; err != nil {
		return cresp.Fail("save failed err:" + err.Error())
	}
	return cresp.Success(data)
}

// FeedBack
// @Description: feed back https://github.com/MQEnergy/chatcat/issues/new?title=&body=
// https://github.com/arco-design/arco-design-vue/issues/new?
// title=vue%E4%B8%BB%E9%A2%98%E8%AE%BE%E7%BD%AE%E4%B8%AD%20%E8%B7%9F%E9%9A%8F%E7%B3%BB%E7%BB%9F%E6%8A%A5%E9%94%99&
// body=-%20%5B%20%5D%20I%27m%20sure%20this%20does%20not%20appear%20in%20%5Bthe%20issue%20list%20of%20the%20repository%5D(https%3A%2F%2Fgithub.com%2Farco-design%2Farco-design-vue%2Fissues)%0A%0A%23%23%20Basic%20Info%20%0A-%20**Package%20Name%20And%20Version%3A**%20%40arco-design%2Fweb-vue%402.45.1%0A-%20**Browser%3A**%20chrome112.0.0.0%0A%0A%23%23%20Steps%20to%20reproduce%0A123123%0A%0A%3C!---%20Disclaimer%3A%20Submitting%20offensive%20issues%20will%20result%20in%20being%20blocked%20from%20arco-design%20organization.%20--%3E%0A%3C!--%20generated%20by%20arco-issue.%20DO%20NOT%20REMOVE%20--%3E
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) FeedBack(data model.FeedBack) *cresp.Response {
	client := github.NewClient(nil)
	orgs, resp, err := client.Issues.Create(s.App.Ctx, s.App.Cfg.Github.Owner, s.App.Cfg.Github.Repo, &github.IssueRequest{
		Title:       &data.Title,
		Body:        &data.Body,
		Labels:      &data.Labels,
		Assignees:   &data.Assignees,
		State:       &data.State,
		StateReason: &data.StateReason,
		Milestone:   &data.Milestone,
	})
	fmt.Println(orgs, resp, err)
	if err != nil {
		return cresp.Fail("feedback failed err:" + err.Error())
	}
	return cresp.Success("")
}

// GetFeedBackUrl
// @Description: get feed back url
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
// 根据传入的FeedbackReq数据，生成一个反馈的URL，并返回一个Response对象
func (s *Service) GetFeedBackUrl(data model.FeedbackReq) *cresp.Response {
	// 初始化一个字符串，用于存储反馈的内容
	body := "- [ ] I'm sure this does not appear in [the issue list of the repository](https://github.com/MQEnergy/chatcat/issues) "
	// 判断反馈的类型
	if data.IssueType == 1 {
		// 如果是问题反馈，则将反馈的内容格式化为字符串
		body += fmt.Sprintf("%s ## Basic Info:%s - Version: %s ## Steps to reproduce: %s", "%0A", "%0A", data.Version+"%0A", "%0A"+data.Body+"%0A")
	} else {
		// 如果是建议反馈，则将反馈的内容格式化为字符串
		body += fmt.Sprintf("%s ## Basic Info:%s - Version: %s ## What is expected?: %s", "%0A", "%0A", data.Version+"%0A", "%0A"+data.Body+"%0A")
	}
	// 解析生成的URL
	parseUrl, _ := url.Parse("https://github.com/" + s.App.Cfg.Github.Owner + "/" + s.App.Cfg.Github.Repo + "/issues/new?title=" + data.Title + "&body=" + body)
	// 返回一个Response对象，其中包含生成的URL
	return cresp.Success(parseUrl.String())
}

// GetGithubReleaseList
// @Description: get GitHub release list
// @receiver s
// @author cx
func (s *Service) GetGithubReleaseList() *cresp.Response {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(s.App.Ctx, s.App.Cfg.Github.Owner, s.App.Cfg.Github.Repo, nil)
	if err != nil {
		return cresp.Fail(err.Error())
	}
	return cresp.Success(releases)
}

// Minimise ...
func (s *Service) Minimise() {
	runtime.WindowMinimise(s.App.Ctx)
}

// Fullscreen ...
func (s *Service) Fullscreen() {
	runtime.WindowFullscreen(s.App.Ctx)
}

// NormalScreen ...
func (s *Service) NormalScreen() {
	runtime.WindowSetSize(s.App.Ctx, s.App.Cfg.App.Width, s.App.Cfg.App.Height)
}

// Hide ...
func (s *Service) Hide() {
	runtime.Hide(s.App.Ctx)
}

// Quit ...
func (s *Service) Quit() {
	runtime.Quit(s.App.Ctx)
}

// ReloadApp ...
func (s *Service) ReloadApp() {
	runtime.WindowReloadApp(s.App.Ctx)
}

// IsWindows ...
func (s *Service) IsWindows() bool {
	if rt.GOOS == "darwin" {
		return false
	}
	return true
}
func (s *Service) GeOLLamaModel(urls string) *cresp.Response {
	var releases, err = chttp.Request("GET", urls, "")
	if err != nil {
		return cresp.Fail(err.Error())
	}
	return cresp.Success(releases)
}
