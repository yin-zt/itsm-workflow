package response

type OneOrderInfo struct {
	Title    string   `json:"title"`
	Type     string   `json:"type"`
	Item     []string `json:"item"`
	IsExpand bool     `json:"is_expand"`
}
