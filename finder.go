package wechatbot

type FinderModule struct {
	client *Client
}

func (m *FinderModule) CreateFinder(proxyIp, signature, headImg, nickName string, sex int, token string) (*Response, error) {
	return m.client.Request("/createFinder", token, map[string]any{
		"proxyIp": proxyIp, "signature": signature, "headImg": headImg,
		"nickName": nickName, "sex": sex,
	})
}

func (m *FinderModule) GetProfile(token string, proxyIp string) (*Response, error) {
	return m.client.Request("/getProfile", token, map[string]any{"proxyIp": proxyIp})
}

func (m *FinderModule) GetQrCode(myUserName string, myRoleType int, token string, proxyIp string) (*Response, error) {
	return m.client.Request("/finderGetQrCode", token, map[string]any{
		"proxyIp": proxyIp, "myUserName": myUserName, "myRoleType": myRoleType,
	})
}

func (m *FinderModule) FinderSearch(content, token string, category int) (*Response, error) {
	return m.client.Request("/finderSearch", token, map[string]any{
		"proxyIp": "", "content": content, "category": category,
		"filter": 0, "page": 0, "cookie": "", "searchId": "", "offset": 0,
	})
}

func (m *FinderModule) Follow(myUserName string, myRoleType, opType int, toUserName, token string) (*Response, error) {
	return m.client.Request("/follow", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "myRoleType": myRoleType,
		"opType": opType, "toUserName": toUserName,
		"searchInfo": map[string]string{"cookies": "", "searchId": "", "docId": ""},
	})
}

func (m *FinderModule) FollowList(myUserName string, myRoleType int, token, proxyIp, lastBuffer string) (*Response, error) {
	return m.client.Request("/followList", token, map[string]any{
		"proxyIp": proxyIp, "myUserName": myUserName,
		"lastBuffer": lastBuffer, "myRoleType": myRoleType,
	})
}

func (m *FinderModule) SearchFollow(myUserName string, myRoleType int, toUserName, keyword, token string) (*Response, error) {
	return m.client.Request("/searchFollow", token, map[string]any{
		"myUserName": myUserName, "myRoleType": myRoleType,
		"toUserName": toUserName, "keyword": keyword,
	})
}

func (m *FinderModule) UserPage(toUserName, token, proxyIp, lastBuffer string, maxId int) (*Response, error) {
	return m.client.Request("/userPage", token, map[string]any{
		"proxyIp": proxyIp, "lastBuffer": lastBuffer,
		"toUserName": toUserName, "maxId": maxId,
	})
}

func (m *FinderModule) Browse(myUserName, objectNonceId, sessionBuffer string, objectId int64, myRoleType int, token string) (*Response, error) {
	return m.client.Request("/browse", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "objectNonceId": objectNonceId,
		"sessionBuffer": sessionBuffer, "objectId": objectId, "myRoleType": myRoleType,
	})
}

func (m *FinderModule) Comment(myUserName string, opType int, objectNonceId, sessionBuffer string, objectId int64, myRoleType int, content, commentId, token string) (*Response, error) {
	return m.client.Request("/comment", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "opType": opType,
		"objectNonceId": objectNonceId, "sessionBuffer": sessionBuffer,
		"objectId": objectId, "myRoleType": myRoleType, "content": content,
		"commentId": commentId, "replyUserName": "", "refCommentId": 0, "rootCommentId": 0,
	})
}

func (m *FinderModule) CommentList(sessionBuffer, objectId, token string) (*Response, error) {
	return m.client.Request("/commentList", token, map[string]any{
		"proxyIp": "", "rootCommentId": 0, "refCommentId": 0,
		"objectNonceId": "", "sessionBuffer": sessionBuffer,
		"lastBuffer": "", "objectId": objectId,
	})
}

func (m *FinderModule) PublishFinderWeb(title, videoUrl, thumbUrl, description, token string) (*Response, error) {
	return m.client.Request("/publishFinderWeb", token, map[string]any{
		"title": title, "videoUrl": videoUrl, "thumbUrl": thumbUrl, "description": description,
	})
}

func (m *FinderModule) UploadFinderVideo(videoUrl, coverImgUrl, token, proxyIp string) (*Response, error) {
	return m.client.Request("/uploadFinderVideo", token, map[string]any{
		"proxyIp": proxyIp, "videoUrl": videoUrl, "coverImgUrl": coverImgUrl,
	})
}

func (m *FinderModule) PublishFinderCdn(myUserName string, myRoleType int, description string, videoCdn map[string]any, token, proxyIp string, topic []string) (*Response, error) {
	if topic == nil {
		topic = []string{}
	}
	return m.client.Request("/publishFinderCdn", token, map[string]any{
		"proxyIp": proxyIp, "topic": topic, "myUserName": myUserName,
		"myRoleType": myRoleType, "description": description, "videoCdn": videoCdn,
	})
}

func (m *FinderModule) MentionList(myUserName string, myRoleType, reqScene int, token, lastBuff string) (*Response, error) {
	return m.client.Request("/mentionList", token, map[string]any{
		"myUserName": myUserName, "lastBuff": lastBuff,
		"myRoleType": myRoleType, "reqScene": reqScene,
	})
}

