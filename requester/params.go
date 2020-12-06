package requester

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
)

// RequestParams Parametros de request
type RequestParams struct {
	Dy                             ParamItem `json:"dy"`
	PL                             ParamItem `json:"p_L"`
	DividaLiquidaEbit              ParamItem `json:"dividaLiquidaEbit"`
	DividaliquidaPatrimonioLiquido ParamItem `json:"dividaliquidaPatrimonioLiquido"`
	Roe                            ParamItem `json:"roe"`
	Roic                           ParamItem `json:"roic"`
	Roa                            ParamItem `json:"roa"`
	LiquidezMediaDiaria            ParamItem `json:"liquidezMediaDiaria"`
	MyRange                        string    `json:"my_range"`
}

// ParamItem Item de parametro de request
type ParamItem struct {
	Item1 float32
	Item2 float32
}

// CreateParamItem Cria item de parametros de request http
func CreateParamItem(min, max float32) ParamItem {
	return ParamItem{Item1: min, Item2: max}
}

// CreateRequestParams Cria parametros para fazer request http
func CreateRequestParams(inputParams config.Filters) RequestParams {
	params := RequestParams{
		Dy:                             CreateParamItem(inputParams.Dividendo[0], inputParams.Dividendo[1]),
		PL:                             CreateParamItem(inputParams.PL[0], inputParams.PL[1]),
		DividaLiquidaEbit:              CreateParamItem(inputParams.DivLiqEbit[0], inputParams.DivLiqEbit[1]),
		DividaliquidaPatrimonioLiquido: CreateParamItem(inputParams.DivLiqPatLiq[0], inputParams.DivLiqPatLiq[1]),
		Roe:                            CreateParamItem(inputParams.Roe[0], inputParams.Roe[1]),
		Roic:                           CreateParamItem(inputParams.Roic[0], inputParams.Roic[1]),
		Roa:                            CreateParamItem(inputParams.Roa[0], inputParams.Roa[1]),
		LiquidezMediaDiaria:            CreateParamItem(inputParams.LiqMedDiaria[0], inputParams.LiqMedDiaria[1]),
		MyRange:                        "0;25",
	}

	return params
}
