export namespace cresp {
	
	export class Response {
	    code: number;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}

}

export namespace model {
	
	export class Chat {
	    id: number;
	    created_at: number;
	    updated_at: number;
	    cate_id: number;
	    name: string;
	    sort: number;
	
	    static createFrom(source: any = {}) {
	        return new Chat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	        this.cate_id = source["cate_id"];
	        this.name = source["name"];
	        this.sort = source["sort"];
	    }
	}
	export class ChatCate {
	    id: number;
	    created_at: number;
	    updated_at: number;
	    name: string;
	    letter: string;
	    color: string;
	    desc: string;
	
	    static createFrom(source: any = {}) {
	        return new ChatCate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	        this.name = source["name"];
	        this.letter = source["letter"];
	        this.color = source["color"];
	        this.desc = source["desc"];
	    }
	}
	export class ChatRecord {
	    id: number;
	    created_at: number;
	    updated_at: number;
	    cate_id: number;
	    chat_id: number;
	    name: string;
	    content: string;
	    reply_content: string;
	    prefix: string;
	    role: string;
	
	    static createFrom(source: any = {}) {
	        return new ChatRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	        this.cate_id = source["cate_id"];
	        this.chat_id = source["chat_id"];
	        this.name = source["name"];
	        this.content = source["content"];
	        this.reply_content = source["reply_content"];
	        this.prefix = source["prefix"];
	        this.role = source["role"];
	    }
	}
	export class FeedBack {
	    title?: string;
	    body?: string;
	    labels?: string[];
	    assignee?: string;
	    state?: string;
	    state_reason?: string;
	    milestone?: number;
	    assignees?: string[];
	
	    static createFrom(source: any = {}) {
	        return new FeedBack(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.body = source["body"];
	        this.labels = source["labels"];
	        this.assignee = source["assignee"];
	        this.state = source["state"];
	        this.state_reason = source["state_reason"];
	        this.milestone = source["milestone"];
	        this.assignees = source["assignees"];
	    }
	}
	export class FeedbackReq {
	    issue_type?: number;
	    title?: string;
	    body?: string;
	    version?: string;
	
	    static createFrom(source: any = {}) {
	        return new FeedbackReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.issue_type = source["issue_type"];
	        this.title = source["title"];
	        this.body = source["body"];
	        this.version = source["version"];
	    }
	}
	export class Prompt {
	    id: number;
	    created_at: number;
	    updated_at: number;
	    name: string;
	    desc: string;
	    prompt: string;
	    enprompt: string;
	    tag_name: string;
	    is_common: number;
	
	    static createFrom(source: any = {}) {
	        return new Prompt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	        this.name = source["name"];
	        this.desc = source["desc"];
	        this.prompt = source["prompt"];
	        this.enprompt = source["enprompt"];
	        this.tag_name = source["tag_name"];
	        this.is_common = source["is_common"];
	    }
	}
	export class Setting {
	    id: number;
	    created_at: number;
	    updated_at: number;
	    api_key: string;
	    chat_platform: string;
	    is_customize: number;
	    chat_model: string;
	    ask_model: string;
	    language: string;
	    theme: number;
	    proxy_url: string;
	    account: string;
	    access_token: string;
	    is_sync: number;
	    sync_time: number;
	    temperature: string;
	    max_tokens: number;
	    presence_penalty: string;
	    frequency_penalty: string;
	    n: number;
	    top_p: string;
	
	    static createFrom(source: any = {}) {
	        return new Setting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	        this.api_key = source["api_key"];
	        this.chat_platform = source["chat_platform"];
	        this.is_customize = source["is_customize"];
	        this.chat_model = source["chat_model"];
	        this.ask_model = source["ask_model"];
	        this.language = source["language"];
	        this.theme = source["theme"];
	        this.proxy_url = source["proxy_url"];
	        this.account = source["account"];
	        this.access_token = source["access_token"];
	        this.is_sync = source["is_sync"];
	        this.sync_time = source["sync_time"];
	        this.temperature = source["temperature"];
	        this.max_tokens = source["max_tokens"];
	        this.presence_penalty = source["presence_penalty"];
	        this.frequency_penalty = source["frequency_penalty"];
	        this.n = source["n"];
	        this.top_p = source["top_p"];
	    }
	}

}

export namespace openai {
	
	export class ChatCompletionMessage {
	    role: string;
	    content: string;
	    name?: string;
	
	    static createFrom(source: any = {}) {
	        return new ChatCompletionMessage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.role = source["role"];
	        this.content = source["content"];
	        this.name = source["name"];
	    }
	}

}

