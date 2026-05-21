# wechatbot-go-sdk

基于WeChatBot开放平台API封装的 Go SDK，提供简洁易用的接口调用方式。

## 特性

- 零外部依赖，仅使用 Go 标准库（net/http + encoding/json）
- Functional Options 模式构造客户端，支持自定义服务地址和超时
- 10 个功能模块完整覆盖平台所有 API
- 返回 `(*Response, error)` 惯用 Go 错误处理风格
- 并发安全，可在多个 goroutine 中共享同一客户端实例

## 环境要求

- Go >= 1.21

## 安装

```bash
go get github.com/oocsoo/wechatbot-go-sdk
```

## 快速开始

```go
package main

import (
    "fmt"
    wechatbot "github.com/user/wechatbot-go-sdk"
)

func main() {
    client := wechatbot.NewClient()
    token := "your_token"

    // 获取登录二维码
    resp, err := client.Auth.GetQrcode(token, "", "")
    if err != nil {
        panic(err)
    }
    fmt.Println(resp)

    // 发送文本消息
    resp, err = client.Message.SendText("wxid_xxxxx", "Hello!", token, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(resp)
}
```

## 自定义配置

```go
import "time"

client := wechatbot.NewClient(
    wechatbot.WithBaseURL("http://your-server-ip"),
    wechatbot.WithTimeout(60 * time.Second),
)
```

## 响应处理

所有接口返回 `(*Response, error)`：

```go
resp, err := client.Auth.CheckStatus(token)
if err != nil {
    // 网络错误
    log.Fatal(err)
}
if resp.IsSuccess() {
    // ret == 200 表示成功
    fmt.Println("在线:", string(resp.Data))
} else {
    fmt.Println("错误:", resp.Msg)
}
```

`resp.Data` 是 `json.RawMessage`，可按需反序列化：

```go
var result map[string]any
json.Unmarshal(resp.Data, &result)
```

## 功能模块总览

| 模块 | 属性 | 说明 |
|------|------|------|
| AuthModule | `client.Auth` | 登录认证（扫码/重连/登出/状态检查） |
| MessageModule | `client.Message` | 消息收发（文本/图片/语音/视频/文件/链接/名片/小程序/撤回/转发） |
| GroupModule | `client.Group` | 群聊管理（创建/邀请/踢人/公告/管理员/置顶/免打扰） |
| ContactModule | `client.Contact` | 联系人管理（通讯录/搜索/加删好友/企微） |
| PersonalModule | `client.Personal` | 个人信息（资料/二维码/隐私设置/修改头像） |
| FavoriteModule | `client.Favorite` | 收藏夹（同步/获取/删除） |
| LabelModule | `client.Label` | 标签管理（增删改查） |
| DownModule | `client.Download` | 资源下载（图片/视频/文件/语音/表情/CDN） |
| SnsModule | `client.Sns` | 朋友圈（发布/点赞/评论/转发/隐私设置） |
| FinderModule | `client.Finder` | 视频号（发布/互动/扫码/私信/关注） |

## Token 获取

