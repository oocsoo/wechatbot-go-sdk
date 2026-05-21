package wechatbot

type PersonalModule struct {
	client *Client
}

func (m *PersonalModule) GetInfo(token string) (*Response, error) {
	return m.client.Request("/getPersonalInfo", token, map[string]any{})
}

func (m *PersonalModule) GetQrcode(token string) (*Response, error) {
	return m.client.Request("/getQrCode", token, map[string]any{})
}

func (m *PersonalModule) GetDeviceRecord(token string) (*Response, error) {
	return m.client.Request("/getSafetyInfo", token, map[string]any{})
}

func (m *PersonalModule) PrivacySettings(open bool, option int, token string) (*Response, error) {
	return m.client.Request("/privacySettings", token, map[string]any{
		"open": open, "option": option,
	})
}

func (m *PersonalModule) UpdateInfo(city, country, nickName, province string, sex int, signature, token string) (*Response, error) {
	return m.client.Request("/updatePersonalInfo", token, map[string]any{
		"city": city, "country": country, "nickName": nickName,
		"province": province, "sex": sex, "signature": signature,
	})
}

func (m *PersonalModule) UpdateHeadImg(headImgUrl, token string) (*Response, error) {
	return m.client.Request("/updateHeadImg", token, map[string]any{"headImgUrl": headImgUrl})
}
