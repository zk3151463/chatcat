export default {
  // general settings
  'settings.general': 'General',
  'settings.general.keyTips': 'Your API key is stored locally in your browser and will never be sent anywhere else. The extension supports GPT-4 if you input a GPT-4 API key.',
  'settings.general.alertTips': 'Please follow the guidelines to obtain API keys and up to $5 credit for free from OpenAI',
  'settings.general.alertTips.drawer.title': '🐱Welcome to Chatcat! Here is your journey to begin:',
  'settings.general.alertTips.drawer.step1.h1': 'How to get an API key.',
  'settings.general.alertTips.drawer.step1.h2': 'Step 1: Get API key',
  'settings.general.alertTips.drawer.step1.li1-1': 'Login your OpenAI account',
  'settings.general.alertTips.drawer.step1.li1-2': 'If you don\'t have an account, Please',
  'settings.general.alertTips.drawer.step1.li1-3': 'Sign up',
  'settings.general.alertTips.drawer.step1.li1-4': 'HERE',
  'settings.general.alertTips.drawer.step1.li2': 'Visit OpenAI API Key page',
  'settings.general.alertTips.drawer.step1.li3': 'Create a new secret key and copy it as follow',
  'settings.general.alertTips.drawer.step2.h2': 'Step 2: Paste API Key in the box below',
  'settings.general.alertTips.drawer.step2.li1': 'Don\'t worry. Your API key is only used locally within the ChatGPT extension. Using the API is more stable than directly using ChatGPT.',
  'settings.general.alertTips.drawer.step3.h2-1': 'Step 3: Check your credits',
  'settings.general.alertTips.drawer.step3.h2-2': 'here',
  'settings.general.alertTips.drawer.step3.h2-3': 'please',
  'settings.general.apiKey': 'API Key',
  'settings.general.apiPlatform': 'API Platform',
  'settings.general.apiPlatform.placeholder': 'Please select an API platform',
  'settings.general.apiModel': 'API Model',
  'settings.general.apiModel.placeholder': 'Please select an API model',
  'settings.language': 'Language',
  'settings.theme': 'Theme',
  'settings.theme.system': 'Auto',
  'settings.theme.toLight': 'Light mode',
  'settings.theme.toDark': 'Dark mode',
  'settings.screen.toFull': 'Click to switch to full screen mode',
  'settings.screen.toExit': 'Click to exit the full screen mode',
  "settings.proxy": "API Address",
  "settings.proxy.placeholder": "Please input a  address eg. http://localhost:11434/v1",
  // advanced settings
  'settings.advanced': 'Advanced',
  'settings.advanced.temperature': 'Randomness (temperature)',
  'settings.advanced.temperature.desc': 'The larger the value, the more random the reply will be, ranging between 0 and 2. Values greater than 1 may result in garbled text.',
  'settings.advanced.max_tokens': 'Reply Limitation per Session (max_tokens)',
  'settings.advanced.max_tokens.desc': 'The maximum number of tokens allowed for a single interaction is set to 0, which means that the maximum number of tokens allowed by the current model will be automatically calculated.',
  'settings.advanced.presence_penalty': 'Conversation freshness (presence_penalty)',
  'settings.advanced.presence_penalty.desc': 'The weight of the token will be punished based on whether it has appeared in the text so far, in order to increase the possibility of the model discussing new topics. It is between -2.0 and 2.0.',
  'settings.advanced.frequency_penalty': 'Repetitiveness in Dialogue (frequency_penalty)',
  'settings.advanced.frequency_penalty.desc': 'The weight will be penalized based on the existing frequency of the new marker in the text, thus reducing the likelihood of the model repeating the same line word by word. It is between -2.0 and 2.0.',
  'settings.advanced.N': 'Return quantity (N)',
  'settings.advanced.N.desc': 'How many possible text options will the API generate for users to choose from, and it will quickly consume your token quota. Please use it with caution. Default 1.',
  // data sync
  'settings.datasync': 'Data Sync',
  'settings.datasync.linkAccount': 'Linked account',
  'settings.datasync.loginBtn': 'Login',
  'settings.datasync.registerBtn': 'Please go to register',
  'settings.datasync.beginSync': 'Start sync',
  'settings.datasync.accountTips': 'Register a Chatcat account to provide cloud storage for your data. You can backup your data to the cloud and synchronize it across devices.',
  'settings.datasync.syncTips': '5s automatically syncs once',
  // prompt
  'settings.prompt': 'Prompt',
  'settings.prompt.common': 'Common prompt',
  'settings.prompt.cate': 'Category',
  'settings.prompt.cate.name': 'Name:',
  'settings.prompt.cate.name.placeholder': 'Please enter a prompt name, for example: Writing Title Generator.',
  'settings.prompt.cate.name.errTip1': 'Name is required.',
  'settings.prompt.cate.desc': 'Description:',
  'settings.prompt.cate.desc.placeholder': 'Please provide an explanation in the remarks section',
  'settings.prompt.cate.prompt': 'Prompt(zh_CN):',
  'settings.prompt.cate.prompt.placeholder': 'Please input Chinese prompts.',
  'settings.prompt.cate.prompt.errTip1': 'Prompt is required.',
  'settings.prompt.cate.enprompt': 'Prompt(EN):',
  'settings.prompt.cate.enprompt.placeholder': 'Please input English prompts.',
  'settings.prompt.cate.cancel': 'Cancel',
  'settings.prompt.cate.submit': 'Submit',
  'settings.prompt.viewBtn': 'View',
  'settings.prompt.addBtn': 'Add Prompt',
  'settings.prompt.drawerTitle': 'Prompt list',
  // chat management
  'settings.chat': 'Chat',
  'settings.chat.drawerTitle': 'Chat list',
  'settings.chat.catmove': 'Cate move',
  // contact us
  'settings.contact': 'Contact',
  'settings.contact.star': 'Give me a star.',
  'settings.contact.feedbackType': 'Feedback type:',
  'settings.contact.feedbackType.placeholder': 'Please select feedback type',
  'settings.contact.feedbackType.option1': 'Bug Report',
  'settings.contact.feedbackType.option2': 'Feature Request',
  'settings.contact.title': 'Title:',
  'settings.contact.title.tips': 'Please state the issue or requirement in a few sentences',
  'settings.contact.title.placeholder': 'Please enter the title name',
  'settings.contact.title.errorTips1': 'Title name is required',
  'settings.contact.title.errorTips2': 'Cannot be less than five characters',
  'settings.contact.desc': 'Detailed description:',
  'settings.contact.desc.tips': 'Please provide a detailed description of the problem, any actions you have taken, and any error messages received',
  'settings.contact.desc.placeholder': 'Please provide a detailed description',
  'settings.contact.desc.errorTips1': 'Detailed description is required',
  'settings.contact.desc.errorTips2': 'Cannot be less than five characters',
  'settings.contact.email.tips': 'Enter your email address to receive updates about this issue',
  'settings.contact.email.placeholder': 'Please provide a detailed description',
  'settings.contact.email.errorTips1': 'Detailed description is required',
  'settings.contact.email.errorTips2': 'Cannot be less than five characters',
  'settings.contact.cancelBtn': 'Cancel',
  'settings.contact.addBtn': 'Submit',
  // update info
  'settings.releaseNotes': 'Update',
  'settings.updateCate': 'Update information',
  'settings.checkUpdate': 'Check update',
};
