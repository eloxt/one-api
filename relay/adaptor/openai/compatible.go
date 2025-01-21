package openai

import (
	"github.com/eloxt/one-api/relay/adaptor/ai360"
	"github.com/eloxt/one-api/relay/adaptor/baichuan"
	"github.com/eloxt/one-api/relay/adaptor/deepseek"
	"github.com/eloxt/one-api/relay/adaptor/doubao"
	"github.com/eloxt/one-api/relay/adaptor/groq"
	"github.com/eloxt/one-api/relay/adaptor/lingyiwanwu"
	"github.com/eloxt/one-api/relay/adaptor/minimax"
	"github.com/eloxt/one-api/relay/adaptor/mistral"
	"github.com/eloxt/one-api/relay/adaptor/moonshot"
	"github.com/eloxt/one-api/relay/adaptor/novita"
	"github.com/eloxt/one-api/relay/adaptor/siliconflow"
	"github.com/eloxt/one-api/relay/adaptor/stepfun"
	"github.com/eloxt/one-api/relay/adaptor/togetherai"
	"github.com/eloxt/one-api/relay/adaptor/xai"
	"github.com/eloxt/one-api/relay/channeltype"
)

var CompatibleChannels = []int{
	channeltype.Azure,
	channeltype.AI360,
	channeltype.Moonshot,
	channeltype.Baichuan,
	channeltype.Minimax,
	channeltype.Doubao,
	channeltype.Mistral,
	channeltype.Groq,
	channeltype.LingYiWanWu,
	channeltype.StepFun,
	channeltype.DeepSeek,
	channeltype.TogetherAI,
	channeltype.Novita,
	channeltype.SiliconFlow,
	channeltype.XAI,
}

func GetCompatibleChannelMeta(channelType int) (string, []string) {
	switch channelType {
	case channeltype.Azure:
		return "azure", ModelList
	case channeltype.AI360:
		return "360", ai360.ModelList
	case channeltype.Moonshot:
		return "moonshot", moonshot.ModelList
	case channeltype.Baichuan:
		return "baichuan", baichuan.ModelList
	case channeltype.Minimax:
		return "minimax", minimax.ModelList
	case channeltype.Mistral:
		return "mistralai", mistral.ModelList
	case channeltype.Groq:
		return "groq", groq.ModelList
	case channeltype.LingYiWanWu:
		return "lingyiwanwu", lingyiwanwu.ModelList
	case channeltype.StepFun:
		return "stepfun", stepfun.ModelList
	case channeltype.DeepSeek:
		return "deepseek", deepseek.ModelList
	case channeltype.TogetherAI:
		return "together.ai", togetherai.ModelList
	case channeltype.Doubao:
		return "doubao", doubao.ModelList
	case channeltype.Novita:
		return "novita", novita.ModelList
	case channeltype.SiliconFlow:
		return "siliconflow", siliconflow.ModelList
	case channeltype.XAI:
		return "xai", xai.ModelList
	default:
		return "openai", ModelList
	}
}
