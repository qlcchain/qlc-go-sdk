package qlcchain

type DebugApi struct {
	client *QLCClient
}

func NewDebugAPI(c *QLCClient) *DebugApi {
	return &DebugApi{client: c}
}

func (l *DebugApi) BlockCacheCount() (map[string]uint64, error) {
	var r map[string]uint64
	err := l.client.getClient().Call(&r, "debug_blockCacheCount")
	if err != nil {
		return nil, err
	}
	return r, nil
}
