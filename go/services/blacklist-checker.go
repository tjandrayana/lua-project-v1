package services

type BlacklistModule struct {
	Dictionary map[string]bool
}

func NewCheckBlacklist() *BlacklistModule {
	return &BlacklistModule{
		Dictionary: getDictionary(),
	}
}

func getDictionary() map[string]bool {
	return map[string]bool{
		"172.256.0.1":   true,
		"188.239.144.1": true,
		"Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405": true,
	}
}

func (m *BlacklistModule) CheckBlacklist(bp BlacklistParameter) bool {
	return m.Dictionary[bp.IPAddress] || m.Dictionary[bp.UserAgent]
}