请访问官网 [www.wechatbot.online](http://www.wechatbot.online) 获取 Token。

---

## 详细用法示例

### 登录认证

```go
// 获取登录二维码
qr, _ := client.Auth.GetQrcode(token, "", "")

// 轮询登录状态
status, _ := client.Auth.RefreshStatus("uuid_value", token, "")

// 检查是否在线
online, _ := client.Auth.CheckStatus(token)

// 已登录过的微信二次弹窗登录
client.Auth.DialogLogin(token)

// 断线重连
client.Auth.Reconnection(token)

// 退出登录
client.Auth.Logout(token)
```

### 消息收发

```go
to := "wxid_xxxxx"

// 发送文本
client.Message.SendText(to, "你好", token, nil)

// 发送文本（群聊中@某人）
client.Message.SendText(to, "大家好", token, []string{"wxid_aaa"})

// 发送图片
client.Message.SendImage(to, "https://example.com/pic.jpg", token)

// 发送语音（silk格式）
client.Message.SendVoice(to, "https://example.com/voice.silk", 5000, token)

// 发送视频
client.Message.SendVideo(to, "https://example.com/video.mp4", "https://example.com/thumb.jpg", 15, token)

// 发送文件
client.Message.SendFile(to, "report.pdf", "https://example.com/report.pdf", token)

// 发送链接卡片
client.Message.SendLink(to, "文章标题", "文章描述", "https://example.com/article", "https://example.com/thumb.jpg", token)

// 发送名片
client.Message.SendCard(to, "张三", "wxid_zhangsan", token)

// 发送小程序
client.Message.SendMiniapp(to, "miniAppId", "userName", "标题", "https://cover.jpg", "/pages/index", "显示名", token)

// 撤回消息
client.Message.RevokeMsg(to, "msg_id", "new_msg_id", 1700000000, token)

// 转发文件（xml从回调消息中获取）
client.Message.ForwardFile(to, "<xml>...</xml>", token)
```

### 群聊管理

```go
roomId := "xxxxx@chatroom"

// 创建群聊（至少3人）
client.Group.CreateGroup([]string{"wxid_aaa", "wxid_bbb", "wxid_ccc"}, token)

// 修改群名
client.Group.ModifyGroupName("新群名", roomId, token)

// 邀请好友进群（多个wxid用逗号隔开）
client.Group.InviteGroupMember("wxid_ddd,wxid_eee", roomId, "欢迎加入", token)

// 移出群成员
client.Group.RemoveGroupMember("wxid_ddd", roomId, token)

// 获取群信息
info, _ := client.Group.GroupInfo(roomId, token)

// 获取群成员详情
detail, _ := client.Group.GroupMemberDetail(roomId, []string{"wxid_aaa", "wxid_bbb"}, token)

// 设置群公告
client.Group.SetAnnouncement(roomId, "本群规则：...", token)

// 管理员操作（1:添加管理 2:删除管理 3:转让群主）
client.Group.AdminOperate(roomId, 1, []string{"wxid_aaa"}, token)

// 聊天置顶 / 消息免打扰
client.Group.PinnedChat(roomId, true, token)
client.Group.SetMsgSilence(roomId, true, token)
```

### 联系人管理

```go
// 获取通讯录（缓存版本更快）
contacts, _ := client.Contact.ContactsListCache(token)

// 获取好友详细信息（最大20个）
detail, _ := client.Contact.DetailInfo([]string{"wxid_aaa", "wxid_bbb"}, token)

// 搜索好友（微信号、手机号）
result, _ := client.Contact.SearchFriend("13800138000", token)

// 添加好友（scene: 3微信号 15手机号；option: 2添加 3同意）
client.Contact.AddContacts(3, "你好，我是xxx", "v4_value", "v3_value", 2, token)

// 设置好友备注
client.Contact.SetFriendRemark("wxid_aaa", "张三-同事", token)

// 检测好友关系
client.Contact.CheckRelation([]string{"wxid_aaa", "wxid_bbb"}, token)
```

### 个人信息

```go
// 获取个人资料
info, _ := client.Personal.GetInfo(token)

// 修改个人资料
client.Personal.UpdateInfo("Shanghai", "CN", "新昵称", "Shanghai", 1, "个性签名", token)

// 修改头像
client.Personal.UpdateHeadImg("https://example.com/avatar.jpg", token)

// 隐私设置（option: 4加我验证 7推荐通讯录 8手机号 25微信号）
client.Personal.PrivacySettings(true, 4, token)
```

### 朋友圈

```go
// 获取我的朋友圈（首次maxId=0，翻页传返回的maxId）
sns, _ := client.Sns.SnsList(0, true, "", token)

// 发送文字朋友圈
client.Sns.SnsSendText(nil, nil, nil, "今天天气真好！", false, nil, nil, token)

// 点赞（operType: 1点赞 2取消）
client.Sns.SnsLike(123456, 1, "wxid_aaa", token)

// 评论
client.Sns.SnsComment(123456, 1, "wxid_aaa", 0, "写得真好！", token)

// 设置朋友圈可见范围（1全部 2半年 3一个月 4三天）
client.Sns.SnsScope(1, token)
```

### 视频号

```go
// 获取我的视频号信息
profile, _ := client.Finder.GetProfile(token, "")

// 搜索视频号
results, _ := client.Finder.FinderSearch("人民日报", token, 1)

// 关注（opType: 1关注 2取消关注）
client.Finder.Follow("my_username", 3, 1, "v2_target@finder", token)

// 发布视频（需ipad协议）
client.Finder.PublishFinderWeb("视频标题", "https://video_url", "https://thumb_url", "#话题", token)

// 浏览视频
client.Finder.Browse("my_username", "nonce_id", "session_buffer", 14195037502970006000, 3, token)

// 点赞
client.Finder.IdFav("my_username", 1, "nonce_id", "session_buffer", 14195037502970006000, "v2_target@finder", 3, token)

// 私信
client.Finder.PostPrivateLetter("你好", "session_id", "my_username", "v2_target@finder", token)
```

---

## 注意事项

- 朋友圈相关接口建议上号 1-3 天后再使用
- 通讯录列表接口（`ContactsList`）为长耗时接口，建议使用缓存版本（`ContactsListCache`）
- 批量查询接口（`BriefInfo`、`DetailInfo`、`CheckRelation`）单次最多 20 个
- 视频号部分接口（`PublishFinderWeb`、`SyncPrivateLetterMsg`）需使用 iPad 协议登录

## 项目结构

```
wechatbot-go-sdk/
├── go.mod           模块定义
├── client.go        主客户端（Functional Options）
├── response.go      统一响应类型
├── auth.go          登录认证
├── message.go       消息收发
├── group.go         群聊管理
├── contact.go       联系人管理
├── personal.go      个人信息
├── favorite.go      收藏夹
├── label.go         标签管理
├── download.go      资源下载
├── sns.go           朋友圈
└── finder.go        视频号
```

## 许可证

[MIT License](LICENSE)
