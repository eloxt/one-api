package novita

import (
	"fmt"

	"github.com/eloxt/one-api/relay/meta"
	"github.com/eloxt/one-api/relay/relaymode"
)

func GetRequestURL(meta *meta.Meta) (string, error) {
	if meta.Mode == relaymode.ChatCompletions {
		return fmt.Sprintf("%s/chat/completions", meta.BaseURL), nil
	}
	return "", fmt.Errorf("unsupported relay mode %d for novita", meta.Mode)
}
