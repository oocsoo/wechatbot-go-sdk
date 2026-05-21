package wechatbot

type LabelModule struct {
	client *Client
}

func (m *LabelModule) AddLabel(labelName, token string) (*Response, error) {
	return m.client.Request("/addLabel", token, map[string]any{"labelName": labelName})
}

func (m *LabelModule) ListLabel(token string) (*Response, error) {
	return m.client.Request("/listLabel", token, map[string]any{})
}

func (m *LabelModule) DeleteLabel(labelIds, token string) (*Response, error) {
	return m.client.Request("/deleteLabel", token, map[string]any{"labelIds": labelIds})
}

func (m *LabelModule) ModifyFriendLabel(labelIds string, wxIds []string, token string) (*Response, error) {
	return m.client.Request("/modifyMemberList", token, map[string]any{
		"labelIds": labelIds, "wxIds": wxIds,
	})
}
