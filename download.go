package wechatbot

import (
	"encoding/base64"
	"os"
)

type DownModule struct {
	client *Client
}

func (m *DownModule) DownloadSilkBase64(base64Data, savePath string) error {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}
	return os.WriteFile(savePath, data, 0644)
}

func (m *DownModule) DownloadSilkRequest(msgId, xml, token string) (*Response, error) {
	return m.client.Request("/downloadVoice", token, map[string]any{
		"msgId": msgId, "xml": xml,
	})
}

func (m *DownModule) DownloadFile(xml, token string) (*Response, error) {
	return m.client.Request("/downloadFile", token, map[string]any{"xml": xml})
}

func (m *DownModule) DownloadImage(imgType int, xml, token string) (*Response, error) {
	return m.client.Request("/downloadImage", token, map[string]any{
		"type": imgType, "xml": xml,
	})
}

func (m *DownModule) DownloadVideo(xml, token string) (*Response, error) {
	return m.client.Request("/downloadVideo", token, map[string]any{"xml": xml})
}

func (m *DownModule) DownloadEmoji(emojiMd5, token string) (*Response, error) {
	return m.client.Request("/downloadEmoji", token, map[string]any{"emojiMd5": emojiMd5})
}

func (m *DownModule) DownloadCdn(aesKey, totalSize, fileType, fileId, suffix, token string) (*Response, error) {
	return m.client.Request("/downloadCdn", token, map[string]any{
		"aesKey": aesKey, "totalSize": totalSize, "type": fileType,
		"fileId": fileId, "suffix": suffix,
	})
}
