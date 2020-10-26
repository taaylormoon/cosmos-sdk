package sdk

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
)

type Client struct {
	endpoint string
	cdc      *codec.Codec
}

// NewClient returns the client to call Cosmos RPC.
func NewClient(endpoint string, cdc *codec.Codec) *Client {
	return &Client{
		endpoint: endpoint,
		cdc:      cdc,
	}
}

func (c Client) getEndpoint(path string) string {
	return fmt.Sprintf("%s%s", c.endpoint, path)
}