package wechatbot

type SnsModule struct {
	client *Client
}

func (m *SnsModule) SnsList(maxId int64, decrypt bool, firstPageMd5, token string) (*Response, error) {
	return m.client.Request("/snsList", token, map[string]any{
		"maxId": maxId, "decrypt": decrypt, "firstPageMd5": firstPageMd5,
	})
}

func (m *SnsModule) FriendsSnsList(maxId int64, decrypt bool, wxid, firstPageMd5, token string) (*Response, error) {
	return m.client.Request("/contactsSnsList", token, map[string]any{
		"maxId": maxId, "decrypt": decrypt, "wxid": wxid, "firstPageMd5": firstPageMd5,
	})
}

func (m *SnsModule) SnsDetails(snsId int64, token string) (*Response, error) {
	return m.client.Request("/snsDetails", token, map[string]any{"snsId": snsId})
}

func (m *SnsModule) SnsLike(snsId int64, operType int, wxid, token string) (*Response, error) {
	return m.client.Request("/likeSns", token, map[string]any{
		"snsId": snsId, "operType": operType, "wxid": wxid,
	})
}

func (m *SnsModule) SnsComment(snsId int64, operType int, wxid string, commentId int, content, token string) (*Response, error) {
	return m.client.Request("/commentSns", token, map[string]any{
		"snsId": snsId, "operType": operType, "wxid": wxid,
		"commentId": commentId, "content": content,
	})
}

func (m *SnsModule) SnsDelete(snsId int64, token string) (*Response, error) {
	return m.client.Request("/delSns", token, map[string]any{"snsId": snsId})
}

func (m *SnsModule) SnsScope(option int, token string) (*Response, error) {
	return m.client.Request("/snsVisibleScope", token, map[string]any{"option": option})
}

func (m *SnsModule) SnsVisibilityEnable(enabled bool, token string) (*Response, error) {
	return m.client.Request("/strangerVisibilityEnabled", token, map[string]any{"enabled": enabled})
}

func (m *SnsModule) SnsSetStatus(snsId int64, open bool, token string) (*Response, error) {
	return m.client.Request("/snsSetPrivacy", token, map[string]any{"snsId": snsId, "open": open})
}

func (m *SnsModule) SnsDownloadVideo(snsXml, token string) (*Response, error) {
	return m.client.Request("/downloadSnsVideo", token, map[string]any{"snsXml": snsXml})
}

func (m *SnsModule) SnsSendText(allowWxIds, atWxIds, disableWxIds []string, content string, privacy bool, allowTagIds, disableTagIds []int, token string) (*Response, error) {
	return m.client.Request("/sendTextSns", token, map[string]any{
		"allowWxIds": allowWxIds, "atWxIds": atWxIds, "disableWxIds": disableWxIds,
		"content": content, "privacy": privacy,
		"allowTagIds": allowTagIds, "disableTagIds": disableTagIds,
	})
}

func (m *SnsModule) SnsUploadImage(imgUrls []string, token string) (*Response, error) {
	return m.client.Request("/uploadSnsImage", token, map[string]any{"imgUrls": imgUrls})
}

func (m *SnsModule) SnsUploadVideo(thumbUrl, videoUrl, token string) (*Response, error) {
	return m.client.Request("/uploadSnsVideo", token, map[string]any{
		"thumbUrl": thumbUrl, "videoUrl": videoUrl,
	})
}

func (m *SnsModule) SnsSendUrl(allowWxIds, atWxIds, disableWxIds []string, content, description, title, linkUrl, thumbUrl string, privacy bool, allowTagIds, disableTagIds []int, token string) (*Response, error) {
	return m.client.Request("/sendUrlSns", token, map[string]any{
		"allowWxIds": allowWxIds, "atWxIds": atWxIds, "disableWxIds": disableWxIds,
		"content": content, "description": description, "title": title,
		"linkUrl": linkUrl, "thumbUrl": thumbUrl, "privacy": privacy,
		"allowTagIds": allowTagIds, "disableTagIds": disableTagIds,
	})
}

func (m *SnsModule) SnsForward(allowWxIds, atWxIds, disableWxIds []string, snsXml string, privacy bool, allowTagIds, disableTagIds []int, token string) (*Response, error) {
	return m.client.Request("/forwardSns", token, map[string]any{
		"allowWxIds": allowWxIds, "atWxIds": atWxIds, "disableWxIds": disableWxIds,
		"snsXml": snsXml, "privacy": privacy,
		"allowTagIds": allowTagIds, "disableTagIds": disableTagIds,
	})
}
