package wechatbot

type ContactModule struct {
	client *Client
}

func (m *ContactModule) ContactsList(token string) (*Response, error) {
	return m.client.Request("/fetchContactsList", token, map[string]any{})
}

func (m *ContactModule) ContactsListCache(token string) (*Response, error) {
	return m.client.Request("/fetchContactsListCache", token, map[string]any{})
}

func (m *ContactModule) BriefInfo(wxids []string, token string) (*Response, error) {
	return m.client.Request("/getBriefInfo", token, map[string]any{"wxids": wxids})
}

func (m *ContactModule) DetailInfo(wxids []string, token string) (*Response, error) {
	return m.client.Request("/getDetailInfo", token, map[string]any{"wxids": wxids})
}

func (m *ContactModule) SearchFriend(contactsInfo, token string) (*Response, error) {
	return m.client.Request("/search", token, map[string]any{"contactsInfo": contactsInfo})
}

func (m *ContactModule) AddContacts(scene int, content, v4, v3 string, option int, token string) (*Response, error) {
	return m.client.Request("/addContacts", token, map[string]any{
		"scene": scene, "content": content, "v4": v4, "v3": v3, "option": option,
	})
}

func (m *ContactModule) DeleteFriend(wxid, token string) (*Response, error) {
	return m.client.Request("/deleteFriend", token, map[string]any{"wxid": wxid})
}

func (m *ContactModule) SetFriendPermissions(wxid string, onlyChat bool, token string) (*Response, error) {
	return m.client.Request("/setFriendPermissions", token, map[string]any{
		"wxid": wxid, "onlyChat": onlyChat,
	})
}

func (m *ContactModule) SetFriendRemark(wxid, remark, token string) (*Response, error) {
	return m.client.Request("/setFriendRemark", token, map[string]any{
		"wxid": wxid, "remark": remark,
	})
}

func (m *ContactModule) GetPhoneList(token string, phones []string) (*Response, error) {
	return m.client.Request("/getPhoneAddressList", token, map[string]any{"phones": phones})
}

func (m *ContactModule) UploadPhoneList(phones []string, opType int, token string) (*Response, error) {
	return m.client.Request("/uploadPhoneAddressList", token, map[string]any{
		"phones": phones, "opType": opType,
	})
}

func (m *ContactModule) ImSearch(scene int, content, token string) (*Response, error) {
	return m.client.Request("/imSearch", token, map[string]any{
		"scene": scene, "content": content,
	})
}

func (m *ContactModule) AddImFriends(v3, v4, token string) (*Response, error) {
	return m.client.Request("/imAddFriends", token, map[string]any{"v3": v3, "v4": v4})
}

func (m *ContactModule) SyncImFriends(token string) (*Response, error) {
	return m.client.Request("/imSyncFriends", token, map[string]any{})
}

func (m *ContactModule) DetailImFriends(toUserName, token string) (*Response, error) {
	return m.client.Request("/imDetailFriends", token, map[string]any{"toUserName": toUserName})
}

func (m *ContactModule) CheckRelation(wxids []string, token string) (*Response, error) {
	return m.client.Request("/checkRelation", token, map[string]any{"wxids": wxids})
}
