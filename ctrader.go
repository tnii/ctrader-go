package ctrader

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/diegobernardes/ctrader/openapi"
)

func mappingResponse(payloadType uint32) (proto.Message, error) {
	var response proto.Message
	switch payloadType {
	case uint32(openapi.ProtoPayloadType_PROTO_MESSAGE):
		response = &openapi.ProtoMessage{}
	case uint32(openapi.ProtoPayloadType_ERROR_RES):
		response = &openapi.ProtoErrorRes{}
	case uint32(openapi.ProtoPayloadType_HEARTBEAT_EVENT):
		response = &openapi.ProtoHeartbeatEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_APPLICATION_AUTH_RES):
		response = &openapi.ProtoOAApplicationAuthRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ACCOUNT_AUTH_RES):
		response = &openapi.ProtoOAAccountAuthRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_VERSION_RES):
		response = &openapi.ProtoOAVersionRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_TRAILING_SL_CHANGED_EVENT):
		response = &openapi.ProtoOATrailingSLChangedEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ASSET_LIST_RES):
		response = &openapi.ProtoOAAssetListRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SYMBOLS_LIST_RES):
		response = &openapi.ProtoOASymbolsListRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SYMBOL_BY_ID_RES):
		response = &openapi.ProtoOASymbolByIdRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SYMBOLS_FOR_CONVERSION_RES):
		response = &openapi.ProtoOASymbolsForConversionRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SYMBOL_CHANGED_EVENT):
		response = &openapi.ProtoOASymbolChangedEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_TRADER_RES):
		response = &openapi.ProtoOATraderRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_TRADER_UPDATE_EVENT):
		response = &openapi.ProtoOAMarginCallUpdateEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_RECONCILE_RES):
		response = &openapi.ProtoOAReconcileRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_EXECUTION_EVENT):
		response = &openapi.ProtoOAExecutionEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SUBSCRIBE_SPOTS_RES):
		response = &openapi.ProtoOASubscribeSpotsRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_UNSUBSCRIBE_SPOTS_RES):
		response = &openapi.ProtoOAUnsubscribeSpotsRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SPOT_EVENT):
		response = &openapi.ProtoOASpotEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ORDER_ERROR_EVENT):
		response = &openapi.ProtoOAOrderErrorEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_DEAL_LIST_RES):
		response = &openapi.ProtoOADealListRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_GET_TRENDBARS_RES):
		response = &openapi.ProtoOAGetTrendbarsRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_EXPECTED_MARGIN_RES):
		response = &openapi.ProtoOAExpectedMarginRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_MARGIN_CHANGED_EVENT):
		response = &openapi.ProtoOAMarginChangedEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ERROR_RES):
		response = &openapi.ProtoOAErrorRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_CASH_FLOW_HISTORY_LIST_RES):
		response = &openapi.ProtoOACashFlowHistoryListRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_GET_TICKDATA_RES):
		response = &openapi.ProtoOAGetTickDataRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ACCOUNTS_TOKEN_INVALIDATED_EVENT):
		response = &openapi.ProtoOAAccountsTokenInvalidatedEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_CLIENT_DISCONNECT_EVENT):
		response = &openapi.ProtoOAClientDisconnectEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_GET_ACCOUNTS_BY_ACCESS_TOKEN_RES):
		response = &openapi.ProtoOAGetAccountListByAccessTokenRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_GET_CTID_PROFILE_BY_TOKEN_RES):
		response = &openapi.ProtoOAGetCtidProfileByTokenRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ASSET_CLASS_LIST_RES):
		response = &openapi.ProtoOAAssetClassListRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_DEPTH_EVENT):
		response = &openapi.ProtoOADepthEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SUBSCRIBE_DEPTH_QUOTES_RES):
		response = &openapi.ProtoOASubscribeDepthQuotesRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_UNSUBSCRIBE_DEPTH_QUOTES_RES):
		response = &openapi.ProtoOAUnsubscribeDepthQuotesRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SYMBOL_CATEGORY_RES):
		response = &openapi.ProtoOASymbolCategoryListRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ACCOUNT_LOGOUT_RES):
		response = &openapi.ProtoOAAccountLogoutRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ACCOUNT_DISCONNECT_EVENT):
		response = &openapi.ProtoOAAccountDisconnectEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_SUBSCRIBE_LIVE_TRENDBAR_RES):
		response = &openapi.ProtoOASubscribeLiveTrendbarRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_UNSUBSCRIBE_LIVE_TRENDBAR_RES):
		response = &openapi.ProtoOAUnsubscribeLiveTrendbarRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_MARGIN_CALL_LIST_RES):
		response = &openapi.ProtoOAMarginCallListRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_MARGIN_CALL_UPDATE_RES):
		response = &openapi.ProtoOAMarginCallUpdateRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_MARGIN_CALL_UPDATE_EVENT):
		response = &openapi.ProtoOAMarginCallUpdateEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_MARGIN_CALL_TRIGGER_EVENT):
		response = &openapi.ProtoOAMarginCallTriggerEvent{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_REFRESH_TOKEN_RES):
		response = &openapi.ProtoOARefreshTokenRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_ORDER_LIST_RES):
		response = &openapi.ProtoOAOrderListRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_GET_DYNAMIC_LEVERAGE_RES):
		response = &openapi.ProtoOAGetDynamicLeverageByIDRes{}
	case uint32(openapi.ProtoOAPayloadType_PROTO_OA_DEAL_LIST_BY_POSITION_ID_RES):
		response = &openapi.ProtoOADealListByPositionIdRes{}
	default:
		return nil, fmt.Errorf("unknow message type '%d'", payloadType)
	}
	return response, nil
}