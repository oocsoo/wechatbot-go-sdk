package wechatbot

type AuthModule struct {
	client *Client
}

func (m *AuthModule) GetQrcode(token string, deviceType string, aid string) (*Response, error) {
	return m.client.Request("/getLoginQrCode", token, map[string]any{
		"type": deviceType,
		"aid":  aid,
	})
}

func (m *AuthModule) RefreshStatus(uuid string, token string, capCode string) (*Response, error) {
	return m.client.Request("/refreshStatus", token, map[string]any{
		"uuid":         uuid,
		"auto_sliding": "false",
		"cap_code":     capCode,
	})
}

func (m *AuthModule) DialogLogin(token string) (*Response, error) {
	return m.client.Request("/dialogLogin", token, map[string]any{})
}

func (m *AuthModule) Reconnection(token string) (*Response, error) {
	return m.client.Request("/reconnection", token, map[string]any{})
}

func (m *AuthModule) Logout(token string) (*Response, error) {
	return m.client.Request("/logout", token, map[string]any{})
}

func (m *AuthModule) CheckStatus(token string) (*Response, error) {
	return m.client.Request("/checkOnline", token, map[string]any{})
}
