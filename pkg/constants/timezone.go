package constants

import (
	"encoding/json"
	"errors"
)

type Timezone string

const (
	TimezoneAfricaAbidjan                  Timezone = "Africa/Abidjan"
	TimezoneAfricaAccra                    Timezone = "Africa/Accra"
	TimezoneAfricaAddisAbaba               Timezone = "Africa/Addis_Ababa"
	TimezoneAfricaAlgiers                  Timezone = "Africa/Algiers"
	TimezoneAfricaAsmara                   Timezone = "Africa/Asmara"
	TimezoneAfricaAsmera                   Timezone = "Africa/Asmera"
	TimezoneAfricaBamako                   Timezone = "Africa/Bamako"
	TimezoneAfricaBangui                   Timezone = "Africa/Bangui"
	TimezoneAfricaBanjul                   Timezone = "Africa/Banjul"
	TimezoneAfricaBissau                   Timezone = "Africa/Bissau"
	TimezoneAfricaBlantyre                 Timezone = "Africa/Blantyre"
	TimezoneAfricaBrazzaville              Timezone = "Africa/Brazzaville"
	TimezoneAfricaBujumbura                Timezone = "Africa/Bujumbura"
	TimezoneAfricaCairo                    Timezone = "Africa/Cairo"
	TimezoneAfricaCasablanca               Timezone = "Africa/Casablanca"
	TimezoneAfricaCeuta                    Timezone = "Africa/Ceuta"
	TimezoneAfricaConakry                  Timezone = "Africa/Conakry"
	TimezoneAfricaDakar                    Timezone = "Africa/Dakar"
	TimezoneAfricaDarEsSalaam              Timezone = "Africa/Dar_es_Salaam"
	TimezoneAfricaDjibouti                 Timezone = "Africa/Djibouti"
	TimezoneAfricaDouala                   Timezone = "Africa/Douala"
	TimezoneAfricaElAaiun                  Timezone = "Africa/El_Aaiun"
	TimezoneAfricaFreetown                 Timezone = "Africa/Freetown"
	TimezoneAfricaGaborone                 Timezone = "Africa/Gaborone"
	TimezoneAfricaHarare                   Timezone = "Africa/Harare"
	TimezoneAfricaJohannesburg             Timezone = "Africa/Johannesburg"
	TimezoneAfricaJuba                     Timezone = "Africa/Juba"
	TimezoneAfricaKampala                  Timezone = "Africa/Kampala"
	TimezoneAfricaKhartoum                 Timezone = "Africa/Khartoum"
	TimezoneAfricaKigali                   Timezone = "Africa/Kigali"
	TimezoneAfricaKinshasa                 Timezone = "Africa/Kinshasa"
	TimezoneAfricaLagos                    Timezone = "Africa/Lagos"
	TimezoneAfricaLibreville               Timezone = "Africa/Libreville"
	TimezoneAfricaLome                     Timezone = "Africa/Lome"
	TimezoneAfricaLuanda                   Timezone = "Africa/Luanda"
	TimezoneAfricaLubumbashi               Timezone = "Africa/Lubumbashi"
	TimezoneAfricaLusaka                   Timezone = "Africa/Lusaka"
	TimezoneAfricaMalabo                   Timezone = "Africa/Malabo"
	TimezoneAfricaMaputo                   Timezone = "Africa/Maputo"
	TimezoneAfricaMaseru                   Timezone = "Africa/Maseru"
	TimezoneAfricaMbabane                  Timezone = "Africa/Mbabane"
	TimezoneAfricaMogadishu                Timezone = "Africa/Mogadishu"
	TimezoneAfricaMonrovia                 Timezone = "Africa/Monrovia"
	TimezoneAfricaNairobi                  Timezone = "Africa/Nairobi"
	TimezoneAfricaNdjamena                 Timezone = "Africa/Ndjamena"
	TimezoneAfricaNiamey                   Timezone = "Africa/Niamey"
	TimezoneAfricaNouakchott               Timezone = "Africa/Nouakchott"
	TimezoneAfricaOuagadougou              Timezone = "Africa/Ouagadougou"
	TimezoneAfricaPortoNovo                Timezone = "Africa/Porto-Novo"
	TimezoneAfricaSaoTome                  Timezone = "Africa/Sao_Tome"
	TimezoneAfricaTimbuktu                 Timezone = "Africa/Timbuktu"
	TimezoneAfricaTripoli                  Timezone = "Africa/Tripoli"
	TimezoneAfricaTunis                    Timezone = "Africa/Tunis"
	TimezoneAfricaWindhoek                 Timezone = "Africa/Windhoek"
	TimezoneAmericaAdak                    Timezone = "America/Adak"
	TimezoneAmericaAnchorage               Timezone = "America/Anchorage"
	TimezoneAmericaAnguilla                Timezone = "America/Anguilla"
	TimezoneAmericaAntigua                 Timezone = "America/Antigua"
	TimezoneAmericaAraguaina               Timezone = "America/Araguaina"
	TimezoneAmericaArgentinaBuenosAires    Timezone = "America/Argentina/Buenos_Aires"
	TimezoneAmericaArgentinaCatamarca      Timezone = "America/Argentina/Catamarca"
	TimezoneAmericaArgentinaComodRivadavia Timezone = "America/Argentina/ComodRivadavia"
	TimezoneAmericaArgentinaCordoba        Timezone = "America/Argentina/Cordoba"
	TimezoneAmericaArgentinaJujuy          Timezone = "America/Argentina/Jujuy"
	TimezoneAmericaArgentinaLaRioja        Timezone = "America/Argentina/La_Rioja"
	TimezoneAmericaArgentinaMendoza        Timezone = "America/Argentina/Mendoza"
	TimezoneAmericaArgentinaRioGallegos    Timezone = "America/Argentina/Rio_Gallegos"
	TimezoneAmericaArgentinaSalta          Timezone = "America/Argentina/Salta"
	TimezoneAmericaArgentinaSanJuan        Timezone = "America/Argentina/San_Juan"
	TimezoneAmericaArgentinaSanLuis        Timezone = "America/Argentina/San_Luis"
	TimezoneAmericaArgentinaTucuman        Timezone = "America/Argentina/Tucuman"
	TimezoneAmericaArgentinaUshuaia        Timezone = "America/Argentina/Ushuaia"
	TimezoneAmericaAruba                   Timezone = "America/Aruba"
	TimezoneAmericaAsuncion                Timezone = "America/Asuncion"
	TimezoneAmericaAtikokan                Timezone = "America/Atikokan"
	TimezoneAmericaAtka                    Timezone = "America/Atka"
	TimezoneAmericaBahia                   Timezone = "America/Bahia"
	TimezoneAmericaBahiaBanderas           Timezone = "America/Bahia_Banderas"
	TimezoneAmericaBarbados                Timezone = "America/Barbados"
	TimezoneAmericaBelem                   Timezone = "America/Belem"
	TimezoneAmericaBelize                  Timezone = "America/Belize"
	TimezoneAmericaBlancSablon             Timezone = "America/Blanc-Sablon"
	TimezoneAmericaBoaVista                Timezone = "America/Boa_Vista"
	TimezoneAmericaBogota                  Timezone = "America/Bogota"
	TimezoneAmericaBoise                   Timezone = "America/Boise"
	TimezoneAmericaBuenosAires             Timezone = "America/Buenos_Aires"
	TimezoneAmericaCambridgeBay            Timezone = "America/Cambridge_Bay"
	TimezoneAmericaCampoGrande             Timezone = "America/Campo_Grande"
	TimezoneAmericaCancun                  Timezone = "America/Cancun"
	TimezoneAmericaCaracas                 Timezone = "America/Caracas"
	TimezoneAmericaCatamarca               Timezone = "America/Catamarca"
	TimezoneAmericaCayenne                 Timezone = "America/Cayenne"
	TimezoneAmericaCayman                  Timezone = "America/Cayman"
	TimezoneAmericaChicago                 Timezone = "America/Chicago"
	TimezoneAmericaChihuahua               Timezone = "America/Chihuahua"
	TimezoneAmericaCoralHarbour            Timezone = "America/Coral_Harbour"
	TimezoneAmericaCordoba                 Timezone = "America/Cordoba"
	TimezoneAmericaCostaRica               Timezone = "America/Costa_Rica"
	TimezoneAmericaCreston                 Timezone = "America/Creston"
	TimezoneAmericaCuiaba                  Timezone = "America/Cuiaba"
	TimezoneAmericaCuracao                 Timezone = "America/Curacao"
	TimezoneAmericaDanmarkshavn            Timezone = "America/Danmarkshavn"
	TimezoneAmericaDawson                  Timezone = "America/Dawson"
	TimezoneAmericaDawsonCreek             Timezone = "America/Dawson_Creek"
	TimezoneAmericaDenver                  Timezone = "America/Denver"
	TimezoneAmericaDetroit                 Timezone = "America/Detroit"
	TimezoneAmericaDominica                Timezone = "America/Dominica"
	TimezoneAmericaEdmonton                Timezone = "America/Edmonton"
	TimezoneAmericaEirunepe                Timezone = "America/Eirunepe"
	TimezoneAmericaElSalvador              Timezone = "America/El_Salvador"
	TimezoneAmericaEnsenada                Timezone = "America/Ensenada"
	TimezoneAmericaFortNelson              Timezone = "America/Fort_Nelson"
	TimezoneAmericaFortWayne               Timezone = "America/Fort_Wayne"
	TimezoneAmericaFortaleza               Timezone = "America/Fortaleza"
	TimezoneAmericaGlaceBay                Timezone = "America/Glace_Bay"
	TimezoneAmericaGodthab                 Timezone = "America/Godthab"
	TimezoneAmericaGooseBay                Timezone = "America/Goose_Bay"
	TimezoneAmericaGrandTurk               Timezone = "America/Grand_Turk"
	TimezoneAmericaGrenada                 Timezone = "America/Grenada"
	TimezoneAmericaGuadeloupe              Timezone = "America/Guadeloupe"
	TimezoneAmericaGuatemala               Timezone = "America/Guatemala"
	TimezoneAmericaGuayaquil               Timezone = "America/Guayaquil"
	TimezoneAmericaGuyana                  Timezone = "America/Guyana"
	TimezoneAmericaHalifax                 Timezone = "America/Halifax"
	TimezoneAmericaHavana                  Timezone = "America/Havana"
	TimezoneAmericaHermosillo              Timezone = "America/Hermosillo"
	TimezoneAmericaIndianaIndianapolis     Timezone = "America/Indiana/Indianapolis"
	TimezoneAmericaIndianaKnox             Timezone = "America/Indiana/Knox"
	TimezoneAmericaIndianaMarengo          Timezone = "America/Indiana/Marengo"
	TimezoneAmericaIndianaPetersburg       Timezone = "America/Indiana/Petersburg"
	TimezoneAmericaIndianaTellCity         Timezone = "America/Indiana/Tell_City"
	TimezoneAmericaIndianaVevay            Timezone = "America/Indiana/Vevay"
	TimezoneAmericaIndianaVincennes        Timezone = "America/Indiana/Vincennes"
	TimezoneAmericaIndianaWinamac          Timezone = "America/Indiana/Winamac"
	TimezoneAmericaIndianapolis            Timezone = "America/Indianapolis"
	TimezoneAmericaInuvik                  Timezone = "America/Inuvik"
	TimezoneAmericaIqaluit                 Timezone = "America/Iqaluit"
	TimezoneAmericaJamaica                 Timezone = "America/Jamaica"
	TimezoneAmericaJujuy                   Timezone = "America/Jujuy"
	TimezoneAmericaJuneau                  Timezone = "America/Juneau"
	TimezoneAmericaKentuckyLouisville      Timezone = "America/Kentucky/Louisville"
	TimezoneAmericaKentuckyMonticello      Timezone = "America/Kentucky/Monticello"
	TimezoneAmericaKnoxIN                  Timezone = "America/Knox_IN"
	TimezoneAmericaKralendijk              Timezone = "America/Kralendijk"
	TimezoneAmericaLaPaz                   Timezone = "America/La_Paz"
	TimezoneAmericaLima                    Timezone = "America/Lima"
	TimezoneAmericaLosAngeles              Timezone = "America/Los_Angeles"
	TimezoneAmericaLouisville              Timezone = "America/Louisville"
	TimezoneAmericaLowerPrinces            Timezone = "America/Lower_Princes"
	TimezoneAmericaMaceio                  Timezone = "America/Maceio"
	TimezoneAmericaManagua                 Timezone = "America/Managua"
	TimezoneAmericaManaus                  Timezone = "America/Manaus"
	TimezoneAmericaMarigot                 Timezone = "America/Marigot"
	TimezoneAmericaMartinique              Timezone = "America/Martinique"
	TimezoneAmericaMatamoros               Timezone = "America/Matamoros"
	TimezoneAmericaMazatlan                Timezone = "America/Mazatlan"
	TimezoneAmericaMendoza                 Timezone = "America/Mendoza"
	TimezoneAmericaMenominee               Timezone = "America/Menominee"
	TimezoneAmericaMerida                  Timezone = "America/Merida"
	TimezoneAmericaMetlakatla              Timezone = "America/Metlakatla"
	TimezoneAmericaMexicoCity              Timezone = "America/Mexico_City"
	TimezoneAmericaMiquelon                Timezone = "America/Miquelon"
	TimezoneAmericaMoncton                 Timezone = "America/Moncton"
	TimezoneAmericaMonterrey               Timezone = "America/Monterrey"
	TimezoneAmericaMontevideo              Timezone = "America/Montevideo"
	TimezoneAmericaMontreal                Timezone = "America/Montreal"
	TimezoneAmericaMontserrat              Timezone = "America/Montserrat"
	TimezoneAmericaNassau                  Timezone = "America/Nassau"
	TimezoneAmericaNewYork                 Timezone = "America/New_York"
	TimezoneAmericaNipigon                 Timezone = "America/Nipigon"
	TimezoneAmericaNome                    Timezone = "America/Nome"
	TimezoneAmericaNoronha                 Timezone = "America/Noronha"
	TimezoneAmericaNorthDakotaBeulah       Timezone = "America/North_Dakota/Beulah"
	TimezoneAmericaNorthDakotaCenter       Timezone = "America/North_Dakota/Center"
	TimezoneAmericaNorthDakotaNewSalem     Timezone = "America/North_Dakota/New_Salem"
	TimezoneAmericaNuuk                    Timezone = "America/Nuuk"
	TimezoneAmericaOjinaga                 Timezone = "America/Ojinaga"
	TimezoneAmericaPanama                  Timezone = "America/Panama"
	TimezoneAmericaPangnirtung             Timezone = "America/Pangnirtung"
	TimezoneAmericaParamaribo              Timezone = "America/Paramaribo"
	TimezoneAmericaPhoenix                 Timezone = "America/Phoenix"
	TimezoneAmericaPortAuPrince            Timezone = "America/Port-au-Prince"
	TimezoneAmericaPortOfSpain             Timezone = "America/Port_of_Spain"
	TimezoneAmericaPortoAcre               Timezone = "America/Porto_Acre"
	TimezoneAmericaPortoVelho              Timezone = "America/Porto_Velho"
	TimezoneAmericaPuertoRico              Timezone = "America/Puerto_Rico"
	TimezoneAmericaPuntaArenas             Timezone = "America/Punta_Arenas"
	TimezoneAmericaRainyRiver              Timezone = "America/Rainy_River"
	TimezoneAmericaRankinInlet             Timezone = "America/Rankin_Inlet"
	TimezoneAmericaRecife                  Timezone = "America/Recife"
	TimezoneAmericaRegina                  Timezone = "America/Regina"
	TimezoneAmericaResolute                Timezone = "America/Resolute"
	TimezoneAmericaRioBranco               Timezone = "America/Rio_Branco"
	TimezoneAmericaRosario                 Timezone = "America/Rosario"
	TimezoneAmericaSantaIsabel             Timezone = "America/Santa_Isabel"
	TimezoneAmericaSantarem                Timezone = "America/Santarem"
	TimezoneAmericaSantiago                Timezone = "America/Santiago"
	TimezoneAmericaSantoDomingo            Timezone = "America/Santo_Domingo"
	TimezoneAmericaSaoPaulo                Timezone = "America/Sao_Paulo"
	TimezoneAmericaScoresbysund            Timezone = "America/Scoresbysund"
	TimezoneAmericaShiprock                Timezone = "America/Shiprock"
	TimezoneAmericaSitka                   Timezone = "America/Sitka"
	TimezoneAmericaStBarthelemy            Timezone = "America/St_Barthelemy"
	TimezoneAmericaStJohns                 Timezone = "America/St_Johns"
	TimezoneAmericaStKitts                 Timezone = "America/St_Kitts"
	TimezoneAmericaStLucia                 Timezone = "America/St_Lucia"
	TimezoneAmericaStThomas                Timezone = "America/St_Thomas"
	TimezoneAmericaStVincent               Timezone = "America/St_Vincent"
	TimezoneAmericaSwiftCurrent            Timezone = "America/Swift_Current"
	TimezoneAmericaTegucigalpa             Timezone = "America/Tegucigalpa"
	TimezoneAmericaThule                   Timezone = "America/Thule"
	TimezoneAmericaThunderBay              Timezone = "America/Thunder_Bay"
	TimezoneAmericaTijuana                 Timezone = "America/Tijuana"
	TimezoneAmericaToronto                 Timezone = "America/Toronto"
	TimezoneAmericaTortola                 Timezone = "America/Tortola"
	TimezoneAmericaVancouver               Timezone = "America/Vancouver"
	TimezoneAmericaVirgin                  Timezone = "America/Virgin"
	TimezoneAmericaWhitehorse              Timezone = "America/Whitehorse"
	TimezoneAmericaWinnipeg                Timezone = "America/Winnipeg"
	TimezoneAmericaYakutat                 Timezone = "America/Yakutat"
	TimezoneAmericaYellowknife             Timezone = "America/Yellowknife"
	TimezoneAntarcticaCasey                Timezone = "Antarctica/Casey"
	TimezoneAntarcticaDavis                Timezone = "Antarctica/Davis"
	TimezoneAntarcticaDumontDUrville       Timezone = "Antarctica/DumontDUrville"
	TimezoneAntarcticaMacquarie            Timezone = "Antarctica/Macquarie"
	TimezoneAntarcticaMawson               Timezone = "Antarctica/Mawson"
	TimezoneAntarcticaMcMurdo              Timezone = "Antarctica/McMurdo"
	TimezoneAntarcticaPalmer               Timezone = "Antarctica/Palmer"
	TimezoneAntarcticaRothera              Timezone = "Antarctica/Rothera"
	TimezoneAntarcticaSouthPole            Timezone = "Antarctica/South_Pole"
	TimezoneAntarcticaSyowa                Timezone = "Antarctica/Syowa"
	TimezoneAntarcticaTroll                Timezone = "Antarctica/Troll"
	TimezoneAntarcticaVostok               Timezone = "Antarctica/Vostok"
	TimezoneArcticLongyearbyen             Timezone = "Arctic/Longyearbyen"
	TimezoneAsiaAden                       Timezone = "Asia/Aden"
	TimezoneAsiaAlmaty                     Timezone = "Asia/Almaty"
	TimezoneAsiaAmman                      Timezone = "Asia/Amman"
	TimezoneAsiaAnadyr                     Timezone = "Asia/Anadyr"
	TimezoneAsiaAqtau                      Timezone = "Asia/Aqtau"
	TimezoneAsiaAqtobe                     Timezone = "Asia/Aqtobe"
	TimezoneAsiaAshgabat                   Timezone = "Asia/Ashgabat"
	TimezoneAsiaAshkhabad                  Timezone = "Asia/Ashkhabad"
	TimezoneAsiaAtyrau                     Timezone = "Asia/Atyrau"
	TimezoneAsiaBaghdad                    Timezone = "Asia/Baghdad"
	TimezoneAsiaBahrain                    Timezone = "Asia/Bahrain"
	TimezoneAsiaBaku                       Timezone = "Asia/Baku"
	TimezoneAsiaBangkok                    Timezone = "Asia/Bangkok"
	TimezoneAsiaBarnaul                    Timezone = "Asia/Barnaul"
	TimezoneAsiaBeirut                     Timezone = "Asia/Beirut"
	TimezoneAsiaBishkek                    Timezone = "Asia/Bishkek"
	TimezoneAsiaBrunei                     Timezone = "Asia/Brunei"
	TimezoneAsiaCalcutta                   Timezone = "Asia/Calcutta"
	TimezoneAsiaChita                      Timezone = "Asia/Chita"
	TimezoneAsiaChoibalsan                 Timezone = "Asia/Choibalsan"
	TimezoneAsiaChongqing                  Timezone = "Asia/Chongqing"
	TimezoneAsiaChungking                  Timezone = "Asia/Chungking"
	TimezoneAsiaColombo                    Timezone = "Asia/Colombo"
	TimezoneAsiaDacca                      Timezone = "Asia/Dacca"
	TimezoneAsiaDamascus                   Timezone = "Asia/Damascus"
	TimezoneAsiaDhaka                      Timezone = "Asia/Dhaka"
	TimezoneAsiaDili                       Timezone = "Asia/Dili"
	TimezoneAsiaDubai                      Timezone = "Asia/Dubai"
	TimezoneAsiaDushanbe                   Timezone = "Asia/Dushanbe"
	TimezoneAsiaFamagusta                  Timezone = "Asia/Famagusta"
	TimezoneAsiaGaza                       Timezone = "Asia/Gaza"
	TimezoneAsiaHarbin                     Timezone = "Asia/Harbin"
	TimezoneAsiaHebron                     Timezone = "Asia/Hebron"
	TimezoneAsiaHoChiMinh                  Timezone = "Asia/Ho_Chi_Minh"
	TimezoneAsiaHongKong                   Timezone = "Asia/Hong_Kong"
	TimezoneAsiaHovd                       Timezone = "Asia/Hovd"
	TimezoneAsiaIrkutsk                    Timezone = "Asia/Irkutsk"
	TimezoneAsiaIstanbul                   Timezone = "Asia/Istanbul"
	TimezoneAsiaJakarta                    Timezone = "Asia/Jakarta"
	TimezoneAsiaJayapura                   Timezone = "Asia/Jayapura"
	TimezoneAsiaJerusalem                  Timezone = "Asia/Jerusalem"
	TimezoneAsiaKabul                      Timezone = "Asia/Kabul"
	TimezoneAsiaKamchatka                  Timezone = "Asia/Kamchatka"
	TimezoneAsiaKarachi                    Timezone = "Asia/Karachi"
	TimezoneAsiaKashgar                    Timezone = "Asia/Kashgar"
	TimezoneAsiaKathmandu                  Timezone = "Asia/Kathmandu"
	TimezoneAsiaKatmandu                   Timezone = "Asia/Katmandu"
	TimezoneAsiaKhandyga                   Timezone = "Asia/Khandyga"
	TimezoneAsiaKolkata                    Timezone = "Asia/Kolkata"
	TimezoneAsiaKrasnoyarsk                Timezone = "Asia/Krasnoyarsk"
	TimezoneAsiaKualaLumpur                Timezone = "Asia/Kuala_Lumpur"
	TimezoneAsiaKuching                    Timezone = "Asia/Kuching"
	TimezoneAsiaKuwait                     Timezone = "Asia/Kuwait"
	TimezoneAsiaMacao                      Timezone = "Asia/Macao"
	TimezoneAsiaMacau                      Timezone = "Asia/Macau"
	TimezoneAsiaMagadan                    Timezone = "Asia/Magadan"
	TimezoneAsiaMakassar                   Timezone = "Asia/Makassar"
	TimezoneAsiaManila                     Timezone = "Asia/Manila"
	TimezoneAsiaMuscat                     Timezone = "Asia/Muscat"
	TimezoneAsiaNicosia                    Timezone = "Asia/Nicosia"
	TimezoneAsiaNovokuznetsk               Timezone = "Asia/Novokuznetsk"
	TimezoneAsiaNovosibirsk                Timezone = "Asia/Novosibirsk"
	TimezoneAsiaOmsk                       Timezone = "Asia/Omsk"
	TimezoneAsiaOral                       Timezone = "Asia/Oral"
	TimezoneAsiaPhnomPenh                  Timezone = "Asia/Phnom_Penh"
	TimezoneAsiaPontianak                  Timezone = "Asia/Pontianak"
	TimezoneAsiaPyongyang                  Timezone = "Asia/Pyongyang"
	TimezoneAsiaQatar                      Timezone = "Asia/Qatar"
	TimezoneAsiaQostanay                   Timezone = "Asia/Qostanay"
	TimezoneAsiaQyzylorda                  Timezone = "Asia/Qyzylorda"
	TimezoneAsiaRangoon                    Timezone = "Asia/Rangoon"
	TimezoneAsiaRiyadh                     Timezone = "Asia/Riyadh"
	TimezoneAsiaSaigon                     Timezone = "Asia/Saigon"
	TimezoneAsiaSakhalin                   Timezone = "Asia/Sakhalin"
	TimezoneAsiaSamarkand                  Timezone = "Asia/Samarkand"
	TimezoneAsiaSeoul                      Timezone = "Asia/Seoul"
	TimezoneAsiaShanghai                   Timezone = "Asia/Shanghai"
	TimezoneAsiaSingapore                  Timezone = "Asia/Singapore"
	TimezoneAsiaSrednekolymsk              Timezone = "Asia/Srednekolymsk"
	TimezoneAsiaTaipei                     Timezone = "Asia/Taipei"
	TimezoneAsiaTashkent                   Timezone = "Asia/Tashkent"
	TimezoneAsiaTbilisi                    Timezone = "Asia/Tbilisi"
	TimezoneAsiaTehran                     Timezone = "Asia/Tehran"
	TimezoneAsiaTelAviv                    Timezone = "Asia/Tel_Aviv"
	TimezoneAsiaThimbu                     Timezone = "Asia/Thimbu"
	TimezoneAsiaThimphu                    Timezone = "Asia/Thimphu"
	TimezoneAsiaTokyo                      Timezone = "Asia/Tokyo"
	TimezoneAsiaTomsk                      Timezone = "Asia/Tomsk"
	TimezoneAsiaUjungPandang               Timezone = "Asia/Ujung_Pandang"
	TimezoneAsiaUlaanbaatar                Timezone = "Asia/Ulaanbaatar"
	TimezoneAsiaUlanBator                  Timezone = "Asia/Ulan_Bator"
	TimezoneAsiaUrumqi                     Timezone = "Asia/Urumqi"
	TimezoneAsiaUstNera                    Timezone = "Asia/Ust-Nera"
	TimezoneAsiaVientiane                  Timezone = "Asia/Vientiane"
	TimezoneAsiaVladivostok                Timezone = "Asia/Vladivostok"
	TimezoneAsiaYakutsk                    Timezone = "Asia/Yakutsk"
	TimezoneAsiaYangon                     Timezone = "Asia/Yangon"
	TimezoneAsiaYekaterinburg              Timezone = "Asia/Yekaterinburg"
	TimezoneAsiaYerevan                    Timezone = "Asia/Yerevan"
	TimezoneAtlanticAzores                 Timezone = "Atlantic/Azores"
	TimezoneAtlanticBermuda                Timezone = "Atlantic/Bermuda"
	TimezoneAtlanticCanary                 Timezone = "Atlantic/Canary"
	TimezoneAtlanticCapeVerde              Timezone = "Atlantic/Cape_Verde"
	TimezoneAtlanticFaeroe                 Timezone = "Atlantic/Faeroe"
	TimezoneAtlanticFaroe                  Timezone = "Atlantic/Faroe"
	TimezoneAtlanticJanMayen               Timezone = "Atlantic/Jan_Mayen"
	TimezoneAtlanticMadeira                Timezone = "Atlantic/Madeira"
	TimezoneAtlanticReykjavik              Timezone = "Atlantic/Reykjavik"
	TimezoneAtlanticSouthGeorgia           Timezone = "Atlantic/South_Georgia"
	TimezoneAtlanticStHelena               Timezone = "Atlantic/St_Helena"
	TimezoneAtlanticStanley                Timezone = "Atlantic/Stanley"
	TimezoneAustraliaACT                   Timezone = "Australia/ACT"
	TimezoneAustraliaAdelaide              Timezone = "Australia/Adelaide"
	TimezoneAustraliaBrisbane              Timezone = "Australia/Brisbane"
	TimezoneAustraliaBrokenHill            Timezone = "Australia/Broken_Hill"
	TimezoneAustraliaCanberra              Timezone = "Australia/Canberra"
	TimezoneAustraliaCurrie                Timezone = "Australia/Currie"
	TimezoneAustraliaDarwin                Timezone = "Australia/Darwin"
	TimezoneAustraliaEucla                 Timezone = "Australia/Eucla"
	TimezoneAustraliaHobart                Timezone = "Australia/Hobart"
	TimezoneAustraliaLHI                   Timezone = "Australia/LHI"
	TimezoneAustraliaLindeman              Timezone = "Australia/Lindeman"
	TimezoneAustraliaLordHowe              Timezone = "Australia/Lord_Howe"
	TimezoneAustraliaMelbourne             Timezone = "Australia/Melbourne"
	TimezoneAustraliaNorth                 Timezone = "Australia/North"
	TimezoneAustraliaNSW                   Timezone = "Australia/NSW"
	TimezoneAustraliaPerth                 Timezone = "Australia/Perth"
	TimezoneAustraliaQueensland            Timezone = "Australia/Queensland"
	TimezoneAustraliaSouth                 Timezone = "Australia/South"
	TimezoneAustraliaSydney                Timezone = "Australia/Sydney"
	TimezoneAustraliaTasmania              Timezone = "Australia/Tasmania"
	TimezoneAustraliaVictoria              Timezone = "Australia/Victoria"
	TimezoneAustraliaWest                  Timezone = "Australia/West"
	TimezoneAustraliaYancowinna            Timezone = "Australia/Yancowinna"
	TimezoneBrazilAcre                     Timezone = "Brazil/Acre"
	TimezoneBrazilDeNoronha                Timezone = "Brazil/DeNoronha"
	TimezoneBrazilEast                     Timezone = "Brazil/East"
	TimezoneBrazilWest                     Timezone = "Brazil/West"
	TimezoneCanadaAtlantic                 Timezone = "Canada/Atlantic"
	TimezoneCanadaCentral                  Timezone = "Canada/Central"
	TimezoneCanadaEastern                  Timezone = "Canada/Eastern"
	TimezoneCanadaMountain                 Timezone = "Canada/Mountain"
	TimezoneCanadaNewfoundland             Timezone = "Canada/Newfoundland"
	TimezoneCanadaPacific                  Timezone = "Canada/Pacific"
	TimezoneCanadaSaskatchewan             Timezone = "Canada/Saskatchewan"
	TimezoneCanadaYukon                    Timezone = "Canada/Yukon"
	TimezoneChileContinental               Timezone = "Chile/Continental"
	TimezoneChileEasterIsland              Timezone = "Chile/EasterIsland"
	TimezoneCuba                           Timezone = "Cuba"
	TimezoneEgypt                          Timezone = "Egypt"
	TimezoneEire                           Timezone = "Eire"
	TimezoneEtcGMT                         Timezone = "Etc/GMT"
	TimezoneEtcGMTPlus0                    Timezone = "Etc/GMT+0"
	TimezoneEtcGMTPlus1                    Timezone = "Etc/GMT+1"
	TimezoneEtcGMTPlus10                   Timezone = "Etc/GMT+10"
	TimezoneEtcGMTPlus11                   Timezone = "Etc/GMT+11"
	TimezoneEtcGMTPlus12                   Timezone = "Etc/GMT+12"
	TimezoneEtcGMTPlus2                    Timezone = "Etc/GMT+2"
	TimezoneEtcGMTPlus3                    Timezone = "Etc/GMT+3"
	TimezoneEtcGMTPlus4                    Timezone = "Etc/GMT+4"
	TimezoneEtcGMTPlus5                    Timezone = "Etc/GMT+5"
	TimezoneEtcGMTPlus6                    Timezone = "Etc/GMT+6"
	TimezoneEtcGMTPlus7                    Timezone = "Etc/GMT+7"
	TimezoneEtcGMTPlus8                    Timezone = "Etc/GMT+8"
	TimezoneEtcGMTPlus9                    Timezone = "Etc/GMT+9"
	TimezoneEtcGMTMinus0                   Timezone = "Etc/GMT-0"
	TimezoneEtcGMTMinus1                   Timezone = "Etc/GMT-1"
	TimezoneEtcGMTMinus10                  Timezone = "Etc/GMT-10"
	TimezoneEtcGMTMinus11                  Timezone = "Etc/GMT-11"
	TimezoneEtcGMTMinus12                  Timezone = "Etc/GMT-12"
	TimezoneEtcGMTMinus13                  Timezone = "Etc/GMT-13"
	TimezoneEtcGMTMinus14                  Timezone = "Etc/GMT-14"
	TimezoneEtcGMTMinus2                   Timezone = "Etc/GMT-2"
	TimezoneEtcGMTMinus3                   Timezone = "Etc/GMT-3"
	TimezoneEtcGMTMinus4                   Timezone = "Etc/GMT-4"
	TimezoneEtcGMTMinus5                   Timezone = "Etc/GMT-5"
	TimezoneEtcGMTMinus6                   Timezone = "Etc/GMT-6"
	TimezoneEtcGMTMinus7                   Timezone = "Etc/GMT-7"
	TimezoneEtcGMTMinus8                   Timezone = "Etc/GMT-8"
	TimezoneEtcGMTMinus9                   Timezone = "Etc/GMT-9"
	TimezoneEtcGMT0                        Timezone = "Etc/GMT0"
	TimezoneEtcGreenwich                   Timezone = "Etc/Greenwich"
	TimezoneEtcUCT                         Timezone = "Etc/UCT"
	TimezoneEtcUniversal                   Timezone = "Etc/Universal"
	TimezoneEtcUTC                         Timezone = "Etc/UTC"
	TimezoneEtcZulu                        Timezone = "Etc/Zulu"
	TimezoneEuropeAmsterdam                Timezone = "Europe/Amsterdam"
	TimezoneEuropeAndorra                  Timezone = "Europe/Andorra"
	TimezoneEuropeAstrakhan                Timezone = "Europe/Astrakhan"
	TimezoneEuropeAthens                   Timezone = "Europe/Athens"
	TimezoneEuropeBelfast                  Timezone = "Europe/Belfast"
	TimezoneEuropeBelgrade                 Timezone = "Europe/Belgrade"
	TimezoneEuropeBerlin                   Timezone = "Europe/Berlin"
	TimezoneEuropeBratislava               Timezone = "Europe/Bratislava"
	TimezoneEuropeBrussels                 Timezone = "Europe/Brussels"
	TimezoneEuropeBucharest                Timezone = "Europe/Bucharest"
	TimezoneEuropeBudapest                 Timezone = "Europe/Budapest"
	TimezoneEuropeBusingen                 Timezone = "Europe/Busingen"
	TimezoneEuropeChisinau                 Timezone = "Europe/Chisinau"
	TimezoneEuropeCopenhagen               Timezone = "Europe/Copenhagen"
	TimezoneEuropeDublin                   Timezone = "Europe/Dublin"
	TimezoneEuropeGibraltar                Timezone = "Europe/Gibraltar"
	TimezoneEuropeGuernsey                 Timezone = "Europe/Guernsey"
	TimezoneEuropeHelsinki                 Timezone = "Europe/Helsinki"
	TimezoneEuropeIsleOfMan                Timezone = "Europe/Isle_of_Man"
	TimezoneEuropeIstanbul                 Timezone = "Europe/Istanbul"
	TimezoneEuropeJersey                   Timezone = "Europe/Jersey"
	TimezoneEuropeKaliningrad              Timezone = "Europe/Kaliningrad"
	TimezoneEuropeKiev                     Timezone = "Europe/Kiev"
	TimezoneEuropeKirov                    Timezone = "Europe/Kirov"
	TimezoneEuropeKyiv                     Timezone = "Europe/Kyiv"
	TimezoneEuropeLisbon                   Timezone = "Europe/Lisbon"
	TimezoneEuropeLjubljana                Timezone = "Europe/Ljubljana"
	TimezoneEuropeLondon                   Timezone = "Europe/London"
	TimezoneEuropeLuxembourg               Timezone = "Europe/Luxembourg"
	TimezoneEuropeMadrid                   Timezone = "Europe/Madrid"
	TimezoneEuropeMalta                    Timezone = "Europe/Malta"
	TimezoneEuropeMariehamn                Timezone = "Europe/Mariehamn"
	TimezoneEuropeMinsk                    Timezone = "Europe/Minsk"
	TimezoneEuropeMonaco                   Timezone = "Europe/Monaco"
	TimezoneEuropeMoscow                   Timezone = "Europe/Moscow"
	TimezoneEuropeNicosia                  Timezone = "Europe/Nicosia"
	TimezoneEuropeOslo                     Timezone = "Europe/Oslo"
	TimezoneEuropeParis                    Timezone = "Europe/Paris"
	TimezoneEuropePodgorica                Timezone = "Europe/Podgorica"
	TimezoneEuropePrague                   Timezone = "Europe/Prague"
	TimezoneEuropeRiga                     Timezone = "Europe/Riga"
	TimezoneEuropeRome                     Timezone = "Europe/Rome"
	TimezoneEuropeSamara                   Timezone = "Europe/Samara"
	TimezoneEuropeSanMarino                Timezone = "Europe/San_Marino"
	TimezoneEuropeSarajevo                 Timezone = "Europe/Sarajevo"
	TimezoneEuropeSaratov                  Timezone = "Europe/Saratov"
	TimezoneEuropeSimferopol               Timezone = "Europe/Simferopol"
	TimezoneEuropeSkopje                   Timezone = "Europe/Skopje"
	TimezoneEuropeSofia                    Timezone = "Europe/Sofia"
	TimezoneEuropeStockholm                Timezone = "Europe/Stockholm"
	TimezoneEuropeTallinn                  Timezone = "Europe/Tallinn"
	TimezoneEuropeTirane                   Timezone = "Europe/Tirane"
	TimezoneEuropeTiraspol                 Timezone = "Europe/Tiraspol"
	TimezoneEuropeUlyanovsk                Timezone = "Europe/Ulyanovsk"
	TimezoneEuropeUzhgorod                 Timezone = "Europe/Uzhgorod"
	TimezoneEuropeVaduz                    Timezone = "Europe/Vaduz"
	TimezoneEuropeVatican                  Timezone = "Europe/Vatican"
	TimezoneEuropeVienna                   Timezone = "Europe/Vienna"
	TimezoneEuropeVilnius                  Timezone = "Europe/Vilnius"
	TimezoneEuropeVolgograd                Timezone = "Europe/Volgograd"
	TimezoneEuropeWarsaw                   Timezone = "Europe/Warsaw"
	TimezoneEuropeZagreb                   Timezone = "Europe/Zagreb"
	TimezoneEuropeZaporozhye               Timezone = "Europe/Zaporozhye"
	TimezoneEuropeZurich                   Timezone = "Europe/Zurich"
	TimezoneGB                             Timezone = "GB"
	TimezoneGBEire                         Timezone = "GB-Eire"
	TimezoneHongkong                       Timezone = "Hongkong"
	TimezoneIceland                        Timezone = "Iceland"
	TimezoneIndianAntananarivo             Timezone = "Indian/Antananarivo"
	TimezoneIndianChagos                   Timezone = "Indian/Chagos"
	TimezoneIndianChristmas                Timezone = "Indian/Christmas"
	TimezoneIndianCocos                    Timezone = "Indian/Cocos"
	TimezoneIndianComoro                   Timezone = "Indian/Comoro"
	TimezoneIndianKerguelen                Timezone = "Indian/Kerguelen"
	TimezoneIndianMahe                     Timezone = "Indian/Mahe"
	TimezoneIndianMaldives                 Timezone = "Indian/Maldives"
	TimezoneIndianMauritius                Timezone = "Indian/Mauritius"
	TimezoneIndianMayotte                  Timezone = "Indian/Mayotte"
	TimezoneIndianReunion                  Timezone = "Indian/Reunion"
	TimezoneIran                           Timezone = "Iran"
	TimezoneIsrael                         Timezone = "Israel"
	TimezoneJamaica                        Timezone = "Jamaica"
	TimezoneJapan                          Timezone = "Japan"
	TimezoneKwajalein                      Timezone = "Kwajalein"
	TimezoneLibya                          Timezone = "Libya"
	TimezoneMexicoBajaNorte                Timezone = "Mexico/BajaNorte"
	TimezoneMexicoBajaSur                  Timezone = "Mexico/BajaSur"
	TimezoneMexicoGeneral                  Timezone = "Mexico/General"
	TimezoneNavajo                         Timezone = "Navajo"
	TimezoneNZ                             Timezone = "NZ"
	TimezoneNZCHAT                         Timezone = "NZ-CHAT"
	TimezonePacificApia                    Timezone = "Pacific/Apia"
	TimezonePacificAuckland                Timezone = "Pacific/Auckland"
	TimezonePacificBougainville            Timezone = "Pacific/Bougainville"
	TimezonePacificChatham                 Timezone = "Pacific/Chatham"
	TimezonePacificChuuk                   Timezone = "Pacific/Chuuk"
	TimezonePacificEaster                  Timezone = "Pacific/Easter"
	TimezonePacificEfate                   Timezone = "Pacific/Efate"
	TimezonePacificEnderbury               Timezone = "Pacific/Enderbury"
	TimezonePacificFakaofo                 Timezone = "Pacific/Fakaofo"
	TimezonePacificFiji                    Timezone = "Pacific/Fiji"
	TimezonePacificFunafuti                Timezone = "Pacific/Funafuti"
	TimezonePacificGalapagos               Timezone = "Pacific/Galapagos"
	TimezonePacificGambier                 Timezone = "Pacific/Gambier"
	TimezonePacificGuadalcanal             Timezone = "Pacific/Guadalcanal"
	TimezonePacificGuam                    Timezone = "Pacific/Guam"
	TimezonePacificHonolulu                Timezone = "Pacific/Honolulu"
	TimezonePacificJohnston                Timezone = "Pacific/Johnston"
	TimezonePacificKanton                  Timezone = "Pacific/Kanton"
	TimezonePacificKiritimati              Timezone = "Pacific/Kiritimati"
	TimezonePacificKosrae                  Timezone = "Pacific/Kosrae"
	TimezonePacificKwajalein               Timezone = "Pacific/Kwajalein"
	TimezonePacificMajuro                  Timezone = "Pacific/Majuro"
	TimezonePacificMarquesas               Timezone = "Pacific/Marquesas"
	TimezonePacificMidway                  Timezone = "Pacific/Midway"
	TimezonePacificNauru                   Timezone = "Pacific/Nauru"
	TimezonePacificNiue                    Timezone = "Pacific/Niue"
	TimezonePacificNorfolk                 Timezone = "Pacific/Norfolk"
	TimezonePacificNoumea                  Timezone = "Pacific/Noumea"
	TimezonePacificPagoPago                Timezone = "Pacific/Pago_Pago"
	TimezonePacificPalau                   Timezone = "Pacific/Palau"
	TimezonePacificPitcairn                Timezone = "Pacific/Pitcairn"
	TimezonePacificPohnpei                 Timezone = "Pacific/Pohnpei"
	TimezonePacificPonape                  Timezone = "Pacific/Ponape"
	TimezonePacificPortMoresby             Timezone = "Pacific/Port_Moresby"
	TimezonePacificRarotonga               Timezone = "Pacific/Rarotonga"
	TimezonePacificSaipan                  Timezone = "Pacific/Saipan"
	TimezonePacificSamoa                   Timezone = "Pacific/Samoa"
	TimezonePacificTahiti                  Timezone = "Pacific/Tahiti"
	TimezonePacificTarawa                  Timezone = "Pacific/Tarawa"
	TimezonePacificTongatapu               Timezone = "Pacific/Tongatapu"
	TimezonePacificTruk                    Timezone = "Pacific/Truk"
	TimezonePacificWake                    Timezone = "Pacific/Wake"
	TimezonePacificWallis                  Timezone = "Pacific/Wallis"
	TimezonePacificYap                     Timezone = "Pacific/Yap"
	TimezonePoland                         Timezone = "Poland"
	TimezonePortugal                       Timezone = "Portugal"
	TimezonePRC                            Timezone = "PRC"
	TimezoneROC                            Timezone = "ROC"
	TimezoneSingapore                      Timezone = "Singapore"
	TimezoneUSAlaska                       Timezone = "US/Alaska"
	TimezoneUSAleutian                     Timezone = "US/Aleutian"
	TimezoneUSArizona                      Timezone = "US/Arizona"
	TimezoneUSCentral                      Timezone = "US/Central"
	TimezoneUSEastIndiana                  Timezone = "US/East-Indiana"
	TimezoneUSEastern                      Timezone = "US/Eastern"
	TimezoneUSHawaii                       Timezone = "US/Hawaii"
	TimezoneUSIndianaStarke                Timezone = "US/Indiana-Starke"
	TimezoneUSMichigan                     Timezone = "US/Michigan"
	TimezoneUSMountain                     Timezone = "US/Mountain"
	TimezoneUSPacific                      Timezone = "US/Pacific"
	TimezoneUSSamoa                        Timezone = "US/Samoa"
)