func (m *FinderModule) LikeFavList(myUserName string, myRoleType, flag int, token, proxyIp, lastBuffer string) (*Response, error) {
	return m.client.Request("/likeFavList", token, map[string]any{
		"proxyIp": proxyIp, "myUserName": myUserName,
		"lastBuffer": lastBuffer, "myRoleType": myRoleType, "flag": flag,
	})
}

func (m *FinderModule) IdFav(myUserName string, opType int, objectNonceId, sessionBuffer string, objectId int64, toUserName string, myRoleType int, token string) (*Response, error) {
	return m.client.Request("/idFav", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "opType": opType,
		"objectNonceId": objectNonceId, "sessionBuffer": sessionBuffer,
		"objectId": objectId, "toUserName": toUserName, "myRoleType": myRoleType,
	})
}

func (m *FinderModule) IdLike(myUserName string, opType int, objectNonceId, sessionBuffer string, objectId int64, toUserName string, myRoleType int, token string) (*Response, error) {
	return m.client.Request("/idLike", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "opType": opType,
		"objectNonceId": objectNonceId, "sessionBuffer": sessionBuffer,
		"objectId": objectId, "toUserName": toUserName, "myRoleType": myRoleType,
	})
}

func (m *FinderModule) FinderOpt(myUserName string, myRoleType int, toUserName string, opType int, id string, remain int, token string) (*Response, error) {
	return m.client.Request("/finderOpt", token, map[string]any{
		"myUserName": myUserName, "myRoleType": myRoleType,
		"toUserName": toUserName, "opType": opType, "id": id, "remain": remain,
	})
}

func (m *FinderModule) SyncPrivateLetterMsg(token, proxyIp, keyBuff string) (*Response, error) {
	return m.client.Request("/syncPrivateLetterMsg", token, map[string]any{
		"proxyIp": proxyIp, "keyBuff": keyBuff,
	})
}

func (m *FinderModule) ContactList(myUserName, queryInfo string, myRoleType int, token, proxyIp string) (*Response, error) {
	return m.client.Request("/contactList", token, map[string]any{
		"proxyIp": proxyIp, "myUserName": myUserName,
		"queryInfo": queryInfo, "myRoleType": myRoleType,
	})
}

func (m *FinderModule) PostPrivateLetter(content, msgSessionId, myUserName, toUserName, token string) (*Response, error) {
	return m.client.Request("/postPrivateLetter", token, map[string]any{
		"proxyIp": "", "content": content, "msgSessionId": msgSessionId,
		"myUserName": myUserName, "toUserName": toUserName,
	})
}

func (m *FinderModule) PostPrivateLetterImg(imgUrl, msgSessionId, myUserName, toUserName, token string) (*Response, error) {
	return m.client.Request("/postPrivateLetterImg", token, map[string]any{
		"proxyIp": "", "imgUrl": imgUrl, "msgSessionId": msgSessionId,
		"myUserName": myUserName, "toUserName": toUserName,
	})
}

func (m *FinderModule) ScanBrowse(myUserName string, myRoleType int, qrContent string, objectId int64, token string) (*Response, error) {
	return m.client.Request("/scanBrowse", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "myRoleType": myRoleType,
		"qrContent": qrContent, "objectId": objectId,
	})
}

func (m *FinderModule) ScanComment(myUserName string, myRoleType int, qrContent string, objectId int64, commentContent, token string) (*Response, error) {
	return m.client.Request("/scanComment", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "myRoleType": myRoleType,
		"qrContent": qrContent, "objectId": objectId, "commentContent": commentContent,
		"replyUsername": "", "refCommentId": 0, "rootCommentId": 0,
	})
}

func (m *FinderModule) ScanFav(myUserName string, myRoleType int, qrContent string, objectId int64, token string) (*Response, error) {
	return m.client.Request("/scanFav", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "myRoleType": myRoleType,
		"qrContent": qrContent, "objectId": objectId,
	})
}

func (m *FinderModule) ScanLike(myUserName string, myRoleType int, qrContent string, objectId int64, token string) (*Response, error) {
	return m.client.Request("/scanLike", token, map[string]any{
		"proxyIp": "", "myUserName": myUserName, "myRoleType": myRoleType,
		"qrContent": qrContent, "objectId": objectId,
	})
}

func (m *FinderModule) ScanLoginChannels(qrContent, token, proxyIp string) (*Response, error) {
	return m.client.Request("/scanLoginChannels", token, map[string]any{
		"proxyIp": proxyIp, "qrContent": qrContent,
	})
}

func (m *FinderModule) ScanQrCode(myUserName string, myRoleType int, qrContent, token, proxyIp string) (*Response, error) {
	return m.client.Request("/scanQrCode", token, map[string]any{
		"proxyIp": proxyIp, "myUserName": myUserName,
		"myRoleType": myRoleType, "qrContent": qrContent,
	})
}

func (m *FinderModule) UpdateProfile(myUserName string, myRoleType int, token string, options map[string]any) (*Response, error) {
	payload := map[string]any{
		"myUserName": myUserName, "myRoleType": myRoleType, "proxyIp": "",
	}
	for k, v := range options {
		payload[k] = v
	}
	return m.client.Request("/updateProfile", token, payload)
}
