package main

import (
	"net/url"
)

var (
	//FormData Formulario para la consulta de vuelos
	FormData = url.Values{}
)

func init() {
	FormData.Add("LOCALE", "es")
	FormData.Add("SITE", "AAVRAAVR")
	FormData.Add("LANGUAGE", "ES")
	FormData.Add("TRIPFLOW", "YES")
	FormData.Add("SO_SITE_OFFICE_ID", "CCSV008AA")
	FormData.Add("SO_QUEUE_OFFICE_ID", "CCSV008AA")
	FormData.Add("DIRECT_LOGIN", "NO")
	FormData.Add("B_ANY_TIME_1", "FALSE")
	FormData.Add("B_ANY_TIME_2", "TRUE")
	FormData.Add("SO_SITE_PRIMARY_CURRENCY", "VEF")
	FormData.Add("SO_SITE_USER_CURRENCY_CODE", "VEF")
	FormData.Add("ARRANGE_BY", "")
	FormData.Add("REFRESH", "0")
	FormData.Add("DISPLAY_TYPE", "1")
	FormData.Add("EMBEDDED_TRANSACTION", "FlexPricerAvailability")
	FormData.Add("DATE_RANGE_VALUE_1", "7")
	FormData.Add("DATE_RANGE_VALUE_2", "7")
	FormData.Add("DATE_RANGE_QUALIFIER_1", "C")
	FormData.Add("DATE_RANGE_QUALIFIER_2", "C")
	FormData.Add("COMMERCIAL_FARE_FAMILY_1", "VENEZUELA")
	FormData.Add("PRICING_TYPE", "I")
	FormData.Add("SEVEN_DAY_SEARCH", "TRUE")

	FormData.Add("B_DATE_1", "201712151200")
	FormData.Add("B_DATE_2", "201612210000")
	FormData.Add("TRAVELLER_TYPE_1", "ADT")
	FormData.Add("TRAVELLER_TYPE_2", "")
	FormData.Add("TRAVELLER_TYPE_3", "")
	FormData.Add("TRAVELLER_TYPE_4", "")
	FormData.Add("HAS_INFANT_1", "FALSE")
	FormData.Add("HAS_INFANT_2", "FALSE")
	FormData.Add("HAS_INFANT_3", "FALSE")
	FormData.Add("HAS_INFANT_4", "FALSE")
	FormData.Add("SWITCHE", "0")
	FormData.Add("TRIP_TYPE", "R")
	FormData.Add("EXTERNAL_ID", "US")
	FormData.Add("SO_SITE_MOD_E_TICKET", "TRUE")
	FormData.Add("SO_LANG_SITE_AGENCY_LINE1", `<a target='blank' href='http", "//conviasa.aero' >Conviasa</a>`)
	FormData.Add("SO_LANG_SITE_AGENCY_LINE2", "VENTAS INTERNET")
	FormData.Add("SO_LANG_SITE_AGENCY_LINE3", "AV. INTERCOMUNAL AEROPUERTO DE MAIQUETIA")
	FormData.Add("SO_LANG_SITE_AGENCY_LINE4", "HANGARES DE CONVIASA")
	FormData.Add("SO_LANG_SITE_AGENCY_LINE5", "MAIQUETIA, EDO. VARGAS, VENEZUELA.")
	FormData.Add("SO_LANG_SITE_AGENCY_LINE6", "+58(212)303.31.37")
	FormData.Add("SO_LANG_SITE_EMAIL_ADDRESS", "ventas.psp@conviasa.aero")
	FormData.Add("SO_SITE_MOP_EXT", "TRUE")
	FormData.Add("opcion_iv", "on")
	FormData.Add("ORIGEN_VISUALIZAR", "Caracas (CCS), VE")

	FormData.Add("DESTINO_VISUALIZAR", "Maracaibo (MAR), VE")

	FormData.Add("fecha_desde", "15/12/2017")
	FormData.Add("fecha_hasta", "21/12/2017")
	FormData.Add("adultos", "1")
	FormData.Add("ninos", "0")
	FormData.Add("bebes", "0")
}