func (t Timezone) String() string {
	return string(t)
}

func FromString(t string) (Timezone, error) {
	switch t {
	case "Africa/Abidjan":
		return TimezoneAfricaAbidjan, nil
	case "Africa/Accra":
		return TimezoneAfricaAccra, nil
	case "Africa/Addis_Ababa":
		return TimezoneAfricaAddisAbaba, nil
	case "Africa/Algiers":
		return TimezoneAfricaAlgiers, nil
	case "Africa/Asmara":
		return TimezoneAfricaAsmara, nil
	case "Africa/Bamako":
		return TimezoneAfricaBamako, nil
	case "Africa/Bangui":
		return TimezoneAfricaBangui, nil
	case "Africa/Banjul":
		return TimezoneAfricaBanjul, nil
	case "Africa/Bissau":
		return TimezoneAfricaBissau, nil
	case "Africa/Blantyre":
		return TimezoneAfricaBlantyre, nil
	case "Africa/Brazzaville":
		return TimezoneAfricaBrazzaville, nil
	case "Africa/Bujumbura":
		return TimezoneAfricaBujumbura, nil
	case "Africa/Cairo":
		return TimezoneAfricaCairo, nil
	case "Africa/Casablanca":
		return TimezoneAfricaCasablanca, nil
	case "Africa/Ceuta":
		return TimezoneAfricaCeuta, nil
	case "Africa/Conakry":
		return TimezoneAfricaConakry, nil
	case "Africa/Dakar":
		return TimezoneAfricaDakar, nil
	case "Africa/Dar_es_Salaam":
		return TimezoneAfricaDarEsSalaam, nil
	case "Africa/Djibouti":
		return TimezoneAfricaDjibouti, nil
	case "Africa/Douala":
		return TimezoneAfricaDouala, nil
	case "Africa/El_Aaiun":
		return TimezoneAfricaElAaiun, nil
	case "Africa/Freetown":
		return TimezoneAfricaFreetown, nil
	case "Africa/Gaborone":
		return TimezoneAfricaGaborone, nil
	case "Africa/Harare":
		return TimezoneAfricaHarare, nil
	case "Africa/Johannesburg":
		return TimezoneAfricaJohannesburg, nil
	case "Africa/Juba":
		return TimezoneAfricaJuba, nil
	case "Africa/Kampala":
		return TimezoneAfricaKampala, nil
	case "Africa/Khartoum":
		return TimezoneAfricaKhartoum, nil
	case "Africa/Kigali":
		return TimezoneAfricaKigali, nil
	case "Africa/Kinshasa":
		return TimezoneAfricaKinshasa, nil
	case "Africa/Lagos":
		return TimezoneAfricaLagos, nil
	case "Africa/Libreville":
		return TimezoneAfricaLibreville, nil
	case "Africa/Lome":
		return TimezoneAfricaLome, nil
	case "Africa/Luanda":
		return TimezoneAfricaLuanda, nil
	case "Africa/Lubumbashi":
		return TimezoneAfricaLubumbashi, nil
	case "Africa/Lusaka":
		return TimezoneAfricaLusaka, nil
	case "Africa/Malabo":
		return TimezoneAfricaMalabo, nil
	case "Africa/Maputo":
		return TimezoneAfricaMaputo, nil
	case "Africa/Maseru":
		return TimezoneAfricaMaseru, nil
	case "Africa/Mbabane":
		return TimezoneAfricaMbabane, nil
	case "Africa/Mogadishu":
		return TimezoneAfricaMogadishu, nil
	case "Africa/Monrovia":
		return TimezoneAfricaMonrovia, nil
	case "Africa/Nairobi":
		return TimezoneAfricaNairobi, nil
	case "Africa/Ndjamena":
		return TimezoneAfricaNdjamena, nil
	case "Africa/Niamey":
		return TimezoneAfricaNiamey, nil
	case "Africa/Nouakchott":
		return TimezoneAfricaNouakchott, nil
	case "Africa/Ouagadougou":
		return TimezoneAfricaOuagadougou, nil
	case "Africa/Porto-Novo":
		return TimezoneAfricaPortoNovo, nil
	case "Africa/Sao_Tome":
		return TimezoneAfricaSaoTome, nil
	case "Africa/Tripoli":
		return TimezoneAfricaTripoli, nil
	case "Africa/Tunis":
		return TimezoneAfricaTunis, nil
	case "Africa/Windhoek":
		return TimezoneAfricaWindhoek, nil
	case "America/Adak":
		return TimezoneAmericaAdak, nil
	case "America/Anchorage":
		return TimezoneAmericaAnchorage, nil
	case "America/Anguilla":
		return TimezoneAmericaAnguilla, nil
	case "America/Antigua":
		return TimezoneAmericaAntigua, nil
	case "America/Araguaina":
		return TimezoneAmericaAraguaina, nil
	case "America/Argentina/Buenos_Aires":
		return TimezoneAmericaArgentinaBuenosAires, nil
	case "America/Argentina/Catamarca":
		return TimezoneAmericaArgentinaCatamarca, nil
	case "America/Argentina/Cordoba":
		return TimezoneAmericaArgentinaCordoba, nil
	case "America/Argentina/Jujuy":
		return TimezoneAmericaArgentinaJujuy, nil
	case "America/Argentina/La_Rioja":
		return TimezoneAmericaArgentinaLaRioja, nil
	case "America/Argentina/Mendoza":
		return TimezoneAmericaArgentinaMendoza, nil
	case "America/Argentina/Rio_Gallegos":
		return TimezoneAmericaArgentinaRioGallegos, nil
	case "America/Argentina/Salta":
		return TimezoneAmericaArgentinaSalta, nil
	case "America/Argentina/San_Juan":
		return TimezoneAmericaArgentinaSanJuan, nil
	case "America/Argentina/San_Luis":
		return TimezoneAmericaArgentinaSanLuis, nil
	case "America/Argentina/Tucuman":
		return TimezoneAmericaArgentinaTucuman, nil
	case "America/Argentina/Ushuaia":
		return TimezoneAmericaArgentinaUshuaia, nil
	case "America/Aruba":
		return TimezoneAmericaAruba, nil
	case "America/Asuncion":
		return TimezoneAmericaAsuncion, nil
	case "America/Atikokan":
		return TimezoneAmericaAtikokan, nil
	case "America/Bahia":
		return TimezoneAmericaBahia, nil
	case "America/Bahia_Banderas":
		return TimezoneAmericaBahiaBanderas, nil
	case "America/Barbados":
		return TimezoneAmericaBarbados, nil
	case "America/Belem":
		return TimezoneAmericaBelem, nil
	case "America/Belize":
		return TimezoneAmericaBelize, nil
	case "America/Blanc-Sablon":
		return TimezoneAmericaBlancSablon, nil
	case "America/Boa_Vista":
		return TimezoneAmericaBoaVista, nil
	case "America/Bogota":
		return TimezoneAmericaBogota, nil
	case "America/Boise":
		return TimezoneAmericaBoise, nil
	case "America/Cambridge_Bay":
		return TimezoneAmericaCambridgeBay, nil
	case "America/Campo_Grande":
		return TimezoneAmericaCampoGrande, nil
	case "America/Cancun":
		return TimezoneAmericaCancun, nil
	case "America/Caracas":
		return TimezoneAmericaCaracas, nil
	case "America/Cayenne":
		return TimezoneAmericaCayenne, nil
	case "America/Cayman":
		return TimezoneAmericaCayman, nil
	case "America/Chicago":
		return TimezoneAmericaChicago, nil
	case "America/Chihuahua":
		return TimezoneAmericaChihuahua, nil
	case "America/Costa_Rica":
		return TimezoneAmericaCostaRica, nil
	case "America/Creston":
		return TimezoneAmericaCreston, nil
	case "America/Cuiaba":
		return TimezoneAmericaCuiaba, nil
	case "America/Curacao":
		return TimezoneAmericaCuracao, nil
	case "America/Danmarkshavn":
		return TimezoneAmericaDanmarkshavn, nil
	case "America/Dawson":
		return TimezoneAmericaDawson, nil
	case "America/Dawson_Creek":
		return TimezoneAmericaDawsonCreek, nil
	case "America/Denver":
		return TimezoneAmericaDenver, nil
	case "America/Detroit":
		return TimezoneAmericaDetroit, nil
	case "America/Dominica":
		return TimezoneAmericaDominica, nil
	case "America/Edmonton":
		return TimezoneAmericaEdmonton, nil
	case "America/Eirunepe":
		return TimezoneAmericaEirunepe, nil
	case "America/El_Salvador":
		return TimezoneAmericaElSalvador, nil
	case "America/Fort_Nelson":
		return TimezoneAmericaFortNelson, nil
	case "America/Fortaleza":
		return TimezoneAmericaFortaleza, nil
	case "America/Glace_Bay":
		return TimezoneAmericaGlaceBay, nil
	case "America/Godthab":
		return TimezoneAmericaGodthab, nil
	case "America/Goose_Bay":
		return TimezoneAmericaGooseBay, nil
	case "America/Grand_Turk":
		return TimezoneAmericaGrandTurk, nil
	case "America/Grenada":
		return TimezoneAmericaGrenada, nil
	case "America/Guadeloupe":
		return TimezoneAmericaGuadeloupe, nil
	case "America/Guatemala":
		return TimezoneAmericaGuatemala, nil
	case "America/Guayaquil":
		return TimezoneAmericaGuayaquil, nil
	case "America/Guyana":
		return TimezoneAmericaGuyana, nil
	case "America/Halifax":
		return TimezoneAmericaHalifax, nil
	case "America/Havana":
		return TimezoneAmericaHavana, nil
	case "America/Hermosillo":
		return TimezoneAmericaHermosillo, nil
	case "America/Indiana/Indianapolis":
		return TimezoneAmericaIndianaIndianapolis, nil
	case "America/Indiana/Knox":
		return TimezoneAmericaIndianaKnox, nil
	case "America/Indiana/Marengo":
		return TimezoneAmericaIndianaMarengo, nil
	case "America/Indiana/Petersburg":
		return TimezoneAmericaIndianaPetersburg, nil
	case "America/Indiana/Tell_City":
		return TimezoneAmericaIndianaTellCity, nil
	case "America/Indiana/Vevay":
		return TimezoneAmericaIndianaVevay, nil
	case "America/Indiana/Vincennes":
		return TimezoneAmericaIndianaVincennes, nil
	case "America/Indiana/Winamac":
		return TimezoneAmericaIndianaWinamac, nil
	case "America/Inuvik":
		return TimezoneAmericaInuvik, nil
	case "America/Iqaluit":
		return TimezoneAmericaIqaluit, nil
	case "America/Jamaica":
		return TimezoneAmericaJamaica, nil
	case "America/Juneau":
		return TimezoneAmericaJuneau, nil
	case "America/Kentucky/Louisville":
		return TimezoneAmericaKentuckyLouisville, nil
	case "America/Kentucky/Monticello":
		return TimezoneAmericaKentuckyMonticello, nil
	case "America/Kralendijk":
		return TimezoneAmericaKralendijk, nil
	case "America/La_Paz":
		return TimezoneAmericaLaPaz, nil
	case "America/Lima":
		return TimezoneAmericaLima, nil
	case "America/Los_Angeles":
		return TimezoneAmericaLosAngeles, nil
	case "America/Lower_Princes":
		return TimezoneAmericaLowerPrinces, nil
	case "America/Maceio":
		return TimezoneAmericaMaceio, nil
	case "America/Managua":
		return TimezoneAmericaManagua, nil
	case "America/Manaus":
		return TimezoneAmericaManaus, nil
	case "America/Marigot":
		return TimezoneAmericaMarigot, nil
	case "America/Martinique":
		return TimezoneAmericaMartinique, nil
	case "America/Matamoros":
		return TimezoneAmericaMatamoros, nil
	case "America/Mazatlan":
		return TimezoneAmericaMazatlan, nil
	case "America/Menominee":
		return TimezoneAmericaMenominee, nil
	case "America/Merida":
		return TimezoneAmericaMerida, nil
	case "America/Metlakatla":
		return TimezoneAmericaMetlakatla, nil
	case "America/Mexico_City":
		return TimezoneAmericaMexicoCity, nil
	case "America/Miquelon":
		return TimezoneAmericaMiquelon, nil
	case "America/Moncton":
		return TimezoneAmericaMoncton, nil
	case "America/Monterrey":
		return TimezoneAmericaMonterrey, nil
	case "America/Montevideo":
		return TimezoneAmericaMontevideo, nil
	case "America/Montserrat":
		return TimezoneAmericaMontserrat, nil
	case "America/Nassau":
		return TimezoneAmericaNassau, nil
	case "America/New_York":
		return TimezoneAmericaNewYork, nil
	case "America/Nipigon":
		return TimezoneAmericaNipigon, nil
	case "America/Nome":
		return TimezoneAmericaNome, nil
	case "America/Noronha":
		return TimezoneAmericaNoronha, nil
	case "America/North_Dakota/Beulah":
		return TimezoneAmericaNorthDakotaBeulah, nil
	case "America/North_Dakota/Center":
		return TimezoneAmericaNorthDakotaCenter, nil
	case "America/North_Dakota/New_Salem":
		return TimezoneAmericaNorthDakotaNewSalem, nil
	case "America/Ojinaga":
		return TimezoneAmericaOjinaga, nil
	case "America/Panama":
		return TimezoneAmericaPanama, nil
	case "America/Pangnirtung":
		return TimezoneAmericaPangnirtung, nil
	case "America/Paramaribo":
		return TimezoneAmericaParamaribo, nil
	case "America/Phoenix":
		return TimezoneAmericaPhoenix, nil
	case "America/Port-au-Prince":
		return TimezoneAmericaPortAuPrince, nil
	case "America/Port_of_Spain":
		return TimezoneAmericaPortOfSpain, nil
	case "America/Porto_Velho":
		return TimezoneAmericaPortoVelho, nil
	case "America/Puerto_Rico":
		return TimezoneAmericaPuertoRico, nil
	case "America/Punta_Arenas":
		return TimezoneAmericaPuntaArenas, nil
	case "America/Rainy_River":
		return TimezoneAmericaRainyRiver, nil
	case "America/Rankin_Inlet":
		return TimezoneAmericaRankinInlet, nil
	case "America/Recife":
		return TimezoneAmericaRecife, nil
	case "America/Regina":
		return TimezoneAmericaRegina, nil
	case "America/Resolute":
		return TimezoneAmericaResolute, nil
	case "America/Rio_Branco":
		return TimezoneAmericaRioBranco, nil
	case "America/Santarem":
		return TimezoneAmericaSantarem, nil
	case "America/Santiago":
		return TimezoneAmericaSantiago, nil
	case "America/Santo_Domingo":
		return TimezoneAmericaSantoDomingo, nil
	case "America/Sao_Paulo":
		return TimezoneAmericaSaoPaulo, nil
	case "America/Scoresbysund":
		return TimezoneAmericaScoresbysund, nil
	case "America/Sitka":
		return TimezoneAmericaSitka, nil
	case "America/St_Barthelemy":
		return TimezoneAmericaStBarthelemy, nil
	case "America/St_Johns":
		return TimezoneAmericaStJohns, nil
	case "America/St_Kitts":
		return TimezoneAmericaStKitts, nil
	case "America/St_Lucia":
		return TimezoneAmericaStLucia, nil
	case "America/St_Thomas":
		return TimezoneAmericaStThomas, nil
	case "America/St_Vincent":
		return TimezoneAmericaStVincent, nil
	case "America/Swift_Current":
		return TimezoneAmericaSwiftCurrent, nil
	case "America/Tegucigalpa":
		return TimezoneAmericaTegucigalpa, nil
	case "America/Thule":
		return TimezoneAmericaThule, nil
	case "America/Thunder_Bay":
		return TimezoneAmericaThunderBay, nil
	case "America/Tijuana":
		return TimezoneAmericaTijuana, nil
	case "America/Toronto":
		return TimezoneAmericaToronto, nil
	case "America/Tortola":
		return TimezoneAmericaTortola, nil
	case "America/Vancouver":
		return TimezoneAmericaVancouver, nil
	case "America/Whitehorse":
		return TimezoneAmericaWhitehorse, nil
	case "America/Winnipeg":
		return TimezoneAmericaWinnipeg, nil
	case "America/Yakutat":
		return TimezoneAmericaYakutat, nil
	case "America/Yellowknife":
		return TimezoneAmericaYellowknife, nil
	case "Antarctica/Casey":
		return TimezoneAntarcticaCasey, nil
	case "Antarctica/Davis":
		return TimezoneAntarcticaDavis, nil
	case "Antarctica/DumontDUrville":
		return TimezoneAntarcticaDumontDUrville, nil
	case "Antarctica/Macquarie":
		return TimezoneAntarcticaMacquarie, nil
	case "Antarctica/Mawson":
		return TimezoneAntarcticaMawson, nil
	case "Antarctica/McMurdo":
		return TimezoneAntarcticaMcMurdo, nil
	case "Antarctica/Palmer":
		return TimezoneAntarcticaPalmer, nil
	case "Antarctica/Rothera":
		return TimezoneAntarcticaRothera, nil
	case "Antarctica/Syowa":
		return TimezoneAntarcticaSyowa, nil
	case "Antarctica/Troll":
		return TimezoneAntarcticaTroll, nil
	case "Antarctica/Vostok":
		return TimezoneAntarcticaVostok, nil
	case "Arctic/Longyearbyen":
		return TimezoneArcticLongyearbyen, nil
	case "Asia/Aden":
		return TimezoneAsiaAden, nil
	case "Asia/Almaty":
		return TimezoneAsiaAlmaty, nil
	case "Asia/Amman":
		return TimezoneAsiaAmman, nil
	case "Asia/Anadyr":
		return TimezoneAsiaAnadyr, nil
	case "Asia/Aqtau":
		return TimezoneAsiaAqtau, nil
	case "Asia/Aqtobe":
		return TimezoneAsiaAqtobe, nil
	case "Asia/Ashgabat":
		return TimezoneAsiaAshgabat, nil
	case "Asia/Atyrau":
		return TimezoneAsiaAtyrau, nil
	case "Asia/Baghdad":
		return TimezoneAsiaBaghdad, nil
	case "Asia/Bahrain":
		return TimezoneAsiaBahrain, nil
	case "Asia/Baku":
		return TimezoneAsiaBaku, nil
	case "Asia/Bangkok":
		return TimezoneAsiaBangkok, nil
	case "Asia/Barnaul":
		return TimezoneAsiaBarnaul, nil
	case "Asia/Beirut":
		return TimezoneAsiaBeirut, nil
	case "Asia/Bishkek":
		return TimezoneAsiaBishkek, nil
	case "Asia/Brunei":
		return TimezoneAsiaBrunei, nil
	case "Asia/Chita":
		return TimezoneAsiaChita, nil
	case "Asia/Choibalsan":
		return TimezoneAsiaChoibalsan, nil
	case "Asia/Colombo":
		return TimezoneAsiaColombo, nil
	case "Asia/Damascus":
		return TimezoneAsiaDamascus, nil
	case "Asia/Dhaka":
		return TimezoneAsiaDhaka, nil
	case "Asia/Dili":
		return TimezoneAsiaDili, nil
	case "Asia/Dubai":
		return TimezoneAsiaDubai, nil
	case "Asia/Dushanbe":
		return TimezoneAsiaDushanbe, nil
	case "Asia/Famagusta":
		return TimezoneAsiaFamagusta, nil
	case "Asia/Gaza":
		return TimezoneAsiaGaza, nil
	case "Asia/Hebron":
		return TimezoneAsiaHebron, nil
	case "Asia/Ho_Chi_Minh":
		return TimezoneAsiaHoChiMinh, nil
	case "Asia/Hong_Kong":
		return TimezoneAsiaHongKong, nil
	case "Asia/Hovd":
		return TimezoneAsiaHovd, nil
	case "Asia/Irkutsk":
		return TimezoneAsiaIrkutsk, nil
	case "Asia/Jakarta":
		return TimezoneAsiaJakarta, nil
	case "Asia/Jayapura":
		return TimezoneAsiaJayapura, nil
	case "Asia/Jerusalem":
		return TimezoneAsiaJerusalem, nil
	case "Asia/Kabul":
		return TimezoneAsiaKabul, nil
	case "Asia/Kamchatka":
		return TimezoneAsiaKamchatka, nil
	case "Asia/Karachi":
		return TimezoneAsiaKarachi, nil
	case "Asia/Kathmandu":
		return TimezoneAsiaKathmandu, nil
	case "Asia/Khandyga":
		return TimezoneAsiaKhandyga, nil
	case "Asia/Kolkata":
		return TimezoneAsiaKolkata, nil
	case "Asia/Krasnoyarsk":
		return TimezoneAsiaKrasnoyarsk, nil
	case "Asia/Kuala_Lumpur":
		return TimezoneAsiaKualaLumpur, nil
	case "Asia/Kuching":
		return TimezoneAsiaKuching, nil
	case "Asia/Macau":
		return TimezoneAsiaMacau, nil
	case "Asia/Magadan":
		return TimezoneAsiaMagadan, nil
	case "Asia/Makassar":
		return TimezoneAsiaMakassar, nil
	case "Asia/Manila":
		return TimezoneAsiaManila, nil
	case "Asia/Nicosia":
		return TimezoneAsiaNicosia, nil
	case "Asia/Novokuznetsk":
		return TimezoneAsiaNovokuznetsk, nil
	case "Asia/Novosibirsk":
		return TimezoneAsiaNovosibirsk, nil
	case "Asia/Omsk":
		return TimezoneAsiaOmsk, nil
	case "Asia/Oral":
		return TimezoneAsiaOral, nil
	case "Asia/Pontianak":
		return TimezoneAsiaPontianak, nil
	case "Asia/Pyongyang":
		return TimezoneAsiaPyongyang, nil
	case "Asia/Qatar":
		return TimezoneAsiaQatar, nil
	case "Asia/Qostanay":
		return TimezoneAsiaQostanay, nil
	case "Asia/Qyzylorda":
		return TimezoneAsiaQyzylorda, nil
	case "Asia/Riyadh":
		return TimezoneAsiaRiyadh, nil
	case "Asia/Sakhalin":
		return TimezoneAsiaSakhalin, nil
	case "Asia/Samarkand":
		return TimezoneAsiaSamarkand, nil
	case "Asia/Seoul":
		return TimezoneAsiaSeoul, nil
	case "Asia/Shanghai":
		return TimezoneAsiaShanghai, nil
	case "Asia/Singapore":
		return TimezoneAsiaSingapore, nil
	case "Asia/Srednekolymsk":
		return TimezoneAsiaSrednekolymsk, nil
	case "Asia/Taipei":
		return TimezoneAsiaTaipei, nil
	case "Asia/Tashkent":
		return TimezoneAsiaTashkent, nil
	case "Asia/Tbilisi":
		return TimezoneAsiaTbilisi, nil
	case "Asia/Tehran":
		return TimezoneAsiaTehran, nil
	case "Asia/Thimphu":
		return TimezoneAsiaThimphu, nil
	case "Asia/Tokyo":
		return TimezoneAsiaTokyo, nil
	case "Asia/Tomsk":
		return TimezoneAsiaTomsk, nil
	case "Asia/Ulaanbaatar":
		return TimezoneAsiaUlaanbaatar, nil
	case "Asia/Urumqi":
		return TimezoneAsiaUrumqi, nil
	case "Asia/Ust-Nera":
		return TimezoneAsiaUstNera, nil
	case "Asia/Vladivostok":
		return TimezoneAsiaVladivostok, nil
	case "Asia/Yakutsk":
		return TimezoneAsiaYakutsk, nil
	case "Asia/Yangon":
		return TimezoneAsiaYangon, nil
	case "Asia/Yekaterinburg":
		return TimezoneAsiaYekaterinburg, nil
	case "Asia/Yerevan":
		return TimezoneAsiaYerevan, nil
	case "Atlantic/Azores":
		return TimezoneAtlanticAzores, nil
	case "Atlantic/Bermuda":
		return TimezoneAtlanticBermuda, nil
	case "Atlantic/Canary":
		return TimezoneAtlanticCanary, nil
	case "Atlantic/Cape_Verde":
		return TimezoneAtlanticCapeVerde, nil
	case "Atlantic/Faroe":
		return TimezoneAtlanticFaroe, nil
	case "Atlantic/Madeira":
		return TimezoneAtlanticMadeira, nil
	case "Atlantic/Reykjavik":
		return TimezoneAtlanticReykjavik, nil
	case "Atlantic/South_Georgia":
		return TimezoneAtlanticSouthGeorgia, nil
	case "Atlantic/Stanley":
		return TimezoneAtlanticStanley, nil
	case "Australia/Adelaide":
		return TimezoneAustraliaAdelaide, nil
	case "Australia/Brisbane":
		return TimezoneAustraliaBrisbane, nil
	case "Australia/Broken_Hill":
		return TimezoneAustraliaBrokenHill, nil
	case "Australia/Currie":
		return TimezoneAustraliaCurrie, nil
	case "Australia/Darwin":
		return TimezoneAustraliaDarwin, nil
	case "Australia/Eucla":
		return TimezoneAustraliaEucla, nil
	case "Australia/Hobart":
		return TimezoneAustraliaHobart, nil
	case "Australia/Lindeman":
		return TimezoneAustraliaLindeman, nil
	case "Australia/Lord_Howe":
		return TimezoneAustraliaLordHowe, nil
	case "Australia/Melbourne":
		return TimezoneAustraliaMelbourne, nil
	case "Australia/Perth":
		return TimezoneAustraliaPerth, nil
	case "Australia/Sydney":
		return TimezoneAustraliaSydney, nil
	case "Australia/Tasmania":
		return TimezoneAustraliaTasmania, nil
	case "Australia/Victoria":
		return TimezoneAustraliaVictoria, nil
	case "Australia/West":
		return TimezoneAustraliaWest, nil
	case "Australia/Yancowinna":
		return TimezoneAustraliaYancowinna, nil
	case "Brazil/Acre":
		return TimezoneBrazilAcre, nil
	case "Brazil/DeNoronha":
		return TimezoneBrazilDeNoronha, nil
	case "Brazil/East":
		return TimezoneBrazilEast, nil
	case "Brazil/West":
		return TimezoneBrazilWest, nil
	case "Canada/Atlantic":
		return TimezoneCanadaAtlantic, nil
	case "Canada/Central":
		return TimezoneCanadaCentral, nil
	case "Canada/Eastern":
		return TimezoneCanadaEastern, nil
	case "Canada/Mountain":
		return TimezoneCanadaMountain, nil
	case "Canada/Newfoundland":
		return TimezoneCanadaNewfoundland, nil
	case "Canada/Pacific":
		return TimezoneCanadaPacific, nil
	case "Canada/Saskatchewan":
		return TimezoneCanadaSaskatchewan, nil
	case "Canada/Yukon":
		return TimezoneCanadaYukon, nil
	case "Chile/Continental":
		return TimezoneChileContinental, nil
	case "Chile/EasterIsland":
		return TimezoneChileEasterIsland, nil
	case "Cuba":
		return TimezoneCuba, nil
	case "Egypt":
		return TimezoneEgypt, nil
	case "Eire":
		return TimezoneEire, nil
	case "Etc/GMT":
		return TimezoneEtcGMT, nil
	case "Etc/GMT+0":
		return TimezoneEtcGMTPlus0, nil
	case "Etc/GMT+1":
		return TimezoneEtcGMTPlus1, nil
	case "Etc/GMT+10":
		return TimezoneEtcGMTPlus10, nil
	case "Etc/GMT+11":
		return TimezoneEtcGMTPlus11, nil
	case "Etc/GMT+12":
		return TimezoneEtcGMTPlus12, nil
	case "Etc/GMT+2":
		return TimezoneEtcGMTPlus2, nil
	case "Etc/GMT+3":
		return TimezoneEtcGMTPlus3, nil
	case "Etc/GMT+4":
		return TimezoneEtcGMTPlus4, nil
	case "Etc/GMT+5":
		return TimezoneEtcGMTPlus5, nil
	case "Etc/GMT+6":
		return TimezoneEtcGMTPlus6, nil
	case "Etc/GMT+7":
		return TimezoneEtcGMTPlus7, nil
	case "Etc/GMT+8":
		return TimezoneEtcGMTPlus8, nil
	case "Etc/GMT+9":
		return TimezoneEtcGMTPlus9, nil
	case "Etc/GMT-0":
		return TimezoneEtcGMTMinus0, nil
	case "Etc/GMT-1":
		return TimezoneEtcGMTMinus1, nil
	case "Etc/GMT-10":
		return TimezoneEtcGMTMinus10, nil
	case "Etc/GMT-11":
		return TimezoneEtcGMTMinus11, nil
	case "Etc/GMT-12":
		return TimezoneEtcGMTMinus12, nil
	case "Etc/GMT-13":
		return TimezoneEtcGMTMinus13, nil
	case "Etc/GMT-14":
		return TimezoneEtcGMTMinus14, nil
	case "Etc/GMT-2":
		return TimezoneEtcGMTMinus2, nil
	case "Etc/GMT-3":
		return TimezoneEtcGMTMinus3, nil
	case "Etc/GMT-4":
		return TimezoneEtcGMTMinus4, nil
	case "Etc/GMT-5":
		return TimezoneEtcGMTMinus5, nil
	case "Etc/GMT-6":
		return TimezoneEtcGMTMinus6, nil
	case "Etc/GMT-7":
		return TimezoneEtcGMTMinus7, nil
	case "Etc/GMT-8":
		return TimezoneEtcGMTMinus8, nil
	case "Etc/GMT-9":
		return TimezoneEtcGMTMinus9, nil
	case "Etc/GMT0":
		return TimezoneEtcGMT0, nil
	case "Etc/Greenwich":
		return TimezoneEtcGreenwich, nil
	case "Etc/UCT":
		return TimezoneEtcUCT, nil
	case "Etc/UTC":
		return TimezoneEtcUTC, nil
	case "Etc/Universal":
		return TimezoneEtcUniversal, nil
	case "Etc/Zulu":
		return TimezoneEtcZulu, nil
	case "Europe/Amsterdam":
		return TimezoneEuropeAmsterdam, nil
	case "Europe/Andorra":
		return TimezoneEuropeAndorra, nil
	case "Europe/Astrakhan":
		return TimezoneEuropeAstrakhan, nil
	case "Europe/Athens":
		return TimezoneEuropeAthens, nil
	case "Europe/Belgrade":
		return TimezoneEuropeBelgrade, nil
	case "Europe/Berlin":
		return TimezoneEuropeBerlin, nil
	case "Europe/Bratislava":
		return TimezoneEuropeBratislava, nil
	case "Europe/Brussels":
		return TimezoneEuropeBrussels, nil
	case "Europe/Bucharest":
		return TimezoneEuropeBucharest, nil
	case "Europe/Budapest":
		return TimezoneEuropeBudapest, nil
	case "Europe/Busingen":
		return TimezoneEuropeBusingen, nil
	case "Europe/Chisinau":
		return TimezoneEuropeChisinau, nil
	case "Europe/Copenhagen":
		return TimezoneEuropeCopenhagen, nil
	case "Europe/Dublin":
		return TimezoneEuropeDublin, nil
	case "Europe/Gibraltar":
		return TimezoneEuropeGibraltar, nil
	case "Europe/Guernsey":
		return TimezoneEuropeGuernsey, nil
	case "Europe/Helsinki":
		return TimezoneEuropeHelsinki, nil
	case "Europe/Isle_of_Man":
		return TimezoneEuropeIsleOfMan, nil
	case "Europe/Istanbul":
		return TimezoneEuropeIstanbul, nil
	case "Europe/Jersey":
		return TimezoneEuropeJersey, nil
	case "Europe/Kaliningrad":
		return TimezoneEuropeKaliningrad, nil
	case "Europe/Kiev":
		return TimezoneEuropeKiev, nil
	case "Europe/Lisbon":
		return TimezoneEuropeLisbon, nil
	case "Europe/Ljubljana":
		return TimezoneEuropeLjubljana, nil
	case "Europe/London":
		return TimezoneEuropeLondon, nil
	case "Europe/Luxembourg":
		return TimezoneEuropeLuxembourg, nil
	case "Europe/Madrid":
		return TimezoneEuropeMadrid, nil
	case "Europe/Malta":
		return TimezoneEuropeMalta, nil
	case "Europe/Mariehamn":
		return TimezoneEuropeMariehamn, nil
	case "Europe/Minsk":
		return TimezoneEuropeMinsk, nil
	case "Europe/Monaco":
		return TimezoneEuropeMonaco, nil
	case "Europe/Moscow":
		return TimezoneEuropeMoscow, nil
	case "Europe/Nicosia":
		return TimezoneEuropeNicosia, nil
	case "Europe/Oslo":
		return TimezoneEuropeOslo, nil
	case "Europe/Paris":
		return TimezoneEuropeParis, nil
	case "Europe/Podgorica":
		return TimezoneEuropePodgorica, nil
	case "Europe/Prague":
		return TimezoneEuropePrague, nil
	case "Europe/Riga":
		return TimezoneEuropeRiga, nil
	case "Europe/Rome":
		return TimezoneEuropeRome, nil
	case "Europe/Samara":
		return TimezoneEuropeSamara, nil
	case "Europe/San_Marino":
		return TimezoneEuropeSanMarino, nil
	case "Europe/Sarajevo":
		return TimezoneEuropeSarajevo, nil
	case "Europe/Simferopol":
		return TimezoneEuropeSimferopol, nil
	case "Europe/Skopje":
		return TimezoneEuropeSkopje, nil
	case "Europe/Sofia":
		return TimezoneEuropeSofia, nil
	case "Europe/Stockholm":
		return TimezoneEuropeStockholm, nil
	case "Europe/Tallinn":
		return TimezoneEuropeTallinn, nil
	case "Europe/Tirane":
		return TimezoneEuropeTirane, nil
	case "Europe/Ulyanovsk":
		return TimezoneEuropeUlyanovsk, nil
	case "Europe/Uzhgorod":
		return TimezoneEuropeUzhgorod, nil
	case "Europe/Vaduz":
		return TimezoneEuropeVaduz, nil
	case "Europe/Vatican":
		return TimezoneEuropeVatican, nil
	case "Europe/Vienna":
		return TimezoneEuropeVienna, nil
	case "Europe/Vilnius":
		return TimezoneEuropeVilnius, nil
	case "Europe/Volgograd":
		return TimezoneEuropeVolgograd, nil
	case "Europe/Warsaw":
		return TimezoneEuropeWarsaw, nil
	case "Europe/Zagreb":
		return TimezoneEuropeZagreb, nil
	case "Europe/Zaporozhye":
		return TimezoneEuropeZaporozhye, nil
	case "Europe/Zurich":
		return TimezoneEuropeZurich, nil
	case "GB":
		return TimezoneGB, nil
	case "GB-Eire":
		return TimezoneGBEire, nil
	case "Hongkong":
		return TimezoneHongkong, nil
	case "Iceland":
		return TimezoneIceland, nil
	case "Indian/Antananarivo":
		return TimezoneIndianAntananarivo, nil
	case "Indian/Chagos":
		return TimezoneIndianChagos, nil
	case "Indian/Christmas":
		return TimezoneIndianChristmas, nil
	case "Indian/Cocos":
		return TimezoneIndianCocos, nil
	case "Indian/Comoro":
		return TimezoneIndianComoro, nil
	case "Indian/Kerguelen":
		return TimezoneIndianKerguelen, nil
	case "Indian/Mahe":
		return TimezoneIndianMahe, nil
	case "Indian/Maldives":
		return TimezoneIndianMaldives, nil
	case "Indian/Mauritius":
		return TimezoneIndianMauritius, nil
	case "Indian/Mayotte":
		return TimezoneIndianMayotte, nil
	case "Indian/Reunion":
		return TimezoneIndianReunion, nil
	case "Iran":
		return TimezoneIran, nil
	case "Israel":
		return TimezoneIsrael, nil
	case "Jamaica":
		return TimezoneJamaica, nil
	case "Japan":
		return TimezoneJapan, nil
	case "Kwajalein":
		return TimezoneKwajalein, nil
	case "Libya":
		return TimezoneLibya, nil
	case "Mexico/BajaNorte":
		return TimezoneMexicoBajaNorte, nil
	case "Mexico/BajaSur":
		return TimezoneMexicoBajaSur, nil
	case "Mexico/General":
		return TimezoneMexicoGeneral, nil
	case "Navajo":
		return TimezoneNavajo, nil
	case "NZ":
		return TimezoneNZ, nil
	case "NZ-CHAT":
		return TimezoneNZCHAT, nil
	case "Pacific/Apia":
		return TimezonePacificApia, nil
	case "Pacific/Auckland":
		return TimezonePacificAuckland, nil
	case "Pacific/Bougainville":
		return TimezonePacificBougainville, nil
	case "Pacific/Chatham":
		return TimezonePacificChatham, nil
	case "Pacific/Chuuk":
		return TimezonePacificChuuk, nil
	case "Pacific/Easter":
		return TimezonePacificEaster, nil
	case "Pacific/Efate":
		return TimezonePacificEfate, nil
	case "Pacific/Enderbury":
		return TimezonePacificEnderbury, nil
	case "Pacific/Fakaofo":
		return TimezonePacificFakaofo, nil
	case "Pacific/Fiji":
		return TimezonePacificFiji, nil
	case "Pacific/Funafuti":
		return TimezonePacificFunafuti, nil
	case "Pacific/Galapagos":
		return TimezonePacificGalapagos, nil
	case "Pacific/Gambier":
		return TimezonePacificGambier, nil
	case "Pacific/Guadalcanal":
		return TimezonePacificGuadalcanal, nil
	case "Pacific/Guam":
		return TimezonePacificGuam, nil
	case "Pacific/Honolulu":
		return TimezonePacificHonolulu, nil
	case "Pacific/Johnston":
		return TimezonePacificJohnston, nil
	case "Pacific/Kiritimati":
		return TimezonePacificKiritimati, nil
	case "Pacific/Kosrae":
		return TimezonePacificKosrae, nil
	case "Pacific/Kwajalein":
		return TimezonePacificKwajalein, nil
	case "Pacific/Majuro":
		return TimezonePacificMajuro, nil
	case "Pacific/Marquesas":
		return TimezonePacificMarquesas, nil
	case "Pacific/Midway":
		return TimezonePacificMidway, nil
	case "Pacific/Nauru":
		return TimezonePacificNauru, nil
	case "Pacific/Niue":
		return TimezonePacificNiue, nil
	case "Pacific/Norfolk":
		return TimezonePacificNorfolk, nil
	case "Pacific/Noumea":
		return TimezonePacificNoumea, nil
	case "Pacific/Pago_Pago":
		return TimezonePacificPagoPago, nil
	case "Pacific/Palau":
		return TimezonePacificPalau, nil
	case "Pacific/Pitcairn":
		return TimezonePacificPitcairn, nil
	case "Pacific/Pohnpei":
		return TimezonePacificPohnpei, nil
	case "Pacific/Ponape":
		return TimezonePacificPonape, nil
	case "Pacific/Port_Moresby":
		return TimezonePacificPortMoresby, nil
	case "Pacific/Rarotonga":
		return TimezonePacificRarotonga, nil
	case "Pacific/Saipan":
		return TimezonePacificSaipan, nil
	case "Pacific/Samoa":
		return TimezonePacificSamoa, nil
	case "Pacific/Tahiti":
		return TimezonePacificTahiti, nil
	case "Pacific/Tarawa":
		return TimezonePacificTarawa, nil
	case "Pacific/Tongatapu":
		return TimezonePacificTongatapu, nil
	case "Pacific/Truk":
		return TimezonePacificTruk, nil
	case "Pacific/Wake":
		return TimezonePacificWake, nil
	case "Pacific/Wallis":
		return TimezonePacificWallis, nil
	case "Pacific/Yap":
		return TimezonePacificYap, nil
	case "Poland":
		return TimezonePoland, nil
	case "Portugal":
		return TimezonePortugal, nil
	case "PRC":
		return TimezonePRC, nil
	case "ROC":
		return TimezoneROC, nil
	case "Singapore":
		return TimezoneSingapore, nil
	case "US/Alaska":
		return TimezoneUSAlaska, nil
	case "US/Aleutian":
		return TimezoneUSAleutian, nil
	case "US/Arizona":
		return TimezoneUSArizona, nil
	case "US/Central":
		return TimezoneUSCentral, nil
	case "US/East-Indiana":
		return TimezoneUSEastIndiana, nil
	case "US/Eastern":
		return TimezoneUSEastern, nil
	case "US/Hawaii":
		return TimezoneUSHawaii, nil
	case "US/Indiana-Starke":
		return TimezoneUSIndianaStarke, nil
	case "US/Michigan":
		return TimezoneUSMichigan, nil
	case "US/Mountain":
		return TimezoneUSMountain, nil
	case "US/Pacific":
		return TimezoneUSPacific, nil
	case "US/Samoa":
		return TimezoneUSSamoa, nil
	default:
		return "", errors.New("unknown timezone")
	}
}

