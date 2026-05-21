package wechatbot

type FavoriteModule struct {
	client *Client
}

func (m *FavoriteModule) SyncFavorite(syncKey, token string) (*Response, error) {
	return m.client.Request("/syncFavorite", token, map[string]any{"syncKey": syncKey})
}

func (m *FavoriteModule) GetFavorite(favId int, token string) (*Response, error) {
	return m.client.Request("/getFavorite", token, map[string]any{"favId": favId})
}

func (m *FavoriteModule) DeleteFavorite(favId int, token string) (*Response, error) {
	return m.client.Request("/deleteFavorite", token, map[string]any{"favId": favId})
}
