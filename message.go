package wechatbot

type MessageModule struct {
	client *Client
}

func (m *MessageModule) SendText(toWxid, content, token string, atList []string) (*Response, error) {
	if atList == nil {
		atList = []string{}
	}
	return m.client.Request("/sendTextMessage", token, map[string]any{
		"toWxid":     toWxid,
		"content":    content,
		"atWxidList": atList,
	})
}

func (m *MessageModule) SendImage(toWxid, imageUrl, token string) (*Response, error) {
	return m.client.Request("/sendImageMessage", token, map[string]any{
		"toWxid": toWxid,
		"imgUrl": imageUrl,
	})
}

func (m *MessageModule) SendVoice(toWxid, silkUrl string, voiceDuration int, token string) (*Response, error) {
	return m.client.Request("/sendVoiceMessage", token, map[string]any{
		"toWxid":        toWxid,
		"voiceUrl":      silkUrl,
		"voiceDuration": voiceDuration,
	})
}

func (m *MessageModule) SendVideo(toWxid, videoUrl, thumbUrl string, videoDuration int, token string) (*Response, error) {
	return m.client.Request("/sendVedioMessage", token, map[string]any{
		"toWxid":        toWxid,
		"videoUrl":      videoUrl,
		"thumbUrl":      thumbUrl,
		"videoDuration": videoDuration,
	})
}

func (m *MessageModule) SendFile(toWxid, fileName, fileUrl, token string) (*Response, error) {
	return m.client.Request("/sendFileMessage", token, map[string]any{
		"toWxid":   toWxid,
		"fileName": fileName,
		"fileUrl":  fileUrl,
	})
}

func (m *MessageModule) SendLink(toWxid, title, desc, linkUrl, thumbUrl, token string) (*Response, error) {
	return m.client.Request("/sendLinkMessage", token, map[string]any{
		"toWxid":  toWxid,
		"title":   title,
		"desc":    desc,
		"linkUrl": linkUrl,
		"thumbUrl": thumbUrl,
	})
}

func (m *MessageModule) SendCard(toWxid, nickname, nameCardWxid, token string) (*Response, error) {
	return m.client.Request("/sendCardMessage", token, map[string]any{
		"toWxid":       toWxid,
		"nickName":     nickname,
		"nameCardWxid": nameCardWxid,
	})
}

func (m *MessageModule) SendEmoji(toWxid, emojiMd5 string, emojiSize int, token string) (*Response, error) {
	return m.client.Request("/sendEmojiMessage", token, map[string]any{
		"toWxid":    toWxid,
		"emojiMd5":  emojiMd5,
		"emojiSize": emojiSize,
	})
}

func (m *MessageModule) SendApp(toWxid, appmsg, token string) (*Response, error) {
	return m.client.Request("/sendAppMessage", token, map[string]any{
		"toWxid": toWxid,
		"appmsg": appmsg,
	})
}

func (m *MessageModule) SendMiniapp(toWxid, miniAppId, userName, title, coverImgUrl, pagePath, displayName, token string) (*Response, error) {
	return m.client.Request("/sendMiniappMessage", token, map[string]any{
		"toWxid":      toWxid,
		"miniAppId":   miniAppId,
		"userName":    userName,
		"title":       title,
		"coverImgUrl": coverImgUrl,
		"pagePath":    pagePath,
		"displayName": displayName,
	})
}

func (m *MessageModule) SendLocation(toWxid, content, token string) (*Response, error) {
	return m.client.Request("/sendLocationMessage", token, map[string]any{
		"toWxid":  toWxid,
		"content": content,
	})
}

func (m *MessageModule) RevokeMsg(toWxid, msgId, newMsgId string, createTime int64, token string) (*Response, error) {
	return m.client.Request("/sendRevokeMessage", token, map[string]any{
		"toWxid":     toWxid,
		"msgId":      msgId,
		"newMsgId":   newMsgId,
		"createTime": createTime,
	})
}

func (m *MessageModule) ForwardFile(toWxid, xml, token string) (*Response, error) {
	return m.client.Request("/sendForwardFileMessage", token, map[string]any{
		"toWxid": toWxid,
		"xml":    xml,
	})
}

func (m *MessageModule) ForwardImage(toWxid, xml, token string) (*Response, error) {
	return m.client.Request("/sendForwardImageMessage", token, map[string]any{
		"toWxid": toWxid,
		"xml":    xml,
	})
}

func (m *MessageModule) ForwardVideo(toWxid, xml, token string) (*Response, error) {
	return m.client.Request("/sendForwardVideoMessage", token, map[string]any{
		"toWxid": toWxid,
		"xml":    xml,
	})
}

func (m *MessageModule) ForwardLink(toWxid, xml, token string) (*Response, error) {
	return m.client.Request("/sendForwardUrlMessage", token, map[string]any{
		"toWxid": toWxid,
		"xml":    xml,
	})
}

func (m *MessageModule) ForwardMiniapp(toWxid, xml, coverImgUrl, token string) (*Response, error) {
	return m.client.Request("/sendForwardMiniAppMessage", token, map[string]any{
		"toWxid":      toWxid,
		"xml":         xml,
		"coverImgUrl": coverImgUrl,
	})
}
