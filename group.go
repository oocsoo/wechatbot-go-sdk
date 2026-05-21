package wechatbot

type GroupModule struct {
	client *Client
}

func (m *GroupModule) CreateGroup(wxids []string, token string) (*Response, error) {
	return m.client.Request("/createGroup", token, map[string]any{"wxids": wxids})
}

func (m *GroupModule) ModifyGroupName(chatroomName, chatroomId, token string) (*Response, error) {
	return m.client.Request("/modifyGroupName", token, map[string]any{
		"chatroomName": chatroomName, "chatroomId": chatroomId,
	})
}

func (m *GroupModule) ModifyGroupRemark(chatroomRemark, chatroomId, token string) (*Response, error) {
	return m.client.Request("/modifyGroupRemark", token, map[string]any{
		"chatroomRemark": chatroomRemark, "chatroomId": chatroomId,
	})
}

func (m *GroupModule) ModifySelfNicknameInGroup(nickName, chatroomId, token string) (*Response, error) {
	return m.client.Request("/modifyGroupNickNameForSelf", token, map[string]any{
		"nickName": nickName, "chatroomId": chatroomId,
	})
}

func (m *GroupModule) InviteGroupMember(wxids, chatroomId, reason, token string) (*Response, error) {
	return m.client.Request("/inviteMember", token, map[string]any{
		"wxids": wxids, "chatroomId": chatroomId, "reason": reason,
	})
}

func (m *GroupModule) RemoveGroupMember(wxids, chatroomId, token string) (*Response, error) {
	return m.client.Request("/removeMember", token, map[string]any{
		"wxids": wxids, "chatroomId": chatroomId,
	})
}

func (m *GroupModule) QuitGroup(chatroomId, token string) (*Response, error) {
	return m.client.Request("/quitGroup", token, map[string]any{"chatroomId": chatroomId})
}

func (m *GroupModule) DisbandGroup(chatroomId, token string) (*Response, error) {
	return m.client.Request("/disbandGroup", token, map[string]any{"chatroomId": chatroomId})
}

func (m *GroupModule) GroupInfo(chatroomId, token string) (*Response, error) {
	return m.client.Request("/getGroupInfo", token, map[string]any{"chatroomId": chatroomId})
}

func (m *GroupModule) GroupMember(chatroomId, token string) (*Response, error) {
	return m.client.Request("/getGroupMemberList", token, map[string]any{"chatroomId": chatroomId})
}

func (m *GroupModule) GroupMemberDetail(chatroomId string, memberWxids []string, token string) (*Response, error) {
	return m.client.Request("/getGroupMemberDetail", token, map[string]any{
		"chatroomId": chatroomId, "memberWxids": memberWxids,
	})
}

func (m *GroupModule) GetAnnouncement(chatroomId, token string) (*Response, error) {
	return m.client.Request("/getGroupAnnouncement", token, map[string]any{"chatroomId": chatroomId})
}

func (m *GroupModule) SetAnnouncement(chatroomId, content, token string) (*Response, error) {
	return m.client.Request("/setGroupAnnouncement", token, map[string]any{
		"chatroomId": chatroomId, "content": content,
	})
}

func (m *GroupModule) AgreeJoinGroup(url, token string) (*Response, error) {
	return m.client.Request("/agreeJoinGroup", token, map[string]any{"url": url})
}

func (m *GroupModule) AddGroupMemberAsFriend(chatroomId, content, memberWxid, token string) (*Response, error) {
	return m.client.Request("/addGroupMemberAsFriend", token, map[string]any{
		"chatroomId": chatroomId, "content": content, "memberWxid": memberWxid,
	})
}

func (m *GroupModule) GetGroupQr(chatroomId, token string) (*Response, error) {
	return m.client.Request("/getGroupQrCode", token, map[string]any{"chatroomId": chatroomId})
}

func (m *GroupModule) SaveContractList(chatroomId string, operType int, token string) (*Response, error) {
	return m.client.Request("/saveContractList", token, map[string]any{
		"chatroomId": chatroomId, "operType": operType,
	})
}

func (m *GroupModule) AdminOperate(chatroomId string, operType int, wxids []string, token string) (*Response, error) {
	return m.client.Request("/adminOperate", token, map[string]any{
		"chatroomId": chatroomId, "operType": operType, "wxids": wxids,
	})
}

func (m *GroupModule) PinnedChat(chatroomId string, top bool, token string) (*Response, error) {
	return m.client.Request("/pinnedChat", token, map[string]any{
		"chatroomId": chatroomId, "top": top,
	})
}

func (m *GroupModule) SetMsgSilence(chatroomId string, silence bool, token string) (*Response, error) {
	return m.client.Request("/setMsgSilence", token, map[string]any{
		"chatroomId": chatroomId, "silence": silence,
	})
}

func (m *GroupModule) JoinGroupByQr(qrUrl, token string) (*Response, error) {
	return m.client.Request("/joinGroupUsingQRCode", token, map[string]any{"qrUrl": qrUrl})
}

func (m *GroupModule) ApplyGroupApprove(chatroomId, msgContent, newMsgId, token string) (*Response, error) {
	return m.client.Request("/groupAccessApplyCheckApprove", token, map[string]any{
		"chatroomId": chatroomId, "msgContent": msgContent, "newMsgId": newMsgId,
	})
}
