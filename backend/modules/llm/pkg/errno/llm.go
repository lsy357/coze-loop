// Code generated by tool. DO NOT EDIT.
// app: cozeloop, biz: llm

package errno

import (
	"github.com/coze-dev/coze-loop/backend/pkg/errorx/code"
)

const (
	CommonNoPermissionCode              = 601500101
	commonNoPermissionMessage           = "no access permission"
	commonNoPermissionNoAffectStability = true

	CommonBadRequestCode              = 601500201
	commonBadRequestMessage           = "bad request"
	commonBadRequestNoAffectStability = true

	CommonInvalidParamCode              = 601500202
	commonInvalidParamMessage           = "invalid param"
	commonInvalidParamNoAffectStability = true

	CommonBizInvalidCode              = 601500203
	commonBizInvalidMessage           = "biz operation is not allowed"
	commonBizInvalidNoAffectStability = true

	CommonResourceDuplicatedCode              = 601500204
	commonResourceDuplicatedMessage           = "resource duplicated"
	commonResourceDuplicatedNoAffectStability = true

	CommonNetworkTimeOutCode              = 601500701
	commonNetworkTimeOutMessage           = "network timeout"
	commonNetworkTimeOutNoAffectStability = false

	CommonInternalErrorCode              = 601500702
	commonInternalErrorMessage           = "internal error"
	commonInternalErrorNoAffectStability = false

	CommonRPCErrorCode              = 601500703
	commonRPCErrorMessage           = "rpc error"
	commonRPCErrorNoAffectStability = false

	CommonMySqlErrorCode              = 601500801
	commonMySqlErrorMessage           = "mysql error"
	commonMySqlErrorNoAffectStability = false

	CommonRedisErrorCode              = 601500803
	commonRedisErrorMessage           = "redis error"
	commonRedisErrorNoAffectStability = false

	UnknownErrCode              = 601505001
	unknownErrMessage           = "unknown err occurred, please oncall"
	unknownErrNoAffectStability = false

	ResourceNotFoundCode              = 601505002
	resourceNotFoundMessage           = "resource not found"
	resourceNotFoundNoAffectStability = true

	ModelQPMLimitCode              = 601505003
	modelQPMLimitMessage           = "request is limited, because qpm of current model reached the upper limit"
	modelQPMLimitNoAffectStability = true

	ModelTPMLimitCode              = 601505004
	modelTPMLimitMessage           = "request is limited, because tpm of current model reached the upper limit"
	modelTPMLimitNoAffectStability = true

	ModelInvalidCode              = 601505005
	modelInvalidMessage           = "model is invalid"
	modelInvalidNoAffectStability = true

	BuildLLMFailedCode              = 601505006
	buildLLMFailedMessage           = "build llm failed"
	buildLLMFailedNoAffectStability = false

	RequestNotValidCode              = 601505007
	requestNotValidMessage           = "request not valid"
	requestNotValidNoAffectStability = false

	RequestNotCompatibleWithModelAbilityCode              = 601505008
	requestNotCompatibleWithModelAbilityMessage           = "request is not compatible with model ability"
	requestNotCompatibleWithModelAbilityNoAffectStability = false

	CallModelFailedCode              = 601505009
	callModelFailedMessage           = "call model failed"
	callModelFailedNoAffectStability = false

	ParseModelRespFailedCode              = 601505010
	parseModelRespFailedMessage           = "parse model response failed"
	parseModelRespFailedNoAffectStability = false
)

func init() {

	code.Register(
		CommonNoPermissionCode,
		commonNoPermissionMessage,
		code.WithAffectStability(!commonNoPermissionNoAffectStability),
	)

	code.Register(
		CommonBadRequestCode,
		commonBadRequestMessage,
		code.WithAffectStability(!commonBadRequestNoAffectStability),
	)

	code.Register(
		CommonInvalidParamCode,
		commonInvalidParamMessage,
		code.WithAffectStability(!commonInvalidParamNoAffectStability),
	)

	code.Register(
		CommonBizInvalidCode,
		commonBizInvalidMessage,
		code.WithAffectStability(!commonBizInvalidNoAffectStability),
	)

	code.Register(
		CommonResourceDuplicatedCode,
		commonResourceDuplicatedMessage,
		code.WithAffectStability(!commonResourceDuplicatedNoAffectStability),
	)

	code.Register(
		CommonNetworkTimeOutCode,
		commonNetworkTimeOutMessage,
		code.WithAffectStability(!commonNetworkTimeOutNoAffectStability),
	)

	code.Register(
		CommonInternalErrorCode,
		commonInternalErrorMessage,
		code.WithAffectStability(!commonInternalErrorNoAffectStability),
	)

	code.Register(
		CommonRPCErrorCode,
		commonRPCErrorMessage,
		code.WithAffectStability(!commonRPCErrorNoAffectStability),
	)

	code.Register(
		CommonMySqlErrorCode,
		commonMySqlErrorMessage,
		code.WithAffectStability(!commonMySqlErrorNoAffectStability),
	)

	code.Register(
		CommonRedisErrorCode,
		commonRedisErrorMessage,
		code.WithAffectStability(!commonRedisErrorNoAffectStability),
	)

	code.Register(
		UnknownErrCode,
		unknownErrMessage,
		code.WithAffectStability(!unknownErrNoAffectStability),
	)

	code.Register(
		ResourceNotFoundCode,
		resourceNotFoundMessage,
		code.WithAffectStability(!resourceNotFoundNoAffectStability),
	)

	code.Register(
		ModelQPMLimitCode,
		modelQPMLimitMessage,
		code.WithAffectStability(!modelQPMLimitNoAffectStability),
	)

	code.Register(
		ModelTPMLimitCode,
		modelTPMLimitMessage,
		code.WithAffectStability(!modelTPMLimitNoAffectStability),
	)

	code.Register(
		ModelInvalidCode,
		modelInvalidMessage,
		code.WithAffectStability(!modelInvalidNoAffectStability),
	)

	code.Register(
		BuildLLMFailedCode,
		buildLLMFailedMessage,
		code.WithAffectStability(!buildLLMFailedNoAffectStability),
	)

	code.Register(
		RequestNotValidCode,
		requestNotValidMessage,
		code.WithAffectStability(!requestNotValidNoAffectStability),
	)

	code.Register(
		RequestNotCompatibleWithModelAbilityCode,
		requestNotCompatibleWithModelAbilityMessage,
		code.WithAffectStability(!requestNotCompatibleWithModelAbilityNoAffectStability),
	)

	code.Register(
		CallModelFailedCode,
		callModelFailedMessage,
		code.WithAffectStability(!callModelFailedNoAffectStability),
	)

	code.Register(
		ParseModelRespFailedCode,
		parseModelRespFailedMessage,
		code.WithAffectStability(!parseModelRespFailedNoAffectStability),
	)

}
