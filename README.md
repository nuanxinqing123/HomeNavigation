## 个人导航页
程序的起因是因为家里申请了公网IP，因此购置了一套服务器在家里放着。但是方便操作和记忆，因为写了此导航来可以快速的从内网或者外网访问家里的某些服务

### config.json 文件
```json
{
  "Config": {
    # 网站LOGO（支持相对&绝对路径、支持图床外链）
    "Logo": "img/logo.png",
    # 网站ico（支持相对&绝对路径、支持图床外链）
    "Favicon": "img/favicon.ico",
    # 网站标题
    "Title": "Cloud Services",
    # 网站底部标题
    "FooterTitle": "陕ICP备xxxxxxxx号-1",
    # 网站底部标题（可空）
    "FooterTitleLink": "https://beian.miit.gov.cn/",
    # 网站底部版权信息
    "FooterTextDataOne": "Copyright © 2021",
    # 网站底部其他文字（可随意自定义）
    "FooterTextDataTwo": "百度",
    # 网站底部其他文字连接（可空）
    "FooterTextDataTwoLink": "https://www.baidu.com/"
  },
  "SoftWare": {
    # 程序运行端口（修改后重启生效）      
    "port": "8100"
  },
  "Data": [
    {
      # 显示顺序（不可以重复）   
      "id": 1,
      # 图标标题
      "title": "我的博客",
      # 图标图片（支持相对&绝对路径、支持图床外链）
      "ico": "img/png/AppStore.png",
      # 外网访问地址（可空）
      "w_link": "https://6b7.org/",
      # 内网访问地址（可空）  
      "n_link": "https://6b7.org/"
    },
    {
      "id": 2,
      "title": "我的网盘",
      "ico": "img/png/iCloud-Drive.png",
      "w_link": "https://alist.6b7.org/",
      "n_link": "https://alist.6b7.org/"
    }
  ]
}
```

### 添加新的站点
```json
    {
      "id": ,
      "title": "",
      "ico": "",
      "w_link": "",
      "n_link": ""
    }
```
如果需要添加新的站点到页面，请在Data的 `[ ]` 里面按照如上格式添加。

**Tips：所有的 `{ }` 后面都需要添加一个英文逗号，除了最后一个 `{ }`**

### 配置文件说明
以上配置文件是整个程序的所有配置项，由于程序采用JSON文件实时读取的方式，所以修改配置过后均可实时生效（除了端口需要重启生效）。


### 前端模板说明
模板经过修改，原版模板来自于：http://www.wdmomo.fun:81/home/