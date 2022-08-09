package types

import (
	"time"

	nodetypes "github.com/sentinel-official/hub/x/node/types"

	clienttypes "github.com/sentinel-official/cli-client/types"
)

type Node struct {
	Info
	Address   string            `json:"address"`
	Provider  string            `json:"provider"`
	Price     clienttypes.Coins `json:"price"`
	RemoteURL string            `json:"remote_url"`
	Status    string            `json:"status"`
	StatusAt  time.Time         `json:"status_at"`
}

func (n Node) WithInfo(v Info) Node { n.Info = v; return n }

func NewNodeFromRaw(v *nodetypes.Node) Node {
	return Node{
		Address:   v.Address,
		Provider:  v.Provider,
		Price:     clienttypes.NewCoinsFromRaw(v.Price),
		RemoteURL: v.RemoteURL,
		Status:    v.Status.String(),
		StatusAt:  v.StatusAt,
	}
}

type Nodes []Node

func NewNodesFromRaw(v nodetypes.Nodes) Nodes {
	items := make(Nodes, 0, len(v))
	for i := 0; i < len(v); i++ {
		items = append(items, NewNodeFromRaw(&v[i]))
	}

	return items
}
