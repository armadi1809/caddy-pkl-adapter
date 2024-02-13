package pkladapter

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/caddyserver/caddy/v2/caddyconfig"
)

func init() {
	caddyconfig.RegisterAdapter("toml", Adapter{})
}

// Adapter converts a pkl Caddy configuration to JSON
type Adapter struct{}

// Adapt pkl config to JSON
func (a Adapter) Adapt(body []byte, _ map[string]interface{}) (
	[]byte, []caddyconfig.Warning, error) {

	evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions, func(options *pkl.EvaluatorOptions) {
		options.OutputFormat = "json"
	})
	if err != nil {
		return nil, nil, err
	}
	defer evaluator.Close()
	res, err := evaluator.EvaluateOutputText(context.Background(), pkl.TextSource(string(body)))
	if err != nil {
		return nil, nil, err
	}
	return []byte(res), nil, nil
}