func (t Timezone) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Timezone) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*t = Timezone(s)
	return nil
}

func (t Timezone) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Timezone) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	*t = Timezone(s)
	return nil
}

func (t Timezone) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *Timezone) UnmarshalText(b []byte) error {
	*t = Timezone(b)
	return nil
}

var TimeZones = []string{
	"Africa/Abidjan",
	"Africa/Accra",
	"Africa/Addis_Ababa",
	"Africa/Algiers",
	"Africa/Asmara",
	"Africa/Asmera",
	"Africa/Bamako",
	"Africa/Bangui",
	"Africa/Banjul",
	"Africa/Bissau",
	"Africa/Blantyre",
	"Africa/Brazzaville",
	"Africa/Bujumbura",
	"Africa/Cairo",
	"Africa/Casablanca",
	"Africa/Ceuta",
	"Africa/Conakry",
	"Africa/Dakar",
	"Africa/Dar_es_Salaam",
	"Africa/Djibouti",
	"Africa/Douala",
	"Africa/El_Aaiun",
	"Africa/Freetown",
	"Africa/Gaborone",
	"Africa/Harare",
	"Africa/Johannesburg",
	"Africa/Juba",
	"Africa/Kampala",
	"Africa/Khartoum",
	"Africa/Kigali",
	"Africa/Kinshasa",
	"Africa/Lagos",
	"Africa/Libreville",
	"Africa/Lome",
	"Africa/Luanda",
	"Africa/Lubumbashi",
	"Africa/Lusaka",
	"Africa/Malabo",
	"Africa/Maputo",
	"Africa/Maseru",
	"Africa/Mbabane",
	"Africa/Mogadishu",
	"Africa/Monrovia",
	"Africa/Nairobi",
	"Africa/Ndjamena",
	"Africa/Niamey",
	"Africa/Nouakchott",
	"Africa/Ouagadougou",
	"Africa/Porto-Novo",
	"Africa/Sao_Tome",
	"Africa/Timbuktu",
	"Africa/Tripoli",
	"Africa/Tunis",
	"Africa/Windhoek",
	"America/Adak",
	"America/Anchorage",
	"America/Anguilla",
	"America/Antigua",
	"America/Araguaina",
	"America/Argentina/Buenos_Aires",
	"America/Argentina/Catamarca",
	"America/Argentina/ComodRivadavia",
	"America/Argentina/Cordoba",
	"America/Argentina/Jujuy",
	"America/Argentina/La_Rioja",
	"America/Argentina/Mendoza",
	"America/Argentina/Rio_Gallegos",
	"America/Argentina/Salta",
	"America/Argentina/San_Juan",
	"America/Argentina/San_Luis",
	"America/Argentina/Tucuman",
	"America/Argentina/Ushuaia",
	"America/Aruba",
	"America/Asuncion",
	"America/Atikokan",
	"America/Atka",
	"America/Bahia",
	"America/Bahia_Banderas",
	"America/Barbados",
	"America/Belem",
	"America/Belize",
	"America/Blanc-Sablon",
	"America/Boa_Vista",
	"America/Bogota",
	"America/Boise",
	"America/Buenos_Aires",
	"America/Cambridge_Bay",
	"America/Campo_Grande",
	"America/Cancun",
	"America/Caracas",
	"America/Catamarca",
	"America/Cayenne",
	"America/Cayman",
	"America/Chicago",
	"America/Chihuahua",
	"America/Coral_Harbour",
	"America/Cordoba",
	"America/Costa_Rica",
	"America/Creston",
	"America/Cuiaba",
	"America/Curacao",
	"America/Danmarkshavn",
	"America/Dawson",
	"America/Dawson_Creek",
	"America/Denver",
	"America/Detroit",
	"America/Dominica",
	"America/Edmonton",
	"America/Eirunepe",
	"America/El_Salvador",
	"America/Ensenada",
	"America/Fort_Nelson",
	"America/Fort_Wayne",
	"America/Fortaleza",
	"America/Glace_Bay",
	"America/Godthab",
	"America/Goose_Bay",
	"America/Grand_Turk",
	"America/Grenada",
	"America/Guadeloupe",
	"America/Guatemala",
	"America/Guayaquil",
	"America/Guyana",
	"America/Halifax",
	"America/Havana",
	"America/Hermosillo",
	"America/Indiana/Indianapolis",
	"America/Indiana/Knox",
	"America/Indiana/Marengo",
	"America/Indiana/Petersburg",
	"America/Indiana/Tell_City",
	"America/Indiana/Vevay",
	"America/Indiana/Vincennes",
	"America/Indiana/Winamac",
	"America/Indianapolis",
	"America/Inuvik",
	"America/Iqaluit",
	"America/Jamaica",
	"America/Jujuy",
	"America/Juneau",
	"America/Kentucky/Louisville",
	"America/Kentucky/Monticello",
	"America/Knox_IN",
	"America/Kralendijk",
	"America/La_Paz",
	"America/Lima",
	"America/Los_Angeles",
	"America/Louisville",
	"America/Lower_Princes",
	"America/Maceio",
	"America/Managua",
	"America/Manaus",
	"America/Marigot",
	"America/Martinique",
	"America/Matamoros",
	"America/Mazatlan",
	"America/Mendoza",
	"America/Menominee",
	"America/Merida",
	"America/Metlakatla",
	"America/Mexico_City",
	"America/Miquelon",
	"America/Moncton",
	"America/Monterrey",
	"America/Montevideo",
	"America/Montreal",
	"America/Montserrat",
	"America/Nassau",
	"America/New_York",
	"America/Nipigon",
	"America/Nome",
	"America/Noronha",
	"America/North_Dakota/Beulah",
	"America/North_Dakota/Center",
	"America/North_Dakota/New_Salem",
	"America/Nuuk",
	"America/Ojinaga",
	"America/Panama",
	"America/Pangnirtung",
	"America/Paramaribo",
	"America/Phoenix",
	"America/Port-au-Prince",
	"America/Port_of_Spain",
	"America/Porto_Acre",
	"America/Porto_Velho",
	"America/Puerto_Rico",
	"America/Punta_Arenas",
	"America/Rainy_River",
	"America/Rankin_Inlet",
	"America/Recife",
	"America/Regina",
	"America/Resolute",
	"America/Rio_Branco",
	"America/Rosario",
	"America/Santa_Isabel",
	"America/Santarem",
	"America/Santiago",
	"America/Santo_Domingo",
	"America/Sao_Paulo",
	"America/Scoresbysund",
	"America/Shiprock",
	"America/Sitka",
	"America/St_Barthelemy",
	"America/St_Johns",
	"America/St_Kitts",
	"America/St_Lucia",
	"America/St_Thomas",
	"America/St_Vincent",
	"America/Swift_Current",
	"America/Tegucigalpa",
	"America/Thule",
	"America/Thunder_Bay",
	"America/Tijuana",
	"America/Toronto",
	"America/Tortola",
	"America/Vancouver",
	"America/Virgin",
	"America/Whitehorse",
	"America/Winnipeg",
	"America/Yakutat",
	"America/Yellowknife",
	"Antarctica/Casey",
	"Antarctica/Davis",
	"Antarctica/DumontDUrville",
	"Antarctica/Macquarie",
	"Antarctica/Mawson",
	"Antarctica/McMurdo",
	"Antarctica/Palmer",
	"Antarctica/Rothera",
	"Antarctica/South_Pole",
	"Antarctica/Syowa",
	"Antarctica/Troll",
	"Antarctica/Vostok",
	"Arctic/Longyearbyen",
	"Asia/Aden",
	"Asia/Almaty",
	"Asia/Amman",
	"Asia/Anadyr",
	"Asia/Aqtau",
	"Asia/Aqtobe",
	"Asia/Ashgabat",
	"Asia/Ashkhabad",
	"Asia/Atyrau",
	"Asia/Baghdad",
	"Asia/Bahrain",
	"Asia/Baku",
	"Asia/Bangkok",
	"Asia/Barnaul",
	"Asia/Beirut",
	"Asia/Bishkek",
	"Asia/Brunei",
	"Asia/Calcutta",
	"Asia/Chita",
	"Asia/Choibalsan",
	"Asia/Chongqing",
	"Asia/Chungking",
	"Asia/Colombo",
	"Asia/Dacca",
	"Asia/Damascus",
	"Asia/Dhaka",
	"Asia/Dili",
	"Asia/Dubai",
	"Asia/Dushanbe",
	"Asia/Famagusta",
	"Asia/Gaza",
	"Asia/Harbin",
	"Asia/Hebron",
	"Asia/Ho_Chi_Minh",
	"Asia/Hong_Kong",
	"Asia/Hovd",
	"Asia/Irkutsk",
	"Asia/Istanbul",
	"Asia/Jakarta",
	"Asia/Jayapura",
	"Asia/Jerusalem",
	"Asia/Kabul",
	"Asia/Kamchatka",
	"Asia/Karachi",
	"Asia/Kashgar",
	"Asia/Kathmandu",
	"Asia/Katmandu",
	"Asia/Khandyga",
	"Asia/Kolkata",
	"Asia/Krasnoyarsk",
	"Asia/Kuala_Lumpur",
	"Asia/Kuching",
	"Asia/Kuwait",
	"Asia/Macao",
	"Asia/Macau",
	"Asia/Magadan",
	"Asia/Makassar",
	"Asia/Manila",
	"Asia/Muscat",
	"Asia/Nicosia",
	"Asia/Novokuznetsk",
	"Asia/Novosibirsk",
	"Asia/Omsk",
	"Asia/Oral",
	"Asia/Phnom_Penh",
	"Asia/Pontianak",
	"Asia/Pyongyang",
	"Asia/Qatar",
	"Asia/Qostanay",
	"Asia/Qyzylorda",
	"Asia/Rangoon",
	"Asia/Riyadh",
	"Asia/Saigon",
	"Asia/Sakhalin",
	"Asia/Samarkand",
	"Asia/Seoul",
	"Asia/Shanghai",
	"Asia/Singapore",
	"Asia/Srednekolymsk",
	"Asia/Taipei",
	"Asia/Tashkent",
	"Asia/Tbilisi",
	"Asia/Tehran",
	"Asia/Tel_Aviv",
	"Asia/Thimbu",
	"Asia/Thimphu",
	"Asia/Tokyo",
	"Asia/Tomsk",
	"Asia/Ujung_Pandang",
	"Asia/Ulaanbaatar",
	"Asia/Ulan_Bator",
	"Asia/Urumqi",
	"Asia/Ust-Nera",
	"Asia/Vientiane",
	"Asia/Vladivostok",
	"Asia/Yakutsk",
	"Asia/Yangon",
	"Asia/Yekaterinburg",
	"Asia/Yerevan",
	"Atlantic/Azores",
	"Atlantic/Bermuda",
	"Atlantic/Canary",
	"Atlantic/Cape_Verde",
	"Atlantic/Faeroe",
	"Atlantic/Faroe",
	"Atlantic/Jan_Mayen",
	"Atlantic/Madeira",
	"Atlantic/Reykjavik",
	"Atlantic/South_Georgia",
	"Atlantic/St_Helena",
	"Atlantic/Stanley",
	"Australia/ACT",
	"Australia/Adelaide",
	"Australia/Brisbane",
	"Australia/Broken_Hill",
	"Australia/Canberra",
	"Australia/Currie",
	"Australia/Darwin",
	"Australia/Eucla",
	"Australia/Hobart",
	"Australia/LHI",
	"Australia/Lindeman",
	"Australia/Lord_Howe",
	"Australia/Melbourne",
	"Australia/North",
	"Australia/NSW",
	"Australia/Perth",
	"Australia/Queensland",
	"Australia/South",
	"Australia/Sydney",
	"Australia/Tasmania",
	"Australia/Victoria",
	"Australia/West",
	"Australia/Yancowinna",
	"Brazil/Acre",
	"Brazil/DeNoronha",
	"Brazil/East",
	"Brazil/West",
	"Canada/Atlantic",
	"Canada/Central",
	"Canada/Eastern",
	"Canada/Mountain",
	"Canada/Newfoundland",
	"Canada/Pacific",
	"Canada/Saskatchewan",
	"Canada/Yukon",
	"Chile/Continental",
	"Chile/EasterIsland",
	"Cuba",
	"Egypt",
	"Eire",
	"Etc/GMT",
	"Etc/GMT+0",
	"Etc/GMT+1",
	"Etc/GMT+10",
	"Etc/GMT+11",
	"Etc/GMT+12",
	"Etc/GMT+2",
	"Etc/GMT+3",
	"Etc/GMT+4",
	"Etc/GMT+5",
	"Etc/GMT+6",
	"Etc/GMT+7",
	"Etc/GMT+8",
	"Etc/GMT+9",
	"Etc/GMT-0",
	"Etc/GMT-1",
	"Etc/GMT-10",
	"Etc/GMT-11",
	"Etc/GMT-12",
	"Etc/GMT-13",
	"Etc/GMT-14",
	"Etc/GMT-2",
	"Etc/GMT-3",
	"Etc/GMT-4",
	"Etc/GMT-5",
	"Etc/GMT-6",
	"Etc/GMT-7",
	"Etc/GMT-8",
	"Etc/GMT-9",
	"Etc/GMT0",
	"Etc/Greenwich",
	"Etc/UCT",
	"Etc/Universal",
	"Etc/UTC",
	"Etc/Zulu",
	"Europe/Amsterdam",
	"Europe/Andorra",
	"Europe/Astrakhan",
	"Europe/Athens",
	"Europe/Belfast",
	"Europe/Belgrade",
	"Europe/Berlin",
	"Europe/Bratislava",
	"Europe/Brussels",
	"Europe/Bucharest",
	"Europe/Budapest",
	"Europe/Busingen",
	"Europe/Chisinau",
	"Europe/Copenhagen",
	"Europe/Dublin",
	"Europe/Gibraltar",
	"Europe/Guernsey",
	"Europe/Helsinki",
	"Europe/Isle_of_Man",
	"Europe/Istanbul",
	"Europe/Jersey",
	"Europe/Kaliningrad",
	"Europe/Kiev",
	"Europe/Kirov",
	"Europe/Kyiv",
	"Europe/Lisbon",
	"Europe/Ljubljana",
	"Europe/London",
	"Europe/Luxembourg",
	"Europe/Madrid",
	"Europe/Malta",
	"Europe/Mariehamn",
	"Europe/Minsk",
	"Europe/Monaco",
	"Europe/Moscow",
	"Europe/Nicosia",
	"Europe/Oslo",
	"Europe/Paris",
	"Europe/Podgorica",
	"Europe/Prague",
	"Europe/Riga",
	"Europe/Rome",
	"Europe/Samara",
	"Europe/San_Marino",
	"Europe/Sarajevo",
	"Europe/Saratov",
	"Europe/Simferopol",
	"Europe/Skopje",
	"Europe/Sofia",
	"Europe/Stockholm",
	"Europe/Tallinn",
	"Europe/Tirane",
	"Europe/Tiraspol",
	"Europe/Ulyanovsk",
	"Europe/Uzhgorod",
	"Europe/Vaduz",
	"Europe/Vatican",
	"Europe/Vienna",
	"Europe/Vilnius",
	"Europe/Volgograd",
	"Europe/Warsaw",
	"Europe/Zagreb",
	"Europe/Zaporozhye",
	"Europe/Zurich",
	"GB",
	"GB-Eire",
	"Hongkong",
	"Iceland",
	"Indian/Antananarivo",
	"Indian/Chagos",
	"Indian/Christmas",
	"Indian/Cocos",
	"Indian/Comoro",
	"Indian/Kerguelen",
	"Indian/Mahe",
	"Indian/Maldives",
	"Indian/Mauritius",
	"Indian/Mayotte",
	"Indian/Reunion",
	"Iran",
	"Israel",
	"Jamaica",
	"Japan",
	"Kwajalein",
	"Libya",
	"Mexico/BajaNorte",
	"Mexico/BajaSur",
	"Mexico/General",
	"Navajo",
	"NZ",
	"NZ-CHAT",
	"Pacific/Apia",
	"Pacific/Auckland",
	"Pacific/Bougainville",
	"Pacific/Chatham",
	"Pacific/Chuuk",
	"Pacific/Easter",
	"Pacific/Efate",
	"Pacific/Enderbury",
	"Pacific/Fakaofo",
	"Pacific/Fiji",
	"Pacific/Funafuti",
	"Pacific/Galapagos",
	"Pacific/Gambier",
	"Pacific/Guadalcanal",
	"Pacific/Guam",
	"Pacific/Honolulu",
	"Pacific/Johnston",
	"Pacific/Kanton",
	"Pacific/Kiritimati",
	"Pacific/Kosrae",
	"Pacific/Kwajalein",
	"Pacific/Majuro",
	"Pacific/Marquesas",
	"Pacific/Midway",
	"Pacific/Nauru",
	"Pacific/Niue",
	"Pacific/Norfolk",
	"Pacific/Noumea",
	"Pacific/Pago_Pago",
	"Pacific/Palau",
	"Pacific/Pitcairn",
	"Pacific/Pohnpei",
	"Pacific/Ponape",
	"Pacific/Port_Moresby",
	"Pacific/Rarotonga",
	"Pacific/Saipan",
	"Pacific/Samoa",
	"Pacific/Tahiti",
	"Pacific/Tarawa",
	"Pacific/Tongatapu",
	"Pacific/Truk",
	"Pacific/Wake",
	"Pacific/Wallis",
	"Pacific/Yap",
	"Poland",
	"Portugal",
	"PRC",
	"ROC",
	"Singapore",
	"US/Alaska",
	"US/Aleutian",
	"US/Arizona",
	"US/Central",
	"US/East-Indiana",
	"US/Eastern",
	"US/Hawaii",
	"US/Indiana-Starke",
	"US/Michigan",
	"US/Mountain",
	"US/Pacific",
	"US/Samoa",
}
