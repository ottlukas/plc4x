//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type KnxManufacturer uint16

type IKnxManufacturer interface {
    Number() uint16
    Name() string
    Serialize(io utils.WriteBuffer) error
}

const(
    KnxManufacturer_M_UNKNOWN KnxManufacturer = 0
    KnxManufacturer_M_SIEMENS KnxManufacturer = 1
    KnxManufacturer_M_ABB KnxManufacturer = 2
    KnxManufacturer_M_ALBRECHT_JUNG KnxManufacturer = 3
    KnxManufacturer_M_BTICINO KnxManufacturer = 4
    KnxManufacturer_M_BERKER KnxManufacturer = 5
    KnxManufacturer_M_BUSCH_JAEGER_ELEKTRO KnxManufacturer = 6
    KnxManufacturer_M_GIRA_GIERSIEPEN KnxManufacturer = 7
    KnxManufacturer_M_HAGER_ELECTRO KnxManufacturer = 8
    KnxManufacturer_M_INSTA_GMBH KnxManufacturer = 9
    KnxManufacturer_M_LEGRAND_APPAREILLAGE_ELECTRIQUE KnxManufacturer = 10
    KnxManufacturer_M_MERTEN KnxManufacturer = 11
    KnxManufacturer_M_ABB_SPA_SACE_DIVISION KnxManufacturer = 12
    KnxManufacturer_M_SIEDLE_AND_SOEHNE KnxManufacturer = 13
    KnxManufacturer_M_EBERLE KnxManufacturer = 14
    KnxManufacturer_M_GEWISS KnxManufacturer = 15
    KnxManufacturer_M_ALBERT_ACKERMANN KnxManufacturer = 16
    KnxManufacturer_M_SCHUPA_GMBH KnxManufacturer = 17
    KnxManufacturer_M_ABB_SCHWEIZ KnxManufacturer = 18
    KnxManufacturer_M_FELLER KnxManufacturer = 19
    KnxManufacturer_M_GLAMOX_AS KnxManufacturer = 20
    KnxManufacturer_M_DEHN_AND_SOEHNE KnxManufacturer = 21
    KnxManufacturer_M_CRABTREE KnxManufacturer = 22
    KnxManufacturer_M_EVOKNX KnxManufacturer = 23
    KnxManufacturer_M_PAUL_HOCHKOEPPER KnxManufacturer = 24
    KnxManufacturer_M_ALTENBURGER_ELECTRONIC KnxManufacturer = 25
    KnxManufacturer_M_GRAESSLIN KnxManufacturer = 26
    KnxManufacturer_M_SIMON_42 KnxManufacturer = 27
    KnxManufacturer_M_VIMAR KnxManufacturer = 28
    KnxManufacturer_M_MOELLER_GEBAEUDEAUTOMATION_KG KnxManufacturer = 29
    KnxManufacturer_M_ELTAKO KnxManufacturer = 30
    KnxManufacturer_M_BOSCH_SIEMENS_HAUSHALTSGERAETE KnxManufacturer = 31
    KnxManufacturer_M_RITTO_GMBHANDCO_KG KnxManufacturer = 32
    KnxManufacturer_M_POWER_CONTROLS KnxManufacturer = 33
    KnxManufacturer_M_ZUMTOBEL KnxManufacturer = 34
    KnxManufacturer_M_PHOENIX_CONTACT KnxManufacturer = 35
    KnxManufacturer_M_WAGO_KONTAKTTECHNIK KnxManufacturer = 36
    KnxManufacturer_M_KNXPRESSO KnxManufacturer = 37
    KnxManufacturer_M_WIELAND_ELECTRIC KnxManufacturer = 38
    KnxManufacturer_M_HERMANN_KLEINHUIS KnxManufacturer = 39
    KnxManufacturer_M_STIEBEL_ELTRON KnxManufacturer = 40
    KnxManufacturer_M_TEHALIT KnxManufacturer = 41
    KnxManufacturer_M_THEBEN_AG KnxManufacturer = 42
    KnxManufacturer_M_WILHELM_RUTENBECK KnxManufacturer = 43
    KnxManufacturer_M_WINKHAUS KnxManufacturer = 44
    KnxManufacturer_M_ROBERT_BOSCH KnxManufacturer = 45
    KnxManufacturer_M_SOMFY KnxManufacturer = 46
    KnxManufacturer_M_WOERTZ KnxManufacturer = 47
    KnxManufacturer_M_VIESSMANN_WERKE KnxManufacturer = 48
    KnxManufacturer_M_IMI_HYDRONIC_ENGINEERING KnxManufacturer = 49
    KnxManufacturer_M_JOH__VAILLANT KnxManufacturer = 50
    KnxManufacturer_M_AMP_DEUTSCHLAND KnxManufacturer = 51
    KnxManufacturer_M_BOSCH_THERMOTECHNIK_GMBH KnxManufacturer = 52
    KnxManufacturer_M_SEF___ECOTEC KnxManufacturer = 53
    KnxManufacturer_M_DORMA_GMBH_Plus_CO__KG KnxManufacturer = 54
    KnxManufacturer_M_WINDOWMASTER_AS KnxManufacturer = 55
    KnxManufacturer_M_WALTHER_WERKE KnxManufacturer = 56
    KnxManufacturer_M_ORAS KnxManufacturer = 57
    KnxManufacturer_M_DAETWYLER KnxManufacturer = 58
    KnxManufacturer_M_ELECTRAK KnxManufacturer = 59
    KnxManufacturer_M_TECHEM KnxManufacturer = 60
    KnxManufacturer_M_SCHNEIDER_ELECTRIC_INDUSTRIES_SAS KnxManufacturer = 61
    KnxManufacturer_M_WHD_WILHELM_HUBER_Plus_SOEHNE KnxManufacturer = 62
    KnxManufacturer_M_BISCHOFF_ELEKTRONIK KnxManufacturer = 63
    KnxManufacturer_M_JEPAZ KnxManufacturer = 64
    KnxManufacturer_M_RTS_AUTOMATION KnxManufacturer = 65
    KnxManufacturer_M_EIBMARKT_GMBH KnxManufacturer = 66
    KnxManufacturer_M_WAREMA_RENKHOFF_SE KnxManufacturer = 67
    KnxManufacturer_M_EELECTRON KnxManufacturer = 68
    KnxManufacturer_M_BELDEN_WIRE_AND_CABLE_B_V_ KnxManufacturer = 69
    KnxManufacturer_M_BECKER_ANTRIEBE_GMBH KnxManufacturer = 70
    KnxManufacturer_M_J_STEHLEPlusSOEHNE_GMBH KnxManufacturer = 71
    KnxManufacturer_M_AGFEO KnxManufacturer = 72
    KnxManufacturer_M_ZENNIO KnxManufacturer = 73
    KnxManufacturer_M_TAPKO_TECHNOLOGIES KnxManufacturer = 74
    KnxManufacturer_M_HDL KnxManufacturer = 75
    KnxManufacturer_M_UPONOR KnxManufacturer = 76
    KnxManufacturer_M_SE_LIGHTMANAGEMENT_AG KnxManufacturer = 77
    KnxManufacturer_M_ARCUS_EDS KnxManufacturer = 78
    KnxManufacturer_M_INTESIS KnxManufacturer = 79
    KnxManufacturer_M_HERHOLDT_CONTROLS_SRL KnxManufacturer = 80
    KnxManufacturer_M_NIKO_ZUBLIN KnxManufacturer = 81
    KnxManufacturer_M_DURABLE_TECHNOLOGIES KnxManufacturer = 82
    KnxManufacturer_M_INNOTEAM KnxManufacturer = 83
    KnxManufacturer_M_ISE_GMBH KnxManufacturer = 84
    KnxManufacturer_M_TEAM_FOR_TRONICS KnxManufacturer = 85
    KnxManufacturer_M_CIAT KnxManufacturer = 86
    KnxManufacturer_M_REMEHA_BV KnxManufacturer = 87
    KnxManufacturer_M_ESYLUX KnxManufacturer = 88
    KnxManufacturer_M_BASALTE KnxManufacturer = 89
    KnxManufacturer_M_VESTAMATIC KnxManufacturer = 90
    KnxManufacturer_M_MDT_TECHNOLOGIES KnxManufacturer = 91
    KnxManufacturer_M_WARENDORFER_KUECHEN_GMBH KnxManufacturer = 92
    KnxManufacturer_M_VIDEO_STAR KnxManufacturer = 93
    KnxManufacturer_M_SITEK KnxManufacturer = 94
    KnxManufacturer_M_CONTROLTRONIC KnxManufacturer = 95
    KnxManufacturer_M_FUNCTION_TECHNOLOGY KnxManufacturer = 96
    KnxManufacturer_M_AMX KnxManufacturer = 97
    KnxManufacturer_M_ELDAT KnxManufacturer = 98
    KnxManufacturer_M_PANASONIC KnxManufacturer = 99
    KnxManufacturer_M_PULSE_TECHNOLOGIES KnxManufacturer = 100
    KnxManufacturer_M_CRESTRON KnxManufacturer = 101
    KnxManufacturer_M_STEINEL_PROFESSIONAL KnxManufacturer = 102
    KnxManufacturer_M_BILTON_LED_LIGHTING KnxManufacturer = 103
    KnxManufacturer_M_DENRO_AG KnxManufacturer = 104
    KnxManufacturer_M_GEPRO KnxManufacturer = 105
    KnxManufacturer_M_PREUSSEN_AUTOMATION KnxManufacturer = 106
    KnxManufacturer_M_ZOPPAS_INDUSTRIES KnxManufacturer = 107
    KnxManufacturer_M_MACTECH KnxManufacturer = 108
    KnxManufacturer_M_TECHNO_TREND KnxManufacturer = 109
    KnxManufacturer_M_FS_CABLES KnxManufacturer = 110
    KnxManufacturer_M_DELTA_DORE KnxManufacturer = 111
    KnxManufacturer_M_EISSOUND KnxManufacturer = 112
    KnxManufacturer_M_CISCO KnxManufacturer = 113
    KnxManufacturer_M_DINUY KnxManufacturer = 114
    KnxManufacturer_M_IKNIX KnxManufacturer = 115
    KnxManufacturer_M_RADEMACHER_GERAETE_ELEKTRONIK_GMBH KnxManufacturer = 116
    KnxManufacturer_M_EGI_ELECTROACUSTICA_GENERAL_IBERICA KnxManufacturer = 117
    KnxManufacturer_M_BES___INGENIUM KnxManufacturer = 118
    KnxManufacturer_M_ELABNET KnxManufacturer = 119
    KnxManufacturer_M_BLUMOTIX KnxManufacturer = 120
    KnxManufacturer_M_HUNTER_DOUGLAS KnxManufacturer = 121
    KnxManufacturer_M_APRICUM KnxManufacturer = 122
    KnxManufacturer_M_TIANSU_AUTOMATION KnxManufacturer = 123
    KnxManufacturer_M_BUBENDORFF KnxManufacturer = 124
    KnxManufacturer_M_MBS_GMBH KnxManufacturer = 125
    KnxManufacturer_M_ENERTEX_BAYERN_GMBH KnxManufacturer = 126
    KnxManufacturer_M_BMS KnxManufacturer = 127
    KnxManufacturer_M_SINAPSI KnxManufacturer = 128
    KnxManufacturer_M_EMBEDDED_SYSTEMS_SIA KnxManufacturer = 129
    KnxManufacturer_M_KNX1 KnxManufacturer = 130
    KnxManufacturer_M_TOKKA KnxManufacturer = 131
    KnxManufacturer_M_NANOSENSE KnxManufacturer = 132
    KnxManufacturer_M_PEAR_AUTOMATION_GMBH KnxManufacturer = 133
    KnxManufacturer_M_DGA KnxManufacturer = 134
    KnxManufacturer_M_LUTRON KnxManufacturer = 135
    KnxManufacturer_M_AIRZONE___ALTRA KnxManufacturer = 136
    KnxManufacturer_M_LITHOSS_DESIGN_SWITCHES KnxManufacturer = 137
    KnxManufacturer_M_THREEATEL KnxManufacturer = 138
    KnxManufacturer_M_PHILIPS_CONTROLS KnxManufacturer = 139
    KnxManufacturer_M_VELUX_AS KnxManufacturer = 140
    KnxManufacturer_M_LOYTEC KnxManufacturer = 141
    KnxManufacturer_M_EKINEX_S_P_A_ KnxManufacturer = 142
    KnxManufacturer_M_SIRLAN_TECHNOLOGIES KnxManufacturer = 143
    KnxManufacturer_M_PROKNX_SAS KnxManufacturer = 144
    KnxManufacturer_M_IT_GMBH KnxManufacturer = 145
    KnxManufacturer_M_RENSON KnxManufacturer = 146
    KnxManufacturer_M_HEP_GROUP KnxManufacturer = 147
    KnxManufacturer_M_BALMART KnxManufacturer = 148
    KnxManufacturer_M_GFS_GMBH KnxManufacturer = 149
    KnxManufacturer_M_SCHENKER_STOREN_AG KnxManufacturer = 150
    KnxManufacturer_M_ALGODUE_ELETTRONICA_S_R_L_ KnxManufacturer = 151
    KnxManufacturer_M_ABB_FRANCE KnxManufacturer = 152
    KnxManufacturer_M_MAINTRONIC KnxManufacturer = 153
    KnxManufacturer_M_VANTAGE KnxManufacturer = 154
    KnxManufacturer_M_FORESIS KnxManufacturer = 155
    KnxManufacturer_M_RESEARCH_AND_PRODUCTION_ASSOCIATION_SEM KnxManufacturer = 156
    KnxManufacturer_M_WEINZIERL_ENGINEERING_GMBH KnxManufacturer = 157
    KnxManufacturer_M_MOEHLENHOFF_WAERMETECHNIK_GMBH KnxManufacturer = 158
    KnxManufacturer_M_PKC_GROUP_OYJ KnxManufacturer = 159
    KnxManufacturer_M_B_E_G_ KnxManufacturer = 160
    KnxManufacturer_M_ELSNER_ELEKTRONIK_GMBH KnxManufacturer = 161
    KnxManufacturer_M_SIEMENS_BUILDING_TECHNOLOGIES_HKCHINA_LTD_ KnxManufacturer = 162
    KnxManufacturer_M_EUTRAC KnxManufacturer = 163
    KnxManufacturer_M_GUSTAV_HENSEL_GMBH_AND_CO__KG KnxManufacturer = 164
    KnxManufacturer_M_GARO_AB KnxManufacturer = 165
    KnxManufacturer_M_WALDMANN_LICHTTECHNIK KnxManufacturer = 166
    KnxManufacturer_M_SCHUECO KnxManufacturer = 167
    KnxManufacturer_M_EMU KnxManufacturer = 168
    KnxManufacturer_M_JNET_SYSTEMS_AG KnxManufacturer = 169
    KnxManufacturer_M_TOTAL_SOLUTION_GMBH KnxManufacturer = 170
    KnxManufacturer_M_O_Y_L__ELECTRONICS KnxManufacturer = 171
    KnxManufacturer_M_GALAX_SYSTEM KnxManufacturer = 172
    KnxManufacturer_M_DISCH KnxManufacturer = 173
    KnxManufacturer_M_AUCOTEAM KnxManufacturer = 174
    KnxManufacturer_M_LUXMATE_CONTROLS KnxManufacturer = 175
    KnxManufacturer_M_DANFOSS KnxManufacturer = 176
    KnxManufacturer_M_AST_GMBH KnxManufacturer = 177
    KnxManufacturer_M_WILA_LEUCHTEN KnxManufacturer = 178
    KnxManufacturer_M_BPlusB_AUTOMATIONS__UND_STEUERUNGSTECHNIK KnxManufacturer = 179
    KnxManufacturer_M_LINGG_AND_JANKE KnxManufacturer = 180
    KnxManufacturer_M_SAUTER KnxManufacturer = 181
    KnxManufacturer_M_SIMU KnxManufacturer = 182
    KnxManufacturer_M_THEBEN_HTS_AG KnxManufacturer = 183
    KnxManufacturer_M_AMANN_GMBH KnxManufacturer = 184
    KnxManufacturer_M_BERG_ENERGIEKONTROLLSYSTEME_GMBH KnxManufacturer = 185
    KnxManufacturer_M_HUEPPE_FORM_SONNENSCHUTZSYSTEME_GMBH KnxManufacturer = 186
    KnxManufacturer_M_OVENTROP_KG KnxManufacturer = 187
    KnxManufacturer_M_GRIESSER_AG KnxManufacturer = 188
    KnxManufacturer_M_IPAS_GMBH KnxManufacturer = 189
    KnxManufacturer_M_ELERO_GMBH KnxManufacturer = 190
    KnxManufacturer_M_ARDAN_PRODUCTION_AND_INDUSTRIAL_CONTROLS_LTD_ KnxManufacturer = 191
    KnxManufacturer_M_METEC_MESSTECHNIK_GMBH KnxManufacturer = 192
    KnxManufacturer_M_ELKA_ELEKTRONIK_GMBH KnxManufacturer = 193
    KnxManufacturer_M_ELEKTROANLAGEN_D__NAGEL KnxManufacturer = 194
    KnxManufacturer_M_TRIDONIC_BAUELEMENTE_GMBH KnxManufacturer = 195
    KnxManufacturer_M_STENGLER_GESELLSCHAFT KnxManufacturer = 196
    KnxManufacturer_M_SCHNEIDER_ELECTRIC_MG KnxManufacturer = 197
    KnxManufacturer_M_KNX_ASSOCIATION KnxManufacturer = 198
    KnxManufacturer_M_VIVO KnxManufacturer = 199
    KnxManufacturer_M_HUGO_MUELLER_GMBH_AND_CO_KG KnxManufacturer = 200
    KnxManufacturer_M_SIEMENS_HVAC KnxManufacturer = 201
    KnxManufacturer_M_APT KnxManufacturer = 202
    KnxManufacturer_M_HIGHDOM KnxManufacturer = 203
    KnxManufacturer_M_TOP_SERVICES KnxManufacturer = 204
    KnxManufacturer_M_AMBIHOME KnxManufacturer = 205
    KnxManufacturer_M_DATEC_ELECTRONIC_AG KnxManufacturer = 206
    KnxManufacturer_M_ABUS_SECURITY_CENTER KnxManufacturer = 207
    KnxManufacturer_M_LITE_PUTER KnxManufacturer = 208
    KnxManufacturer_M_TANTRON_ELECTRONIC KnxManufacturer = 209
    KnxManufacturer_M_INTERRA KnxManufacturer = 210
    KnxManufacturer_M_DKX_TECH KnxManufacturer = 211
    KnxManufacturer_M_VIATRON KnxManufacturer = 212
    KnxManufacturer_M_NAUTIBUS KnxManufacturer = 213
    KnxManufacturer_M_ON_SEMICONDUCTOR KnxManufacturer = 214
    KnxManufacturer_M_LONGCHUANG KnxManufacturer = 215
    KnxManufacturer_M_AIR_ON_AG KnxManufacturer = 216
    KnxManufacturer_M_IB_COMPANY_GMBH KnxManufacturer = 217
    KnxManufacturer_M_SATION_FACTORY KnxManufacturer = 218
    KnxManufacturer_M_AGENTILO_GMBH KnxManufacturer = 219
    KnxManufacturer_M_MAKEL_ELEKTRIK KnxManufacturer = 220
    KnxManufacturer_M_HELIOS_VENTILATOREN KnxManufacturer = 221
    KnxManufacturer_M_OTTO_SOLUTIONS_PTE_LTD KnxManufacturer = 222
    KnxManufacturer_M_AIRMASTER KnxManufacturer = 223
    KnxManufacturer_M_VALLOX_GMBH KnxManufacturer = 224
    KnxManufacturer_M_DALITEK KnxManufacturer = 225
    KnxManufacturer_M_ASIN KnxManufacturer = 226
    KnxManufacturer_M_BRIDGES_INTELLIGENCE_TECHNOLOGY_INC_ KnxManufacturer = 227
    KnxManufacturer_M_ARBONIA KnxManufacturer = 228
    KnxManufacturer_M_KERMI KnxManufacturer = 229
    KnxManufacturer_M_PROLUX KnxManufacturer = 230
    KnxManufacturer_M_CLICHOME KnxManufacturer = 231
    KnxManufacturer_M_COMMAX KnxManufacturer = 232
    KnxManufacturer_M_EAE KnxManufacturer = 233
    KnxManufacturer_M_TENSE KnxManufacturer = 234
    KnxManufacturer_M_SEYOUNG_ELECTRONICS KnxManufacturer = 235
    KnxManufacturer_M_LIFEDOMUS KnxManufacturer = 236
    KnxManufacturer_M_EUROTRONIC_TECHNOLOGY_GMBH KnxManufacturer = 237
    KnxManufacturer_M_TCI KnxManufacturer = 238
    KnxManufacturer_M_RISHUN_ELECTRONIC KnxManufacturer = 239
    KnxManufacturer_M_ZIPATO KnxManufacturer = 240
    KnxManufacturer_M_CM_SECURITY_GMBH_AND_CO_KG KnxManufacturer = 241
    KnxManufacturer_M_QING_CABLES KnxManufacturer = 242
    KnxManufacturer_M_LABIO KnxManufacturer = 243
    KnxManufacturer_M_COSTER_TECNOLOGIE_ELETTRONICHE_S_P_A_ KnxManufacturer = 244
    KnxManufacturer_M_E_G_E KnxManufacturer = 245
    KnxManufacturer_M_NETXAUTOMATION KnxManufacturer = 246
    KnxManufacturer_M_TECALOR KnxManufacturer = 247
    KnxManufacturer_M_URMET_ELECTRONICS_HUIZHOU_LTD_ KnxManufacturer = 248
    KnxManufacturer_M_PEIYING_BUILDING_CONTROL KnxManufacturer = 249
    KnxManufacturer_M_BPT_S_P_A__A_SOCIO_UNICO KnxManufacturer = 250
    KnxManufacturer_M_KANONTEC___KANONBUS KnxManufacturer = 251
    KnxManufacturer_M_ISER_TECH KnxManufacturer = 252
    KnxManufacturer_M_FINELINE KnxManufacturer = 253
    KnxManufacturer_M_CP_ELECTRONICS_LTD KnxManufacturer = 254
    KnxManufacturer_M_NIKO_SERVODAN_AS KnxManufacturer = 255
    KnxManufacturer_M_SIMON_309 KnxManufacturer = 256
    KnxManufacturer_M_GM_MODULAR_PVT__LTD_ KnxManufacturer = 257
    KnxManufacturer_M_FU_CHENG_INTELLIGENCE KnxManufacturer = 258
    KnxManufacturer_M_NEXKON KnxManufacturer = 259
    KnxManufacturer_M_FEEL_S_R_L KnxManufacturer = 260
    KnxManufacturer_M_NOT_ASSIGNED_314 KnxManufacturer = 261
    KnxManufacturer_M_SHENZHEN_FANHAI_SANJIANG_ELECTRONICS_CO___LTD_ KnxManufacturer = 262
    KnxManufacturer_M_JIUZHOU_GREEBLE KnxManufacturer = 263
    KnxManufacturer_M_AUMUELLER_AUMATIC_GMBH KnxManufacturer = 264
    KnxManufacturer_M_ETMAN_ELECTRIC KnxManufacturer = 265
    KnxManufacturer_M_BLACK_NOVA KnxManufacturer = 266
    KnxManufacturer_M_ZIDATECH_AG KnxManufacturer = 267
    KnxManufacturer_M_IDGS_BVBA KnxManufacturer = 268
    KnxManufacturer_M_DAKANIMO KnxManufacturer = 269
    KnxManufacturer_M_TREBOR_AUTOMATION_AB KnxManufacturer = 270
    KnxManufacturer_M_SATEL_SP__Z_O_O_ KnxManufacturer = 271
    KnxManufacturer_M_RUSSOUND__INC_ KnxManufacturer = 272
    KnxManufacturer_M_MIDEA_HEATING_AND_VENTILATING_EQUIPMENT_CO_LTD KnxManufacturer = 273
    KnxManufacturer_M_CONSORZIO_TERRANUOVA KnxManufacturer = 274
    KnxManufacturer_M_WOLF_HEIZTECHNIK_GMBH KnxManufacturer = 275
    KnxManufacturer_M_SONTEC KnxManufacturer = 276
    KnxManufacturer_M_BELCOM_CABLES_LTD_ KnxManufacturer = 277
    KnxManufacturer_M_GUANGZHOU_SEAWIN_ELECTRICAL_TECHNOLOGIES_CO___LTD_ KnxManufacturer = 278
    KnxManufacturer_M_ACREL KnxManufacturer = 279
    KnxManufacturer_M_FRANKE_AQUAROTTER_GMBH KnxManufacturer = 280
    KnxManufacturer_M_ORION_SYSTEMS KnxManufacturer = 281
    KnxManufacturer_M_SCHRACK_TECHNIK_GMBH KnxManufacturer = 282
    KnxManufacturer_M_INSPRID KnxManufacturer = 283
    KnxManufacturer_M_SUNRICHER KnxManufacturer = 284
    KnxManufacturer_M_MENRED_AUTOMATION_SYSTEMSHANGHAI_CO__LTD_ KnxManufacturer = 285
    KnxManufacturer_M_AUREX KnxManufacturer = 286
    KnxManufacturer_M_JOSEF_BARTHELME_GMBH_AND_CO__KG KnxManufacturer = 287
    KnxManufacturer_M_ARCHITECTURE_NUMERIQUE KnxManufacturer = 288
    KnxManufacturer_M_UP_GROUP KnxManufacturer = 289
    KnxManufacturer_M_TEKNOS_AVINNO KnxManufacturer = 290
    KnxManufacturer_M_NINGBO_DOOYA_MECHANIC_AND_ELECTRONIC_TECHNOLOGY KnxManufacturer = 291
    KnxManufacturer_M_THERMOKON_SENSORTECHNIK_GMBH KnxManufacturer = 292
    KnxManufacturer_M_BELIMO_AUTOMATION_AG KnxManufacturer = 293
    KnxManufacturer_M_ZEHNDER_GROUP_INTERNATIONAL_AG KnxManufacturer = 294
    KnxManufacturer_M_SKS_KINKEL_ELEKTRONIK KnxManufacturer = 295
    KnxManufacturer_M_ECE_WURMITZER_GMBH KnxManufacturer = 296
    KnxManufacturer_M_LARS KnxManufacturer = 297
    KnxManufacturer_M_URC KnxManufacturer = 298
    KnxManufacturer_M_LIGHTCONTROL KnxManufacturer = 299
    KnxManufacturer_M_SHENZHEN_YM KnxManufacturer = 300
    KnxManufacturer_M_MEAN_WELL_ENTERPRISES_CO__LTD_ KnxManufacturer = 301
    KnxManufacturer_M_OSIX KnxManufacturer = 302
    KnxManufacturer_M_AYPRO_TECHNOLOGY KnxManufacturer = 303
    KnxManufacturer_M_HEFEI_ECOLITE_SOFTWARE KnxManufacturer = 304
    KnxManufacturer_M_ENNO KnxManufacturer = 305
    KnxManufacturer_M_OHOSURE KnxManufacturer = 306
    KnxManufacturer_M_GAREFOWL KnxManufacturer = 307
    KnxManufacturer_M_GEZE KnxManufacturer = 308
    KnxManufacturer_M_LG_ELECTRONICS_INC_ KnxManufacturer = 309
    KnxManufacturer_M_SMC_INTERIORS KnxManufacturer = 310
    KnxManufacturer_M_NOT_ASSIGNED_364 KnxManufacturer = 311
    KnxManufacturer_M_SCS_CABLE KnxManufacturer = 312
    KnxManufacturer_M_HOVAL KnxManufacturer = 313
    KnxManufacturer_M_CANST KnxManufacturer = 314
    KnxManufacturer_M_HANGZHOU_BERLIN KnxManufacturer = 315
    KnxManufacturer_M_EVN_LICHTTECHNIK KnxManufacturer = 316
    KnxManufacturer_M_RUTEC KnxManufacturer = 317
    KnxManufacturer_M_FINDER KnxManufacturer = 318
    KnxManufacturer_M_FUJITSU_GENERAL_LIMITED KnxManufacturer = 319
    KnxManufacturer_M_ZF_FRIEDRICHSHAFEN_AG KnxManufacturer = 320
    KnxManufacturer_M_CREALED KnxManufacturer = 321
    KnxManufacturer_M_MILES_MAGIC_AUTOMATION_PRIVATE_LIMITED KnxManufacturer = 322
    KnxManufacturer_M_EPlus KnxManufacturer = 323
    KnxManufacturer_M_ITALCOND KnxManufacturer = 324
    KnxManufacturer_M_SATION KnxManufacturer = 325
    KnxManufacturer_M_NEWBEST KnxManufacturer = 326
    KnxManufacturer_M_GDS_DIGITAL_SYSTEMS KnxManufacturer = 327
    KnxManufacturer_M_IDDERO KnxManufacturer = 328
    KnxManufacturer_M_MBNLED KnxManufacturer = 329
    KnxManufacturer_M_VITRUM KnxManufacturer = 330
    KnxManufacturer_M_EKEY_BIOMETRIC_SYSTEMS_GMBH KnxManufacturer = 331
    KnxManufacturer_M_AMC KnxManufacturer = 332
    KnxManufacturer_M_TRILUX_GMBH_AND_CO__KG KnxManufacturer = 333
    KnxManufacturer_M_WEXCEDO KnxManufacturer = 334
    KnxManufacturer_M_VEMER_SPA KnxManufacturer = 335
    KnxManufacturer_M_ALEXANDER_BUERKLE_GMBH_AND_CO_KG KnxManufacturer = 336
    KnxManufacturer_M_CITRON KnxManufacturer = 337
    KnxManufacturer_M_SHENZHEN_HEGUANG KnxManufacturer = 338
    KnxManufacturer_M_NOT_ASSIGNED_392 KnxManufacturer = 339
    KnxManufacturer_M_TRANE_B_V_B_A KnxManufacturer = 340
    KnxManufacturer_M_CAREL KnxManufacturer = 341
    KnxManufacturer_M_PROLITE_CONTROLS KnxManufacturer = 342
    KnxManufacturer_M_BOSMER KnxManufacturer = 343
    KnxManufacturer_M_EUCHIPS KnxManufacturer = 344
    KnxManufacturer_M_CONNECT_THINKA_CONNECT KnxManufacturer = 345
    KnxManufacturer_M_PEAKNX_A_DOGAWIST_COMPANY KnxManufacturer = 346
    KnxManufacturer_M_ACEMATIC KnxManufacturer = 347
    KnxManufacturer_M_ELAUSYS KnxManufacturer = 348
    KnxManufacturer_M_ITK_ENGINEERING_AG KnxManufacturer = 349
    KnxManufacturer_M_INTEGRA_METERING_AG KnxManufacturer = 350
    KnxManufacturer_M_FMS_HOSPITALITY_PTE_LTD KnxManufacturer = 351
    KnxManufacturer_M_NUVO KnxManufacturer = 352
    KnxManufacturer_M_U__LUX_GMBH KnxManufacturer = 353
    KnxManufacturer_M_BRUMBERG_LEUCHTEN KnxManufacturer = 354
    KnxManufacturer_M_LIME KnxManufacturer = 355
    KnxManufacturer_M_GREAT_EMPIRE_INTERNATIONAL_GROUP_CO___LTD_ KnxManufacturer = 356
    KnxManufacturer_M_KAVOSHPISHRO_ASIA KnxManufacturer = 357
    KnxManufacturer_M_V2_SPA KnxManufacturer = 358
    KnxManufacturer_M_JOHNSON_CONTROLS KnxManufacturer = 359
    KnxManufacturer_M_ARKUD KnxManufacturer = 360
    KnxManufacturer_M_IRIDIUM_LTD_ KnxManufacturer = 361
    KnxManufacturer_M_BSMART KnxManufacturer = 362
    KnxManufacturer_M_BAB_TECHNOLOGIE_GMBH KnxManufacturer = 363
    KnxManufacturer_M_NICE_SPA KnxManufacturer = 364
    KnxManufacturer_M_REDFISH_GROUP_PTY_LTD KnxManufacturer = 365
    KnxManufacturer_M_SABIANA_SPA KnxManufacturer = 366
    KnxManufacturer_M_UBEE_INTERACTIVE_EUROPE KnxManufacturer = 367
    KnxManufacturer_M_REXEL KnxManufacturer = 368
    KnxManufacturer_M_GES_TEKNIK_A_S_ KnxManufacturer = 369
    KnxManufacturer_M_AVE_S_P_A_ KnxManufacturer = 370
    KnxManufacturer_M_ZHUHAI_LTECH_TECHNOLOGY_CO___LTD_ KnxManufacturer = 371
    KnxManufacturer_M_ARCOM KnxManufacturer = 372
    KnxManufacturer_M_VIA_TECHNOLOGIES__INC_ KnxManufacturer = 373
    KnxManufacturer_M_FEELSMART_ KnxManufacturer = 374
    KnxManufacturer_M_SUPCON KnxManufacturer = 375
    KnxManufacturer_M_MANIC KnxManufacturer = 376
    KnxManufacturer_M_TDE_GMBH KnxManufacturer = 377
    KnxManufacturer_M_NANJING_SHUFAN_INFORMATION_TECHNOLOGY_CO__LTD_ KnxManufacturer = 378
    KnxManufacturer_M_EWTECH KnxManufacturer = 379
    KnxManufacturer_M_KLUGER_AUTOMATION_GMBH KnxManufacturer = 380
    KnxManufacturer_M_JOONGANG_CONTROL KnxManufacturer = 381
    KnxManufacturer_M_GREENCONTROLS_TECHNOLOGY_SDN__BHD_ KnxManufacturer = 382
    KnxManufacturer_M_IME_S_P_A_ KnxManufacturer = 383
    KnxManufacturer_M_SICHUAN_HAODING KnxManufacturer = 384
    KnxManufacturer_M_MINDJAGA_LTD_ KnxManufacturer = 385
    KnxManufacturer_M_RUILI_SMART_CONTROL KnxManufacturer = 386
    KnxManufacturer_M_CODESYS_GMBH KnxManufacturer = 387
    KnxManufacturer_M_MOORGEN_DEUTSCHLAND_GMBH KnxManufacturer = 388
    KnxManufacturer_M_CULLMANN_TECH KnxManufacturer = 389
    KnxManufacturer_M_MERCK_WINDOW_TECHNOLOGIES_B_V_ KnxManufacturer = 390
    KnxManufacturer_M_ABEGO KnxManufacturer = 391
    KnxManufacturer_M_MYGEKKO KnxManufacturer = 392
    KnxManufacturer_M_ERGO3_SARL KnxManufacturer = 393
    KnxManufacturer_M_STMICROELECTRONICS_INTERNATIONAL_N_V_ KnxManufacturer = 394
    KnxManufacturer_M_CJC_SYSTEMS KnxManufacturer = 395
    KnxManufacturer_M_SUDOKU KnxManufacturer = 396
    KnxManufacturer_M_AZ_E_LITE_PTE_LTD KnxManufacturer = 397
    KnxManufacturer_M_ARLIGHT KnxManufacturer = 398
    KnxManufacturer_M_GRUENBECK_WASSERAUFBEREITUNG_GMBH KnxManufacturer = 399
    KnxManufacturer_M_MODULE_ELECTRONIC KnxManufacturer = 400
    KnxManufacturer_M_KOPLAT KnxManufacturer = 401
    KnxManufacturer_M_GUANGZHOU_LETOUR_LIFE_TECHNOLOGY_CO___LTD KnxManufacturer = 402
    KnxManufacturer_M_ILEVIA KnxManufacturer = 403
    KnxManufacturer_M_LN_SYSTEMTEQ KnxManufacturer = 404
    KnxManufacturer_M_HISENSE_SMARTHOME KnxManufacturer = 405
    KnxManufacturer_M_FLINK_AUTOMATION_SYSTEM KnxManufacturer = 406
    KnxManufacturer_M_XXTER_BV KnxManufacturer = 407
    KnxManufacturer_M_LYNXUS_TECHNOLOGY KnxManufacturer = 408
    KnxManufacturer_M_ROBOT_S_A_ KnxManufacturer = 409
    KnxManufacturer_M_SHENZHEN_ATTE_SMART_LIFE_CO__LTD_ KnxManufacturer = 410
    KnxManufacturer_M_NOBLESSE KnxManufacturer = 411
    KnxManufacturer_M_ADVANCED_DEVICES KnxManufacturer = 412
    KnxManufacturer_M_ATRINA_BUILDING_AUTOMATION_CO__LTD KnxManufacturer = 413
    KnxManufacturer_M_GUANGDONG_DAMING_LAFFEY_ELECTRIC_CO___LTD_ KnxManufacturer = 414
    KnxManufacturer_M_WESTERSTRAND_URFABRIK_AB KnxManufacturer = 415
    KnxManufacturer_M_CONTROL4_CORPORATE KnxManufacturer = 416
    KnxManufacturer_M_ONTROL KnxManufacturer = 417
    KnxManufacturer_M_STARNET KnxManufacturer = 418
    KnxManufacturer_M_BETA_CAVI KnxManufacturer = 419
    KnxManufacturer_M_EASEMORE KnxManufacturer = 420
    KnxManufacturer_M_VIVALDI_SRL KnxManufacturer = 421
    KnxManufacturer_M_GREE_ELECTRIC_APPLIANCES_INC__OF_ZHUHAI KnxManufacturer = 422
    KnxManufacturer_M_HWISCON KnxManufacturer = 423
    KnxManufacturer_M_SHANGHAI_ELECON_INTELLIGENT_TECHNOLOGY_CO___LTD_ KnxManufacturer = 424
    KnxManufacturer_M_KAMPMANN KnxManufacturer = 425
    KnxManufacturer_M_IMPOLUX_GMBH_LEDIMAX KnxManufacturer = 426
    KnxManufacturer_M_EVAUX KnxManufacturer = 427
    KnxManufacturer_M_WEBRO_CABLES_AND_CONNECTORS_LIMITED KnxManufacturer = 428
    KnxManufacturer_M_SHANGHAI_E_TECH_SOLUTION KnxManufacturer = 429
    KnxManufacturer_M_GUANGZHOU_HOKO_ELECTRIC_CO__LTD_ KnxManufacturer = 430
    KnxManufacturer_M_LAMMIN_HIGH_TECH_CO__LTD KnxManufacturer = 431
    KnxManufacturer_M_SHENZHEN_MERRYTEK_TECHNOLOGY_CO___LTD KnxManufacturer = 432
    KnxManufacturer_M_I_LUXUS KnxManufacturer = 433
    KnxManufacturer_M_ELMOS_SEMICONDUCTOR_AG KnxManufacturer = 434
    KnxManufacturer_M_EMCOM_TECHNOLOGY_INC KnxManufacturer = 435
    KnxManufacturer_M_PROJECT_INNOVATIONS_GMBH KnxManufacturer = 436
    KnxManufacturer_M_ITC KnxManufacturer = 437
    KnxManufacturer_M_ABB_LV_INSTALLATION_MATERIALS_COMPANY_LTD__BEIJING KnxManufacturer = 438
    KnxManufacturer_M_MAICO KnxManufacturer = 439
    KnxManufacturer_M_ELAN_SRL KnxManufacturer = 440
    KnxManufacturer_M_MINHHA_TECHNOLOGY_CO__LTD KnxManufacturer = 441
    KnxManufacturer_M_ZHEJIANG_TIANJIE_INDUSTRIAL_CORP_ KnxManufacturer = 442
    KnxManufacturer_M_IAUTOMATION_PTY_LIMITED KnxManufacturer = 443
    KnxManufacturer_M_EXTRON KnxManufacturer = 444
    KnxManufacturer_M_FREEDOMPRO KnxManufacturer = 445
    KnxManufacturer_M_ONEHOME KnxManufacturer = 446
    KnxManufacturer_M_EOS_SAUNATECHNIK_GMBH KnxManufacturer = 447
    KnxManufacturer_M_KUSATEK_GMBH KnxManufacturer = 448
    KnxManufacturer_M_EISBAER_SCADA KnxManufacturer = 449
    KnxManufacturer_M_AUTOMATISMI_BENINCA_S_P_A_ KnxManufacturer = 450
    KnxManufacturer_M_BLENDOM KnxManufacturer = 451
    KnxManufacturer_M_MADEL_AIR_TECHNICAL_DIFFUSION KnxManufacturer = 452
    KnxManufacturer_M_NIKO KnxManufacturer = 453
    KnxManufacturer_M_BOSCH_REXROTH_AG KnxManufacturer = 454
    KnxManufacturer_M_CANDM_PRODUCTS KnxManufacturer = 455
    KnxManufacturer_M_HOERMANN_KG_VERKAUFSGESELLSCHAFT KnxManufacturer = 456
    KnxManufacturer_M_SHANGHAI_RAJAYASA_CO__LTD KnxManufacturer = 457
    KnxManufacturer_M_SUZUKI KnxManufacturer = 458
    KnxManufacturer_M_SILENT_GLISS_INTERNATIONAL_LTD_ KnxManufacturer = 459
    KnxManufacturer_M_BEE_CONTROLS_ADGSC_GROUP KnxManufacturer = 460
    KnxManufacturer_M_XDTECGMBH KnxManufacturer = 461
    KnxManufacturer_M_OSRAM KnxManufacturer = 462
    KnxManufacturer_M_LEBENOR KnxManufacturer = 463
    KnxManufacturer_M_AUTOMANENG KnxManufacturer = 464
    KnxManufacturer_M_HONEYWELL_AUTOMATION_SOLUTION_CONTROLCHINA KnxManufacturer = 465
    KnxManufacturer_M_HANGZHOU_BINTHEN_INTELLIGENCE_TECHNOLOGY_CO__LTD KnxManufacturer = 466
    KnxManufacturer_M_ETA_HEIZTECHNIK KnxManufacturer = 467
    KnxManufacturer_M_DIVUS_GMBH KnxManufacturer = 468
    KnxManufacturer_M_NANJING_TAIJIESAI_INTELLIGENT_TECHNOLOGY_CO__LTD_ KnxManufacturer = 469
    KnxManufacturer_M_LUNATONE KnxManufacturer = 470
    KnxManufacturer_M_ZHEJIANG_SCTECH_BUILDING_INTELLIGENT KnxManufacturer = 471
    KnxManufacturer_M_FOSHAN_QITE_TECHNOLOGY_CO___LTD_ KnxManufacturer = 472
    KnxManufacturer_M_NOKE KnxManufacturer = 473
    KnxManufacturer_M_LANDCOM KnxManufacturer = 474
    KnxManufacturer_M_STORK_AS KnxManufacturer = 475
    KnxManufacturer_M_HANGZHOU_SHENDU_TECHNOLOGY_CO___LTD_ KnxManufacturer = 476
    KnxManufacturer_M_COOLAUTOMATION KnxManufacturer = 477
    KnxManufacturer_M_APRSTERN KnxManufacturer = 478
    KnxManufacturer_M_SONNEN KnxManufacturer = 479
    KnxManufacturer_M_DNAKE KnxManufacturer = 480
    KnxManufacturer_M_NEUBERGER_GEBAEUDEAUTOMATION_GMBH KnxManufacturer = 481
    KnxManufacturer_M_STILIGER KnxManufacturer = 482
    KnxManufacturer_M_BERGHOF_AUTOMATION_GMBH KnxManufacturer = 483
    KnxManufacturer_M_TOTAL_AUTOMATION_AND_CONTROLS_GMBH KnxManufacturer = 484
    KnxManufacturer_M_DOVIT KnxManufacturer = 485
    KnxManufacturer_M_INSTALIGHTING_GMBH KnxManufacturer = 486
    KnxManufacturer_M_UNI_TEC KnxManufacturer = 487
    KnxManufacturer_M_CASATUNES KnxManufacturer = 488
    KnxManufacturer_M_EMT KnxManufacturer = 489
    KnxManufacturer_M_SENFFICIENT KnxManufacturer = 490
    KnxManufacturer_M_AUROLITE_ELECTRICAL_PANYU_GUANGZHOU_LIMITED KnxManufacturer = 491
    KnxManufacturer_M_ABB_XIAMEN_SMART_TECHNOLOGY_CO___LTD_ KnxManufacturer = 492
    KnxManufacturer_M_SAMSON_ELECTRIC_WIRE KnxManufacturer = 493
    KnxManufacturer_M_T_TOUCHING KnxManufacturer = 494
    KnxManufacturer_M_CORE_SMART_HOME KnxManufacturer = 495
    KnxManufacturer_M_GREENCONNECT_SOLUTIONS_SA KnxManufacturer = 496
    KnxManufacturer_M_ELETTRONICA_CONDUTTORI KnxManufacturer = 497
    KnxManufacturer_M_MKFC KnxManufacturer = 498
    KnxManufacturer_M_AUTOMATIONPlus KnxManufacturer = 499
    KnxManufacturer_M_BLUE_AND_RED KnxManufacturer = 500
    KnxManufacturer_M_FROGBLUE KnxManufacturer = 501
    KnxManufacturer_M_SAVESOR KnxManufacturer = 502
    KnxManufacturer_M_APP_TECH KnxManufacturer = 503
    KnxManufacturer_M_SENSORTEC_AG KnxManufacturer = 504
    KnxManufacturer_M_NYSA_TECHNOLOGY_AND_SOLUTIONS KnxManufacturer = 505
    KnxManufacturer_M_FARADITE KnxManufacturer = 506
    KnxManufacturer_M_OPTIMUS KnxManufacturer = 507
    KnxManufacturer_M_KTS_S_R_L_ KnxManufacturer = 508
    KnxManufacturer_M_RAMCRO_SPA KnxManufacturer = 509
    KnxManufacturer_M_WUHAN_WISECREATE_UNIVERSE_TECHNOLOGY_CO___LTD KnxManufacturer = 510
    KnxManufacturer_M_BEMI_SMART_HOME_LTD KnxManufacturer = 511
    KnxManufacturer_M_ARDOMUS KnxManufacturer = 512
    KnxManufacturer_M_CHANGXING KnxManufacturer = 513
    KnxManufacturer_M_E_CONTROLS KnxManufacturer = 514
    KnxManufacturer_M_AIB_TECHNOLOGY KnxManufacturer = 515
    KnxManufacturer_M_NVC KnxManufacturer = 516
    KnxManufacturer_M_KBOX KnxManufacturer = 517
    KnxManufacturer_M_CNS KnxManufacturer = 518
    KnxManufacturer_M_TYBA KnxManufacturer = 519
    KnxManufacturer_M_ATREL KnxManufacturer = 520
    KnxManufacturer_M_SIMON_ELECTRIC_CHINA_CO___LTD KnxManufacturer = 521
    KnxManufacturer_M_KORDZ_GROUP KnxManufacturer = 522
    KnxManufacturer_M_ND_ELECTRIC KnxManufacturer = 523
    KnxManufacturer_M_CONTROLIUM KnxManufacturer = 524
    KnxManufacturer_M_FAMO_GMBH_AND_CO__KG KnxManufacturer = 525
    KnxManufacturer_M_CDN_SMART KnxManufacturer = 526
    KnxManufacturer_M_HESTON KnxManufacturer = 527
    KnxManufacturer_M_ESLA_CONEXIONES_S_L_ KnxManufacturer = 528
    KnxManufacturer_M_WEISHAUPT KnxManufacturer = 529
    KnxManufacturer_M_ASTRUM_TECHNOLOGY KnxManufacturer = 530
    KnxManufacturer_M_WUERTH_ELEKTRONIK_STELVIO_KONTEK_S_P_A_ KnxManufacturer = 531
    KnxManufacturer_M_NANOTECO_CORPORATION KnxManufacturer = 532
    KnxManufacturer_M_NIETIAN KnxManufacturer = 533
    KnxManufacturer_M_SUMSIR KnxManufacturer = 534
    KnxManufacturer_M_ORBIS_TECNOLOGIA_ELECTRICA_SA KnxManufacturer = 535
    KnxManufacturer_M_ABB___RESERVED KnxManufacturer = 536
    KnxManufacturer_M_BUSCH_JAEGER_ELEKTRO___RESERVED KnxManufacturer = 537
)


func (e KnxManufacturer) Number() uint16 {
    switch e  {
        case 0: { /* '0' */
            return 0
        }
        case 1: { /* '1' */
            return 1
        }
        case 10: { /* '10' */
            return 11
        }
        case 100: { /* '100' */
            return 140
        }
        case 101: { /* '101' */
            return 141
        }
        case 102: { /* '102' */
            return 142
        }
        case 103: { /* '103' */
            return 143
        }
        case 104: { /* '104' */
            return 144
        }
        case 105: { /* '105' */
            return 145
        }
        case 106: { /* '106' */
            return 146
        }
        case 107: { /* '107' */
            return 147
        }
        case 108: { /* '108' */
            return 148
        }
        case 109: { /* '109' */
            return 149
        }
        case 11: { /* '11' */
            return 12
        }
        case 110: { /* '110' */
            return 150
        }
        case 111: { /* '111' */
            return 151
        }
        case 112: { /* '112' */
            return 152
        }
        case 113: { /* '113' */
            return 153
        }
        case 114: { /* '114' */
            return 154
        }
        case 115: { /* '115' */
            return 155
        }
        case 116: { /* '116' */
            return 156
        }
        case 117: { /* '117' */
            return 157
        }
        case 118: { /* '118' */
            return 158
        }
        case 119: { /* '119' */
            return 159
        }
        case 12: { /* '12' */
            return 14
        }
        case 120: { /* '120' */
            return 160
        }
        case 121: { /* '121' */
            return 161
        }
        case 122: { /* '122' */
            return 162
        }
        case 123: { /* '123' */
            return 163
        }
        case 124: { /* '124' */
            return 164
        }
        case 125: { /* '125' */
            return 165
        }
        case 126: { /* '126' */
            return 166
        }
        case 127: { /* '127' */
            return 167
        }
        case 128: { /* '128' */
            return 168
        }
        case 129: { /* '129' */
            return 169
        }
        case 13: { /* '13' */
            return 22
        }
        case 130: { /* '130' */
            return 170
        }
        case 131: { /* '131' */
            return 171
        }
        case 132: { /* '132' */
            return 172
        }
        case 133: { /* '133' */
            return 173
        }
        case 134: { /* '134' */
            return 174
        }
        case 135: { /* '135' */
            return 175
        }
        case 136: { /* '136' */
            return 176
        }
        case 137: { /* '137' */
            return 177
        }
        case 138: { /* '138' */
            return 178
        }
        case 139: { /* '139' */
            return 179
        }
        case 14: { /* '14' */
            return 24
        }
        case 140: { /* '140' */
            return 180
        }
        case 141: { /* '141' */
            return 181
        }
        case 142: { /* '142' */
            return 182
        }
        case 143: { /* '143' */
            return 183
        }
        case 144: { /* '144' */
            return 184
        }
        case 145: { /* '145' */
            return 185
        }
        case 146: { /* '146' */
            return 186
        }
        case 147: { /* '147' */
            return 187
        }
        case 148: { /* '148' */
            return 188
        }
        case 149: { /* '149' */
            return 189
        }
        case 15: { /* '15' */
            return 25
        }
        case 150: { /* '150' */
            return 190
        }
        case 151: { /* '151' */
            return 191
        }
        case 152: { /* '152' */
            return 192
        }
        case 153: { /* '153' */
            return 193
        }
        case 154: { /* '154' */
            return 194
        }
        case 155: { /* '155' */
            return 195
        }
        case 156: { /* '156' */
            return 196
        }
        case 157: { /* '157' */
            return 197
        }
        case 158: { /* '158' */
            return 198
        }
        case 159: { /* '159' */
            return 199
        }
        case 16: { /* '16' */
            return 27
        }
        case 160: { /* '160' */
            return 200
        }
        case 161: { /* '161' */
            return 201
        }
        case 162: { /* '162' */
            return 202
        }
        case 163: { /* '163' */
            return 204
        }
        case 164: { /* '164' */
            return 205
        }
        case 165: { /* '165' */
            return 206
        }
        case 166: { /* '166' */
            return 207
        }
        case 167: { /* '167' */
            return 208
        }
        case 168: { /* '168' */
            return 209
        }
        case 169: { /* '169' */
            return 210
        }
        case 17: { /* '17' */
            return 28
        }
        case 170: { /* '170' */
            return 211
        }
        case 171: { /* '171' */
            return 214
        }
        case 172: { /* '172' */
            return 215
        }
        case 173: { /* '173' */
            return 216
        }
        case 174: { /* '174' */
            return 217
        }
        case 175: { /* '175' */
            return 218
        }
        case 176: { /* '176' */
            return 219
        }
        case 177: { /* '177' */
            return 220
        }
        case 178: { /* '178' */
            return 222
        }
        case 179: { /* '179' */
            return 223
        }
        case 18: { /* '18' */
            return 29
        }
        case 180: { /* '180' */
            return 225
        }
        case 181: { /* '181' */
            return 227
        }
        case 182: { /* '182' */
            return 228
        }
        case 183: { /* '183' */
            return 232
        }
        case 184: { /* '184' */
            return 233
        }
        case 185: { /* '185' */
            return 234
        }
        case 186: { /* '186' */
            return 235
        }
        case 187: { /* '187' */
            return 237
        }
        case 188: { /* '188' */
            return 238
        }
        case 189: { /* '189' */
            return 239
        }
        case 19: { /* '19' */
            return 30
        }
        case 190: { /* '190' */
            return 240
        }
        case 191: { /* '191' */
            return 241
        }
        case 192: { /* '192' */
            return 242
        }
        case 193: { /* '193' */
            return 244
        }
        case 194: { /* '194' */
            return 245
        }
        case 195: { /* '195' */
            return 246
        }
        case 196: { /* '196' */
            return 248
        }
        case 197: { /* '197' */
            return 249
        }
        case 198: { /* '198' */
            return 250
        }
        case 199: { /* '199' */
            return 251
        }
        case 2: { /* '2' */
            return 2
        }
        case 20: { /* '20' */
            return 31
        }
        case 200: { /* '200' */
            return 252
        }
        case 201: { /* '201' */
            return 253
        }
        case 202: { /* '202' */
            return 254
        }
        case 203: { /* '203' */
            return 256
        }
        case 204: { /* '204' */
            return 257
        }
        case 205: { /* '205' */
            return 258
        }
        case 206: { /* '206' */
            return 259
        }
        case 207: { /* '207' */
            return 260
        }
        case 208: { /* '208' */
            return 261
        }
        case 209: { /* '209' */
            return 262
        }
        case 21: { /* '21' */
            return 32
        }
        case 210: { /* '210' */
            return 263
        }
        case 211: { /* '211' */
            return 264
        }
        case 212: { /* '212' */
            return 265
        }
        case 213: { /* '213' */
            return 266
        }
        case 214: { /* '214' */
            return 267
        }
        case 215: { /* '215' */
            return 268
        }
        case 216: { /* '216' */
            return 269
        }
        case 217: { /* '217' */
            return 270
        }
        case 218: { /* '218' */
            return 271
        }
        case 219: { /* '219' */
            return 272
        }
        case 22: { /* '22' */
            return 33
        }
        case 220: { /* '220' */
            return 273
        }
        case 221: { /* '221' */
            return 274
        }
        case 222: { /* '222' */
            return 275
        }
        case 223: { /* '223' */
            return 276
        }
        case 224: { /* '224' */
            return 277
        }
        case 225: { /* '225' */
            return 278
        }
        case 226: { /* '226' */
            return 279
        }
        case 227: { /* '227' */
            return 280
        }
        case 228: { /* '228' */
            return 281
        }
        case 229: { /* '229' */
            return 282
        }
        case 23: { /* '23' */
            return 34
        }
        case 230: { /* '230' */
            return 283
        }
        case 231: { /* '231' */
            return 284
        }
        case 232: { /* '232' */
            return 285
        }
        case 233: { /* '233' */
            return 286
        }
        case 234: { /* '234' */
            return 287
        }
        case 235: { /* '235' */
            return 288
        }
        case 236: { /* '236' */
            return 289
        }
        case 237: { /* '237' */
            return 290
        }
        case 238: { /* '238' */
            return 291
        }
        case 239: { /* '239' */
            return 292
        }
        case 24: { /* '24' */
            return 36
        }
        case 240: { /* '240' */
            return 293
        }
        case 241: { /* '241' */
            return 294
        }
        case 242: { /* '242' */
            return 295
        }
        case 243: { /* '243' */
            return 296
        }
        case 244: { /* '244' */
            return 297
        }
        case 245: { /* '245' */
            return 298
        }
        case 246: { /* '246' */
            return 299
        }
        case 247: { /* '247' */
            return 300
        }
        case 248: { /* '248' */
            return 301
        }
        case 249: { /* '249' */
            return 302
        }
        case 25: { /* '25' */
            return 37
        }
        case 250: { /* '250' */
            return 303
        }
        case 251: { /* '251' */
            return 304
        }
        case 252: { /* '252' */
            return 305
        }
        case 253: { /* '253' */
            return 306
        }
        case 254: { /* '254' */
            return 307
        }
        case 255: { /* '255' */
            return 308
        }
        case 256: { /* '256' */
            return 309
        }
        case 257: { /* '257' */
            return 310
        }
        case 258: { /* '258' */
            return 311
        }
        case 259: { /* '259' */
            return 312
        }
        case 26: { /* '26' */
            return 41
        }
        case 260: { /* '260' */
            return 313
        }
        case 261: { /* '261' */
            return 314
        }
        case 262: { /* '262' */
            return 315
        }
        case 263: { /* '263' */
            return 316
        }
        case 264: { /* '264' */
            return 317
        }
        case 265: { /* '265' */
            return 318
        }
        case 266: { /* '266' */
            return 319
        }
        case 267: { /* '267' */
            return 320
        }
        case 268: { /* '268' */
            return 321
        }
        case 269: { /* '269' */
            return 322
        }
        case 27: { /* '27' */
            return 42
        }
        case 270: { /* '270' */
            return 323
        }
        case 271: { /* '271' */
            return 324
        }
        case 272: { /* '272' */
            return 325
        }
        case 273: { /* '273' */
            return 326
        }
        case 274: { /* '274' */
            return 327
        }
        case 275: { /* '275' */
            return 328
        }
        case 276: { /* '276' */
            return 329
        }
        case 277: { /* '277' */
            return 330
        }
        case 278: { /* '278' */
            return 331
        }
        case 279: { /* '279' */
            return 332
        }
        case 28: { /* '28' */
            return 44
        }
        case 280: { /* '280' */
            return 333
        }
        case 281: { /* '281' */
            return 334
        }
        case 282: { /* '282' */
            return 335
        }
        case 283: { /* '283' */
            return 336
        }
        case 284: { /* '284' */
            return 337
        }
        case 285: { /* '285' */
            return 338
        }
        case 286: { /* '286' */
            return 339
        }
        case 287: { /* '287' */
            return 340
        }
        case 288: { /* '288' */
            return 341
        }
        case 289: { /* '289' */
            return 342
        }
        case 29: { /* '29' */
            return 45
        }
        case 290: { /* '290' */
            return 343
        }
        case 291: { /* '291' */
            return 344
        }
        case 292: { /* '292' */
            return 345
        }
        case 293: { /* '293' */
            return 346
        }
        case 294: { /* '294' */
            return 347
        }
        case 295: { /* '295' */
            return 348
        }
        case 296: { /* '296' */
            return 349
        }
        case 297: { /* '297' */
            return 350
        }
        case 298: { /* '298' */
            return 351
        }
        case 299: { /* '299' */
            return 352
        }
        case 3: { /* '3' */
            return 4
        }
        case 30: { /* '30' */
            return 46
        }
        case 300: { /* '300' */
            return 353
        }
        case 301: { /* '301' */
            return 354
        }
        case 302: { /* '302' */
            return 355
        }
        case 303: { /* '303' */
            return 356
        }
        case 304: { /* '304' */
            return 357
        }
        case 305: { /* '305' */
            return 358
        }
        case 306: { /* '306' */
            return 359
        }
        case 307: { /* '307' */
            return 360
        }
        case 308: { /* '308' */
            return 361
        }
        case 309: { /* '309' */
            return 362
        }
        case 31: { /* '31' */
            return 49
        }
        case 310: { /* '310' */
            return 363
        }
        case 311: { /* '311' */
            return 364
        }
        case 312: { /* '312' */
            return 365
        }
        case 313: { /* '313' */
            return 366
        }
        case 314: { /* '314' */
            return 367
        }
        case 315: { /* '315' */
            return 368
        }
        case 316: { /* '316' */
            return 369
        }
        case 317: { /* '317' */
            return 370
        }
        case 318: { /* '318' */
            return 371
        }
        case 319: { /* '319' */
            return 372
        }
        case 32: { /* '32' */
            return 52
        }
        case 320: { /* '320' */
            return 373
        }
        case 321: { /* '321' */
            return 374
        }
        case 322: { /* '322' */
            return 375
        }
        case 323: { /* '323' */
            return 376
        }
        case 324: { /* '324' */
            return 377
        }
        case 325: { /* '325' */
            return 378
        }
        case 326: { /* '326' */
            return 379
        }
        case 327: { /* '327' */
            return 380
        }
        case 328: { /* '328' */
            return 381
        }
        case 329: { /* '329' */
            return 382
        }
        case 33: { /* '33' */
            return 53
        }
        case 330: { /* '330' */
            return 383
        }
        case 331: { /* '331' */
            return 384
        }
        case 332: { /* '332' */
            return 385
        }
        case 333: { /* '333' */
            return 386
        }
        case 334: { /* '334' */
            return 387
        }
        case 335: { /* '335' */
            return 388
        }
        case 336: { /* '336' */
            return 389
        }
        case 337: { /* '337' */
            return 390
        }
        case 338: { /* '338' */
            return 391
        }
        case 339: { /* '339' */
            return 392
        }
        case 34: { /* '34' */
            return 55
        }
        case 340: { /* '340' */
            return 393
        }
        case 341: { /* '341' */
            return 394
        }
        case 342: { /* '342' */
            return 395
        }
        case 343: { /* '343' */
            return 396
        }
        case 344: { /* '344' */
            return 397
        }
        case 345: { /* '345' */
            return 398
        }
        case 346: { /* '346' */
            return 399
        }
        case 347: { /* '347' */
            return 400
        }
        case 348: { /* '348' */
            return 401
        }
        case 349: { /* '349' */
            return 402
        }
        case 35: { /* '35' */
            return 57
        }
        case 350: { /* '350' */
            return 403
        }
        case 351: { /* '351' */
            return 404
        }
        case 352: { /* '352' */
            return 405
        }
        case 353: { /* '353' */
            return 406
        }
        case 354: { /* '354' */
            return 407
        }
        case 355: { /* '355' */
            return 408
        }
        case 356: { /* '356' */
            return 409
        }
        case 357: { /* '357' */
            return 410
        }
        case 358: { /* '358' */
            return 411
        }
        case 359: { /* '359' */
            return 412
        }
        case 36: { /* '36' */
            return 61
        }
        case 360: { /* '360' */
            return 413
        }
        case 361: { /* '361' */
            return 414
        }
        case 362: { /* '362' */
            return 415
        }
        case 363: { /* '363' */
            return 416
        }
        case 364: { /* '364' */
            return 417
        }
        case 365: { /* '365' */
            return 418
        }
        case 366: { /* '366' */
            return 419
        }
        case 367: { /* '367' */
            return 420
        }
        case 368: { /* '368' */
            return 421
        }
        case 369: { /* '369' */
            return 422
        }
        case 37: { /* '37' */
            return 62
        }
        case 370: { /* '370' */
            return 423
        }
        case 371: { /* '371' */
            return 424
        }
        case 372: { /* '372' */
            return 425
        }
        case 373: { /* '373' */
            return 426
        }
        case 374: { /* '374' */
            return 427
        }
        case 375: { /* '375' */
            return 428
        }
        case 376: { /* '376' */
            return 429
        }
        case 377: { /* '377' */
            return 430
        }
        case 378: { /* '378' */
            return 431
        }
        case 379: { /* '379' */
            return 432
        }
        case 38: { /* '38' */
            return 66
        }
        case 380: { /* '380' */
            return 433
        }
        case 381: { /* '381' */
            return 434
        }
        case 382: { /* '382' */
            return 435
        }
        case 383: { /* '383' */
            return 436
        }
        case 384: { /* '384' */
            return 437
        }
        case 385: { /* '385' */
            return 438
        }
        case 386: { /* '386' */
            return 439
        }
        case 387: { /* '387' */
            return 440
        }
        case 388: { /* '388' */
            return 441
        }
        case 389: { /* '389' */
            return 442
        }
        case 39: { /* '39' */
            return 67
        }
        case 390: { /* '390' */
            return 443
        }
        case 391: { /* '391' */
            return 444
        }
        case 392: { /* '392' */
            return 445
        }
        case 393: { /* '393' */
            return 446
        }
        case 394: { /* '394' */
            return 447
        }
        case 395: { /* '395' */
            return 448
        }
        case 396: { /* '396' */
            return 449
        }
        case 397: { /* '397' */
            return 451
        }
        case 398: { /* '398' */
            return 452
        }
        case 399: { /* '399' */
            return 453
        }
        case 4: { /* '4' */
            return 5
        }
        case 40: { /* '40' */
            return 69
        }
        case 400: { /* '400' */
            return 454
        }
        case 401: { /* '401' */
            return 455
        }
        case 402: { /* '402' */
            return 456
        }
        case 403: { /* '403' */
            return 457
        }
        case 404: { /* '404' */
            return 458
        }
        case 405: { /* '405' */
            return 459
        }
        case 406: { /* '406' */
            return 460
        }
        case 407: { /* '407' */
            return 461
        }
        case 408: { /* '408' */
            return 462
        }
        case 409: { /* '409' */
            return 463
        }
        case 41: { /* '41' */
            return 71
        }
        case 410: { /* '410' */
            return 464
        }
        case 411: { /* '411' */
            return 465
        }
        case 412: { /* '412' */
            return 466
        }
        case 413: { /* '413' */
            return 467
        }
        case 414: { /* '414' */
            return 468
        }
        case 415: { /* '415' */
            return 469
        }
        case 416: { /* '416' */
            return 470
        }
        case 417: { /* '417' */
            return 471
        }
        case 418: { /* '418' */
            return 472
        }
        case 419: { /* '419' */
            return 473
        }
        case 42: { /* '42' */
            return 72
        }
        case 420: { /* '420' */
            return 474
        }
        case 421: { /* '421' */
            return 475
        }
        case 422: { /* '422' */
            return 476
        }
        case 423: { /* '423' */
            return 477
        }
        case 424: { /* '424' */
            return 478
        }
        case 425: { /* '425' */
            return 479
        }
        case 426: { /* '426' */
            return 480
        }
        case 427: { /* '427' */
            return 481
        }
        case 428: { /* '428' */
            return 482
        }
        case 429: { /* '429' */
            return 483
        }
        case 43: { /* '43' */
            return 73
        }
        case 430: { /* '430' */
            return 484
        }
        case 431: { /* '431' */
            return 485
        }
        case 432: { /* '432' */
            return 486
        }
        case 433: { /* '433' */
            return 487
        }
        case 434: { /* '434' */
            return 488
        }
        case 435: { /* '435' */
            return 489
        }
        case 436: { /* '436' */
            return 490
        }
        case 437: { /* '437' */
            return 491
        }
        case 438: { /* '438' */
            return 492
        }
        case 439: { /* '439' */
            return 493
        }
        case 44: { /* '44' */
            return 75
        }
        case 440: { /* '440' */
            return 495
        }
        case 441: { /* '441' */
            return 496
        }
        case 442: { /* '442' */
            return 497
        }
        case 443: { /* '443' */
            return 498
        }
        case 444: { /* '444' */
            return 499
        }
        case 445: { /* '445' */
            return 500
        }
        case 446: { /* '446' */
            return 501
        }
        case 447: { /* '447' */
            return 502
        }
        case 448: { /* '448' */
            return 503
        }
        case 449: { /* '449' */
            return 504
        }
        case 45: { /* '45' */
            return 76
        }
        case 450: { /* '450' */
            return 505
        }
        case 451: { /* '451' */
            return 506
        }
        case 452: { /* '452' */
            return 507
        }
        case 453: { /* '453' */
            return 508
        }
        case 454: { /* '454' */
            return 509
        }
        case 455: { /* '455' */
            return 512
        }
        case 456: { /* '456' */
            return 513
        }
        case 457: { /* '457' */
            return 514
        }
        case 458: { /* '458' */
            return 515
        }
        case 459: { /* '459' */
            return 516
        }
        case 46: { /* '46' */
            return 78
        }
        case 460: { /* '460' */
            return 517
        }
        case 461: { /* '461' */
            return 518
        }
        case 462: { /* '462' */
            return 519
        }
        case 463: { /* '463' */
            return 520
        }
        case 464: { /* '464' */
            return 521
        }
        case 465: { /* '465' */
            return 522
        }
        case 466: { /* '466' */
            return 523
        }
        case 467: { /* '467' */
            return 524
        }
        case 468: { /* '468' */
            return 525
        }
        case 469: { /* '469' */
            return 526
        }
        case 47: { /* '47' */
            return 80
        }
        case 470: { /* '470' */
            return 527
        }
        case 471: { /* '471' */
            return 528
        }
        case 472: { /* '472' */
            return 529
        }
        case 473: { /* '473' */
            return 530
        }
        case 474: { /* '474' */
            return 531
        }
        case 475: { /* '475' */
            return 532
        }
        case 476: { /* '476' */
            return 533
        }
        case 477: { /* '477' */
            return 534
        }
        case 478: { /* '478' */
            return 535
        }
        case 479: { /* '479' */
            return 536
        }
        case 48: { /* '48' */
            return 81
        }
        case 480: { /* '480' */
            return 537
        }
        case 481: { /* '481' */
            return 538
        }
        case 482: { /* '482' */
            return 539
        }
        case 483: { /* '483' */
            return 540
        }
        case 484: { /* '484' */
            return 541
        }
        case 485: { /* '485' */
            return 542
        }
        case 486: { /* '486' */
            return 543
        }
        case 487: { /* '487' */
            return 544
        }
        case 488: { /* '488' */
            return 545
        }
        case 489: { /* '489' */
            return 546
        }
        case 49: { /* '49' */
            return 82
        }
        case 490: { /* '490' */
            return 547
        }
        case 491: { /* '491' */
            return 548
        }
        case 492: { /* '492' */
            return 549
        }
        case 493: { /* '493' */
            return 550
        }
        case 494: { /* '494' */
            return 551
        }
        case 495: { /* '495' */
            return 552
        }
        case 496: { /* '496' */
            return 553
        }
        case 497: { /* '497' */
            return 554
        }
        case 498: { /* '498' */
            return 555
        }
        case 499: { /* '499' */
            return 556
        }
        case 5: { /* '5' */
            return 6
        }
        case 50: { /* '50' */
            return 83
        }
        case 500: { /* '500' */
            return 557
        }
        case 501: { /* '501' */
            return 558
        }
        case 502: { /* '502' */
            return 559
        }
        case 503: { /* '503' */
            return 560
        }
        case 504: { /* '504' */
            return 561
        }
        case 505: { /* '505' */
            return 562
        }
        case 506: { /* '506' */
            return 563
        }
        case 507: { /* '507' */
            return 564
        }
        case 508: { /* '508' */
            return 565
        }
        case 509: { /* '509' */
            return 566
        }
        case 51: { /* '51' */
            return 85
        }
        case 510: { /* '510' */
            return 567
        }
        case 511: { /* '511' */
            return 568
        }
        case 512: { /* '512' */
            return 569
        }
        case 513: { /* '513' */
            return 570
        }
        case 514: { /* '514' */
            return 571
        }
        case 515: { /* '515' */
            return 572
        }
        case 516: { /* '516' */
            return 573
        }
        case 517: { /* '517' */
            return 574
        }
        case 518: { /* '518' */
            return 575
        }
        case 519: { /* '519' */
            return 576
        }
        case 52: { /* '52' */
            return 89
        }
        case 520: { /* '520' */
            return 577
        }
        case 521: { /* '521' */
            return 578
        }
        case 522: { /* '522' */
            return 579
        }
        case 523: { /* '523' */
            return 580
        }
        case 524: { /* '524' */
            return 581
        }
        case 525: { /* '525' */
            return 582
        }
        case 526: { /* '526' */
            return 583
        }
        case 527: { /* '527' */
            return 584
        }
        case 528: { /* '528' */
            return 585
        }
        case 529: { /* '529' */
            return 586
        }
        case 53: { /* '53' */
            return 90
        }
        case 530: { /* '530' */
            return 587
        }
        case 531: { /* '531' */
            return 588
        }
        case 532: { /* '532' */
            return 589
        }
        case 533: { /* '533' */
            return 590
        }
        case 534: { /* '534' */
            return 591
        }
        case 535: { /* '535' */
            return 592
        }
        case 536: { /* '536' */
            return 43954
        }
        case 537: { /* '537' */
            return 43959
        }
        case 54: { /* '54' */
            return 92
        }
        case 55: { /* '55' */
            return 93
        }
        case 56: { /* '56' */
            return 94
        }
        case 57: { /* '57' */
            return 95
        }
        case 58: { /* '58' */
            return 97
        }
        case 59: { /* '59' */
            return 98
        }
        case 6: { /* '6' */
            return 7
        }
        case 60: { /* '60' */
            return 99
        }
        case 61: { /* '61' */
            return 100
        }
        case 62: { /* '62' */
            return 101
        }
        case 63: { /* '63' */
            return 102
        }
        case 64: { /* '64' */
            return 104
        }
        case 65: { /* '65' */
            return 105
        }
        case 66: { /* '66' */
            return 106
        }
        case 67: { /* '67' */
            return 107
        }
        case 68: { /* '68' */
            return 108
        }
        case 69: { /* '69' */
            return 109
        }
        case 7: { /* '7' */
            return 8
        }
        case 70: { /* '70' */
            return 110
        }
        case 71: { /* '71' */
            return 111
        }
        case 72: { /* '72' */
            return 112
        }
        case 73: { /* '73' */
            return 113
        }
        case 74: { /* '74' */
            return 114
        }
        case 75: { /* '75' */
            return 115
        }
        case 76: { /* '76' */
            return 116
        }
        case 77: { /* '77' */
            return 117
        }
        case 78: { /* '78' */
            return 118
        }
        case 79: { /* '79' */
            return 119
        }
        case 8: { /* '8' */
            return 9
        }
        case 80: { /* '80' */
            return 120
        }
        case 81: { /* '81' */
            return 121
        }
        case 82: { /* '82' */
            return 122
        }
        case 83: { /* '83' */
            return 123
        }
        case 84: { /* '84' */
            return 124
        }
        case 85: { /* '85' */
            return 125
        }
        case 86: { /* '86' */
            return 126
        }
        case 87: { /* '87' */
            return 127
        }
        case 88: { /* '88' */
            return 128
        }
        case 89: { /* '89' */
            return 129
        }
        case 9: { /* '9' */
            return 10
        }
        case 90: { /* '90' */
            return 130
        }
        case 91: { /* '91' */
            return 131
        }
        case 92: { /* '92' */
            return 132
        }
        case 93: { /* '93' */
            return 133
        }
        case 94: { /* '94' */
            return 134
        }
        case 95: { /* '95' */
            return 135
        }
        case 96: { /* '96' */
            return 136
        }
        case 97: { /* '97' */
            return 137
        }
        case 98: { /* '98' */
            return 138
        }
        case 99: { /* '99' */
            return 139
        }
        default: {
            return 0
        }
    }
}

func (e KnxManufacturer) Name() string {
    switch e  {
        case 0: { /* '0' */
            return "Unknown Manufacturer"
        }
        case 1: { /* '1' */
            return "Siemens"
        }
        case 10: { /* '10' */
            return "LEGRAND Appareillage électrique"
        }
        case 100: { /* '100' */
            return "Pulse Technologies"
        }
        case 101: { /* '101' */
            return "Crestron"
        }
        case 102: { /* '102' */
            return "STEINEL professional"
        }
        case 103: { /* '103' */
            return "BILTON LED Lighting"
        }
        case 104: { /* '104' */
            return "denro AG"
        }
        case 105: { /* '105' */
            return "GePro"
        }
        case 106: { /* '106' */
            return "preussen automation"
        }
        case 107: { /* '107' */
            return "Zoppas Industries"
        }
        case 108: { /* '108' */
            return "MACTECH"
        }
        case 109: { /* '109' */
            return "TECHNO-TREND"
        }
        case 11: { /* '11' */
            return "Merten"
        }
        case 110: { /* '110' */
            return "FS Cables"
        }
        case 111: { /* '111' */
            return "Delta Dore"
        }
        case 112: { /* '112' */
            return "Eissound"
        }
        case 113: { /* '113' */
            return "Cisco"
        }
        case 114: { /* '114' */
            return "Dinuy"
        }
        case 115: { /* '115' */
            return "iKNiX"
        }
        case 116: { /* '116' */
            return "Rademacher Geräte-Elektronik GmbH"
        }
        case 117: { /* '117' */
            return "EGi Electroacustica General Iberica"
        }
        case 118: { /* '118' */
            return "Bes – Ingenium"
        }
        case 119: { /* '119' */
            return "ElabNET"
        }
        case 12: { /* '12' */
            return "ABB SpA-SACE Division"
        }
        case 120: { /* '120' */
            return "Blumotix"
        }
        case 121: { /* '121' */
            return "Hunter Douglas"
        }
        case 122: { /* '122' */
            return "APRICUM"
        }
        case 123: { /* '123' */
            return "TIANSU Automation"
        }
        case 124: { /* '124' */
            return "Bubendorff"
        }
        case 125: { /* '125' */
            return "MBS GmbH"
        }
        case 126: { /* '126' */
            return "Enertex Bayern GmbH"
        }
        case 127: { /* '127' */
            return "BMS"
        }
        case 128: { /* '128' */
            return "Sinapsi"
        }
        case 129: { /* '129' */
            return "Embedded Systems SIA"
        }
        case 13: { /* '13' */
            return "Siedle & Söhne"
        }
        case 130: { /* '130' */
            return "KNX1"
        }
        case 131: { /* '131' */
            return "Tokka"
        }
        case 132: { /* '132' */
            return "NanoSense"
        }
        case 133: { /* '133' */
            return "PEAR Automation GmbH"
        }
        case 134: { /* '134' */
            return "DGA"
        }
        case 135: { /* '135' */
            return "Lutron"
        }
        case 136: { /* '136' */
            return "AIRZONE – ALTRA"
        }
        case 137: { /* '137' */
            return "Lithoss Design Switches"
        }
        case 138: { /* '138' */
            return "3ATEL"
        }
        case 139: { /* '139' */
            return "Philips Controls"
        }
        case 14: { /* '14' */
            return "Eberle"
        }
        case 140: { /* '140' */
            return "VELUX A/S"
        }
        case 141: { /* '141' */
            return "LOYTEC"
        }
        case 142: { /* '142' */
            return "Ekinex S.p.A."
        }
        case 143: { /* '143' */
            return "SIRLAN Technologies"
        }
        case 144: { /* '144' */
            return "ProKNX SAS"
        }
        case 145: { /* '145' */
            return "IT GmbH"
        }
        case 146: { /* '146' */
            return "RENSON"
        }
        case 147: { /* '147' */
            return "HEP Group"
        }
        case 148: { /* '148' */
            return "Balmart"
        }
        case 149: { /* '149' */
            return "GFS GmbH"
        }
        case 15: { /* '15' */
            return "GEWISS"
        }
        case 150: { /* '150' */
            return "Schenker Storen AG"
        }
        case 151: { /* '151' */
            return "Algodue Elettronica S.r.L."
        }
        case 152: { /* '152' */
            return "ABB France"
        }
        case 153: { /* '153' */
            return "maintronic"
        }
        case 154: { /* '154' */
            return "Vantage"
        }
        case 155: { /* '155' */
            return "Foresis"
        }
        case 156: { /* '156' */
            return "Research & Production Association SEM"
        }
        case 157: { /* '157' */
            return "Weinzierl Engineering GmbH"
        }
        case 158: { /* '158' */
            return "Möhlenhoff Wärmetechnik GmbH"
        }
        case 159: { /* '159' */
            return "PKC-GROUP Oyj"
        }
        case 16: { /* '16' */
            return "Albert Ackermann"
        }
        case 160: { /* '160' */
            return "B.E.G."
        }
        case 161: { /* '161' */
            return "Elsner Elektronik GmbH"
        }
        case 162: { /* '162' */
            return "Siemens Building Technologies (HK/China) Ltd."
        }
        case 163: { /* '163' */
            return "Eutrac"
        }
        case 164: { /* '164' */
            return "Gustav Hensel GmbH & Co. KG"
        }
        case 165: { /* '165' */
            return "GARO AB"
        }
        case 166: { /* '166' */
            return "Waldmann Lichttechnik"
        }
        case 167: { /* '167' */
            return "SCHÜCO"
        }
        case 168: { /* '168' */
            return "EMU"
        }
        case 169: { /* '169' */
            return "JNet Systems AG"
        }
        case 17: { /* '17' */
            return "Schupa GmbH"
        }
        case 170: { /* '170' */
            return "Total Solution GmbH"
        }
        case 171: { /* '171' */
            return "O.Y.L. Electronics"
        }
        case 172: { /* '172' */
            return "Galax System"
        }
        case 173: { /* '173' */
            return "Disch"
        }
        case 174: { /* '174' */
            return "Aucoteam"
        }
        case 175: { /* '175' */
            return "Luxmate Controls"
        }
        case 176: { /* '176' */
            return "Danfoss"
        }
        case 177: { /* '177' */
            return "AST GmbH"
        }
        case 178: { /* '178' */
            return "WILA Leuchten"
        }
        case 179: { /* '179' */
            return "b+b Automations- und Steuerungstechnik"
        }
        case 18: { /* '18' */
            return "ABB SCHWEIZ"
        }
        case 180: { /* '180' */
            return "Lingg & Janke"
        }
        case 181: { /* '181' */
            return "Sauter"
        }
        case 182: { /* '182' */
            return "SIMU"
        }
        case 183: { /* '183' */
            return "Theben HTS AG"
        }
        case 184: { /* '184' */
            return "Amann GmbH"
        }
        case 185: { /* '185' */
            return "BERG Energiekontrollsysteme GmbH"
        }
        case 186: { /* '186' */
            return "Hüppe Form Sonnenschutzsysteme GmbH"
        }
        case 187: { /* '187' */
            return "Oventrop KG"
        }
        case 188: { /* '188' */
            return "Griesser AG"
        }
        case 189: { /* '189' */
            return "IPAS GmbH"
        }
        case 19: { /* '19' */
            return "Feller"
        }
        case 190: { /* '190' */
            return "elero GmbH"
        }
        case 191: { /* '191' */
            return "Ardan Production and Industrial Controls Ltd."
        }
        case 192: { /* '192' */
            return "Metec Meßtechnik GmbH"
        }
        case 193: { /* '193' */
            return "ELKA-Elektronik GmbH"
        }
        case 194: { /* '194' */
            return "ELEKTROANLAGEN D. NAGEL"
        }
        case 195: { /* '195' */
            return "Tridonic Bauelemente GmbH"
        }
        case 196: { /* '196' */
            return "Stengler Gesellschaft"
        }
        case 197: { /* '197' */
            return "Schneider Electric (MG)"
        }
        case 198: { /* '198' */
            return "KNX Association"
        }
        case 199: { /* '199' */
            return "VIVO"
        }
        case 2: { /* '2' */
            return "ABB"
        }
        case 20: { /* '20' */
            return "Glamox AS"
        }
        case 200: { /* '200' */
            return "Hugo Müller GmbH & Co KG"
        }
        case 201: { /* '201' */
            return "Siemens HVAC"
        }
        case 202: { /* '202' */
            return "APT"
        }
        case 203: { /* '203' */
            return "HighDom"
        }
        case 204: { /* '204' */
            return "Top Services"
        }
        case 205: { /* '205' */
            return "ambiHome"
        }
        case 206: { /* '206' */
            return "DATEC electronic AG"
        }
        case 207: { /* '207' */
            return "ABUS Security-Center"
        }
        case 208: { /* '208' */
            return "Lite-Puter"
        }
        case 209: { /* '209' */
            return "Tantron Electronic"
        }
        case 21: { /* '21' */
            return "DEHN & SÖHNE"
        }
        case 210: { /* '210' */
            return "Interra"
        }
        case 211: { /* '211' */
            return "DKX Tech"
        }
        case 212: { /* '212' */
            return "Viatron"
        }
        case 213: { /* '213' */
            return "Nautibus"
        }
        case 214: { /* '214' */
            return "ON Semiconductor"
        }
        case 215: { /* '215' */
            return "Longchuang"
        }
        case 216: { /* '216' */
            return "Air-On AG"
        }
        case 217: { /* '217' */
            return "ib-company GmbH"
        }
        case 218: { /* '218' */
            return "Sation Factory"
        }
        case 219: { /* '219' */
            return "Agentilo GmbH"
        }
        case 22: { /* '22' */
            return "CRABTREE"
        }
        case 220: { /* '220' */
            return "Makel Elektrik"
        }
        case 221: { /* '221' */
            return "Helios Ventilatoren"
        }
        case 222: { /* '222' */
            return "Otto Solutions Pte Ltd"
        }
        case 223: { /* '223' */
            return "Airmaster"
        }
        case 224: { /* '224' */
            return "Vallox GmbH"
        }
        case 225: { /* '225' */
            return "Dalitek"
        }
        case 226: { /* '226' */
            return "ASIN"
        }
        case 227: { /* '227' */
            return "Bridges Intelligence Technology Inc."
        }
        case 228: { /* '228' */
            return "ARBONIA"
        }
        case 229: { /* '229' */
            return "KERMI"
        }
        case 23: { /* '23' */
            return "eVoKNX"
        }
        case 230: { /* '230' */
            return "PROLUX"
        }
        case 231: { /* '231' */
            return "ClicHome"
        }
        case 232: { /* '232' */
            return "COMMAX"
        }
        case 233: { /* '233' */
            return "EAE"
        }
        case 234: { /* '234' */
            return "Tense"
        }
        case 235: { /* '235' */
            return "Seyoung Electronics"
        }
        case 236: { /* '236' */
            return "Lifedomus"
        }
        case 237: { /* '237' */
            return "EUROtronic Technology GmbH"
        }
        case 238: { /* '238' */
            return "tci"
        }
        case 239: { /* '239' */
            return "Rishun Electronic"
        }
        case 24: { /* '24' */
            return "Paul Hochköpper"
        }
        case 240: { /* '240' */
            return "Zipato"
        }
        case 241: { /* '241' */
            return "cm-security GmbH & Co KG"
        }
        case 242: { /* '242' */
            return "Qing Cables"
        }
        case 243: { /* '243' */
            return "LABIO"
        }
        case 244: { /* '244' */
            return "Coster Tecnologie Elettroniche S.p.A."
        }
        case 245: { /* '245' */
            return "E.G.E"
        }
        case 246: { /* '246' */
            return "NETxAutomation"
        }
        case 247: { /* '247' */
            return "tecalor"
        }
        case 248: { /* '248' */
            return "Urmet Electronics (Huizhou) Ltd."
        }
        case 249: { /* '249' */
            return "Peiying Building Control"
        }
        case 25: { /* '25' */
            return "Altenburger Electronic"
        }
        case 250: { /* '250' */
            return "BPT S.p.A. a Socio Unico"
        }
        case 251: { /* '251' */
            return "Kanontec - KanonBUS"
        }
        case 252: { /* '252' */
            return "ISER Tech"
        }
        case 253: { /* '253' */
            return "Fineline"
        }
        case 254: { /* '254' */
            return "CP Electronics Ltd"
        }
        case 255: { /* '255' */
            return "Niko-Servodan A/S"
        }
        case 256: { /* '256' */
            return "Simon"
        }
        case 257: { /* '257' */
            return "GM modular pvt. Ltd."
        }
        case 258: { /* '258' */
            return "FU CHENG Intelligence"
        }
        case 259: { /* '259' */
            return "NexKon"
        }
        case 26: { /* '26' */
            return "Grässlin"
        }
        case 260: { /* '260' */
            return "FEEL s.r.l"
        }
        case 261: { /* '261' */
            return "Not Assigned"
        }
        case 262: { /* '262' */
            return "Shenzhen Fanhai Sanjiang Electronics Co., Ltd."
        }
        case 263: { /* '263' */
            return "Jiuzhou Greeble"
        }
        case 264: { /* '264' */
            return "Aumüller Aumatic GmbH"
        }
        case 265: { /* '265' */
            return "Etman Electric"
        }
        case 266: { /* '266' */
            return "Black Nova"
        }
        case 267: { /* '267' */
            return "ZidaTech AG"
        }
        case 268: { /* '268' */
            return "IDGS bvba"
        }
        case 269: { /* '269' */
            return "dakanimo"
        }
        case 27: { /* '27' */
            return "Simon"
        }
        case 270: { /* '270' */
            return "Trebor Automation AB"
        }
        case 271: { /* '271' */
            return "Satel sp. z o.o."
        }
        case 272: { /* '272' */
            return "Russound, Inc."
        }
        case 273: { /* '273' */
            return "Midea Heating & Ventilating Equipment CO LTD"
        }
        case 274: { /* '274' */
            return "Consorzio Terranuova"
        }
        case 275: { /* '275' */
            return "Wolf Heiztechnik GmbH"
        }
        case 276: { /* '276' */
            return "SONTEC"
        }
        case 277: { /* '277' */
            return "Belcom Cables Ltd."
        }
        case 278: { /* '278' */
            return "Guangzhou SeaWin Electrical Technologies Co., Ltd."
        }
        case 279: { /* '279' */
            return "Acrel"
        }
        case 28: { /* '28' */
            return "VIMAR"
        }
        case 280: { /* '280' */
            return "Franke Aquarotter GmbH"
        }
        case 281: { /* '281' */
            return "Orion Systems"
        }
        case 282: { /* '282' */
            return "Schrack Technik GmbH"
        }
        case 283: { /* '283' */
            return "INSPRID"
        }
        case 284: { /* '284' */
            return "Sunricher"
        }
        case 285: { /* '285' */
            return "Menred automation system(shanghai) Co.,Ltd."
        }
        case 286: { /* '286' */
            return "Aurex"
        }
        case 287: { /* '287' */
            return "Josef Barthelme GmbH & Co. KG"
        }
        case 288: { /* '288' */
            return "Architecture Numerique"
        }
        case 289: { /* '289' */
            return "UP GROUP"
        }
        case 29: { /* '29' */
            return "Moeller Gebäudeautomation KG"
        }
        case 290: { /* '290' */
            return "Teknos-Avinno"
        }
        case 291: { /* '291' */
            return "Ningbo Dooya Mechanic & Electronic Technology"
        }
        case 292: { /* '292' */
            return "Thermokon Sensortechnik GmbH"
        }
        case 293: { /* '293' */
            return "BELIMO Automation AG"
        }
        case 294: { /* '294' */
            return "Zehnder Group International AG"
        }
        case 295: { /* '295' */
            return "sks Kinkel Elektronik"
        }
        case 296: { /* '296' */
            return "ECE Wurmitzer GmbH"
        }
        case 297: { /* '297' */
            return "LARS"
        }
        case 298: { /* '298' */
            return "URC"
        }
        case 299: { /* '299' */
            return "LightControl"
        }
        case 3: { /* '3' */
            return "Albrecht Jung"
        }
        case 30: { /* '30' */
            return "Eltako"
        }
        case 300: { /* '300' */
            return "ShenZhen YM"
        }
        case 301: { /* '301' */
            return "MEAN WELL Enterprises Co. Ltd."
        }
        case 302: { /* '302' */
            return "OSix"
        }
        case 303: { /* '303' */
            return "AYPRO Technology"
        }
        case 304: { /* '304' */
            return "Hefei Ecolite Software"
        }
        case 305: { /* '305' */
            return "Enno"
        }
        case 306: { /* '306' */
            return "OHOSURE"
        }
        case 307: { /* '307' */
            return "Garefowl"
        }
        case 308: { /* '308' */
            return "GEZE"
        }
        case 309: { /* '309' */
            return "LG Electronics Inc."
        }
        case 31: { /* '31' */
            return "Bosch-Siemens Haushaltsgeräte"
        }
        case 310: { /* '310' */
            return "SMC interiors"
        }
        case 311: { /* '311' */
            return "Not Assigned"
        }
        case 312: { /* '312' */
            return "SCS Cable"
        }
        case 313: { /* '313' */
            return "Hoval"
        }
        case 314: { /* '314' */
            return "CANST"
        }
        case 315: { /* '315' */
            return "HangZhou Berlin"
        }
        case 316: { /* '316' */
            return "EVN-Lichttechnik"
        }
        case 317: { /* '317' */
            return "rutec"
        }
        case 318: { /* '318' */
            return "Finder"
        }
        case 319: { /* '319' */
            return "Fujitsu General Limited"
        }
        case 32: { /* '32' */
            return "RITTO GmbH&Co.KG"
        }
        case 320: { /* '320' */
            return "ZF Friedrichshafen AG"
        }
        case 321: { /* '321' */
            return "Crealed"
        }
        case 322: { /* '322' */
            return "Miles Magic Automation Private Limited"
        }
        case 323: { /* '323' */
            return "E+"
        }
        case 324: { /* '324' */
            return "Italcond"
        }
        case 325: { /* '325' */
            return "SATION"
        }
        case 326: { /* '326' */
            return "NewBest"
        }
        case 327: { /* '327' */
            return "GDS DIGITAL SYSTEMS"
        }
        case 328: { /* '328' */
            return "Iddero"
        }
        case 329: { /* '329' */
            return "MBNLED"
        }
        case 33: { /* '33' */
            return "Power Controls"
        }
        case 330: { /* '330' */
            return "VITRUM"
        }
        case 331: { /* '331' */
            return "ekey biometric systems GmbH"
        }
        case 332: { /* '332' */
            return "AMC"
        }
        case 333: { /* '333' */
            return "TRILUX GmbH & Co. KG"
        }
        case 334: { /* '334' */
            return "WExcedo"
        }
        case 335: { /* '335' */
            return "VEMER SPA"
        }
        case 336: { /* '336' */
            return "Alexander Bürkle GmbH & Co KG"
        }
        case 337: { /* '337' */
            return "Citron"
        }
        case 338: { /* '338' */
            return "Shenzhen HeGuang"
        }
        case 339: { /* '339' */
            return "Not Assigned"
        }
        case 34: { /* '34' */
            return "ZUMTOBEL"
        }
        case 340: { /* '340' */
            return "TRANE B.V.B.A"
        }
        case 341: { /* '341' */
            return "CAREL"
        }
        case 342: { /* '342' */
            return "Prolite Controls"
        }
        case 343: { /* '343' */
            return "BOSMER"
        }
        case 344: { /* '344' */
            return "EUCHIPS"
        }
        case 345: { /* '345' */
            return "connect (Thinka connect)"
        }
        case 346: { /* '346' */
            return "PEAKnx a DOGAWIST company"
        }
        case 347: { /* '347' */
            return "ACEMATIC"
        }
        case 348: { /* '348' */
            return "ELAUSYS"
        }
        case 349: { /* '349' */
            return "ITK Engineering AG"
        }
        case 35: { /* '35' */
            return "Phoenix Contact"
        }
        case 350: { /* '350' */
            return "INTEGRA METERING AG"
        }
        case 351: { /* '351' */
            return "FMS Hospitality Pte Ltd"
        }
        case 352: { /* '352' */
            return "Nuvo"
        }
        case 353: { /* '353' */
            return "u::Lux GmbH"
        }
        case 354: { /* '354' */
            return "Brumberg Leuchten"
        }
        case 355: { /* '355' */
            return "Lime"
        }
        case 356: { /* '356' */
            return "Great Empire International Group Co., Ltd."
        }
        case 357: { /* '357' */
            return "Kavoshpishro Asia"
        }
        case 358: { /* '358' */
            return "V2 SpA"
        }
        case 359: { /* '359' */
            return "Johnson Controls"
        }
        case 36: { /* '36' */
            return "WAGO Kontakttechnik"
        }
        case 360: { /* '360' */
            return "Arkud"
        }
        case 361: { /* '361' */
            return "Iridium Ltd."
        }
        case 362: { /* '362' */
            return "bsmart"
        }
        case 363: { /* '363' */
            return "BAB TECHNOLOGIE GmbH"
        }
        case 364: { /* '364' */
            return "NICE Spa"
        }
        case 365: { /* '365' */
            return "Redfish Group Pty Ltd"
        }
        case 366: { /* '366' */
            return "SABIANA spa"
        }
        case 367: { /* '367' */
            return "Ubee Interactive Europe"
        }
        case 368: { /* '368' */
            return "Rexel"
        }
        case 369: { /* '369' */
            return "Ges Teknik A.S."
        }
        case 37: { /* '37' */
            return "knXpresso"
        }
        case 370: { /* '370' */
            return "Ave S.p.A."
        }
        case 371: { /* '371' */
            return "Zhuhai Ltech Technology Co., Ltd."
        }
        case 372: { /* '372' */
            return "ARCOM"
        }
        case 373: { /* '373' */
            return "VIA Technologies, Inc."
        }
        case 374: { /* '374' */
            return "FEELSMART."
        }
        case 375: { /* '375' */
            return "SUPCON"
        }
        case 376: { /* '376' */
            return "MANIC"
        }
        case 377: { /* '377' */
            return "TDE GmbH"
        }
        case 378: { /* '378' */
            return "Nanjing Shufan Information technology Co.,Ltd."
        }
        case 379: { /* '379' */
            return "EWTech"
        }
        case 38: { /* '38' */
            return "Wieland Electric"
        }
        case 380: { /* '380' */
            return "Kluger Automation GmbH"
        }
        case 381: { /* '381' */
            return "JoongAng Control"
        }
        case 382: { /* '382' */
            return "GreenControls Technology Sdn. Bhd."
        }
        case 383: { /* '383' */
            return "IME S.p.a."
        }
        case 384: { /* '384' */
            return "SiChuan HaoDing"
        }
        case 385: { /* '385' */
            return "Mindjaga Ltd."
        }
        case 386: { /* '386' */
            return "RuiLi Smart Control"
        }
        case 387: { /* '387' */
            return "CODESYS GmbH"
        }
        case 388: { /* '388' */
            return "Moorgen Deutschland GmbH"
        }
        case 389: { /* '389' */
            return "CULLMANN TECH"
        }
        case 39: { /* '39' */
            return "Hermann Kleinhuis"
        }
        case 390: { /* '390' */
            return "Merck Window Technologies B.V."
        }
        case 391: { /* '391' */
            return "ABEGO"
        }
        case 392: { /* '392' */
            return "myGEKKO"
        }
        case 393: { /* '393' */
            return "Ergo3 Sarl"
        }
        case 394: { /* '394' */
            return "STmicroelectronics International N.V."
        }
        case 395: { /* '395' */
            return "cjc systems"
        }
        case 396: { /* '396' */
            return "Sudoku"
        }
        case 397: { /* '397' */
            return "AZ e-lite Pte Ltd"
        }
        case 398: { /* '398' */
            return "Arlight"
        }
        case 399: { /* '399' */
            return "Grünbeck Wasseraufbereitung GmbH"
        }
        case 4: { /* '4' */
            return "Bticino"
        }
        case 40: { /* '40' */
            return "Stiebel Eltron"
        }
        case 400: { /* '400' */
            return "Module Electronic"
        }
        case 401: { /* '401' */
            return "KOPLAT"
        }
        case 402: { /* '402' */
            return "Guangzhou Letour Life Technology Co., Ltd"
        }
        case 403: { /* '403' */
            return "ILEVIA"
        }
        case 404: { /* '404' */
            return "LN SYSTEMTEQ"
        }
        case 405: { /* '405' */
            return "Hisense SmartHome"
        }
        case 406: { /* '406' */
            return "Flink Automation System"
        }
        case 407: { /* '407' */
            return "xxter bv"
        }
        case 408: { /* '408' */
            return "lynxus technology"
        }
        case 409: { /* '409' */
            return "ROBOT S.A."
        }
        case 41: { /* '41' */
            return "Tehalit"
        }
        case 410: { /* '410' */
            return "Shenzhen Atte Smart Life Co.,Ltd."
        }
        case 411: { /* '411' */
            return "Noblesse"
        }
        case 412: { /* '412' */
            return "Advanced Devices"
        }
        case 413: { /* '413' */
            return "Atrina Building Automation Co. Ltd"
        }
        case 414: { /* '414' */
            return "Guangdong Daming Laffey electric Co., Ltd."
        }
        case 415: { /* '415' */
            return "Westerstrand Urfabrik AB"
        }
        case 416: { /* '416' */
            return "Control4 Corporate"
        }
        case 417: { /* '417' */
            return "Ontrol"
        }
        case 418: { /* '418' */
            return "Starnet"
        }
        case 419: { /* '419' */
            return "BETA CAVI"
        }
        case 42: { /* '42' */
            return "Theben AG"
        }
        case 420: { /* '420' */
            return "EaseMore"
        }
        case 421: { /* '421' */
            return "Vivaldi srl"
        }
        case 422: { /* '422' */
            return "Gree Electric Appliances,Inc. of Zhuhai"
        }
        case 423: { /* '423' */
            return "HWISCON"
        }
        case 424: { /* '424' */
            return "Shanghai ELECON Intelligent Technology Co., Ltd."
        }
        case 425: { /* '425' */
            return "Kampmann"
        }
        case 426: { /* '426' */
            return "Impolux GmbH / LEDIMAX"
        }
        case 427: { /* '427' */
            return "Evaux"
        }
        case 428: { /* '428' */
            return "Webro Cables & Connectors Limited"
        }
        case 429: { /* '429' */
            return "Shanghai E-tech Solution"
        }
        case 43: { /* '43' */
            return "Wilhelm Rutenbeck"
        }
        case 430: { /* '430' */
            return "Guangzhou HOKO Electric Co.,Ltd."
        }
        case 431: { /* '431' */
            return "LAMMIN HIGH TECH CO.,LTD"
        }
        case 432: { /* '432' */
            return "Shenzhen Merrytek Technology Co., Ltd"
        }
        case 433: { /* '433' */
            return "I-Luxus"
        }
        case 434: { /* '434' */
            return "Elmos Semiconductor AG"
        }
        case 435: { /* '435' */
            return "EmCom Technology Inc"
        }
        case 436: { /* '436' */
            return "project innovations GmbH"
        }
        case 437: { /* '437' */
            return "Itc"
        }
        case 438: { /* '438' */
            return "ABB LV Installation Materials Company Ltd, Beijing"
        }
        case 439: { /* '439' */
            return "Maico"
        }
        case 44: { /* '44' */
            return "Winkhaus"
        }
        case 440: { /* '440' */
            return "ELAN SRL"
        }
        case 441: { /* '441' */
            return "MinhHa Technology co.,Ltd"
        }
        case 442: { /* '442' */
            return "Zhejiang Tianjie Industrial CORP."
        }
        case 443: { /* '443' */
            return "iAutomation Pty Limited"
        }
        case 444: { /* '444' */
            return "Extron"
        }
        case 445: { /* '445' */
            return "Freedompro"
        }
        case 446: { /* '446' */
            return "1Home"
        }
        case 447: { /* '447' */
            return "EOS Saunatechnik GmbH"
        }
        case 448: { /* '448' */
            return "KUSATEK GmbH"
        }
        case 449: { /* '449' */
            return "EisBär Scada"
        }
        case 45: { /* '45' */
            return "Robert Bosch"
        }
        case 450: { /* '450' */
            return "AUTOMATISMI BENINCA S.P.A."
        }
        case 451: { /* '451' */
            return "Blendom"
        }
        case 452: { /* '452' */
            return "Madel Air Technical diffusion"
        }
        case 453: { /* '453' */
            return "NIKO"
        }
        case 454: { /* '454' */
            return "Bosch Rexroth AG"
        }
        case 455: { /* '455' */
            return "C&M Products"
        }
        case 456: { /* '456' */
            return "Hörmann KG Verkaufsgesellschaft"
        }
        case 457: { /* '457' */
            return "Shanghai Rajayasa co.,LTD"
        }
        case 458: { /* '458' */
            return "SUZUKI"
        }
        case 459: { /* '459' */
            return "Silent Gliss International Ltd."
        }
        case 46: { /* '46' */
            return "Somfy"
        }
        case 460: { /* '460' */
            return "BEE Controls (ADGSC Group)"
        }
        case 461: { /* '461' */
            return "xDTecGmbH"
        }
        case 462: { /* '462' */
            return "OSRAM"
        }
        case 463: { /* '463' */
            return "Lebenor"
        }
        case 464: { /* '464' */
            return "automaneng"
        }
        case 465: { /* '465' */
            return "Honeywell Automation Solution control(China)"
        }
        case 466: { /* '466' */
            return "Hangzhou binthen Intelligence Technology Co.,Ltd"
        }
        case 467: { /* '467' */
            return "ETA Heiztechnik"
        }
        case 468: { /* '468' */
            return "DIVUS GmbH"
        }
        case 469: { /* '469' */
            return "Nanjing Taijiesai Intelligent Technology Co. Ltd."
        }
        case 47: { /* '47' */
            return "Woertz"
        }
        case 470: { /* '470' */
            return "Lunatone"
        }
        case 471: { /* '471' */
            return "ZHEJIANG SCTECH BUILDING INTELLIGENT"
        }
        case 472: { /* '472' */
            return "Foshan Qite Technology Co., Ltd."
        }
        case 473: { /* '473' */
            return "NOKE"
        }
        case 474: { /* '474' */
            return "LANDCOM"
        }
        case 475: { /* '475' */
            return "Stork AS"
        }
        case 476: { /* '476' */
            return "Hangzhou Shendu Technology Co., Ltd."
        }
        case 477: { /* '477' */
            return "CoolAutomation"
        }
        case 478: { /* '478' */
            return "Aprstern"
        }
        case 479: { /* '479' */
            return "sonnen"
        }
        case 48: { /* '48' */
            return "Viessmann Werke"
        }
        case 480: { /* '480' */
            return "DNAKE"
        }
        case 481: { /* '481' */
            return "Neuberger Gebäudeautomation GmbH"
        }
        case 482: { /* '482' */
            return "Stiliger"
        }
        case 483: { /* '483' */
            return "Berghof Automation GmbH"
        }
        case 484: { /* '484' */
            return "Total Automation and controls GmbH"
        }
        case 485: { /* '485' */
            return "dovit"
        }
        case 486: { /* '486' */
            return "Instalighting GmbH"
        }
        case 487: { /* '487' */
            return "UNI-TEC"
        }
        case 488: { /* '488' */
            return "CasaTunes"
        }
        case 489: { /* '489' */
            return "EMT"
        }
        case 49: { /* '49' */
            return "IMI Hydronic Engineering"
        }
        case 490: { /* '490' */
            return "Senfficient"
        }
        case 491: { /* '491' */
            return "Aurolite electrical panyu guangzhou limited"
        }
        case 492: { /* '492' */
            return "ABB Xiamen Smart Technology Co., Ltd."
        }
        case 493: { /* '493' */
            return "Samson Electric Wire"
        }
        case 494: { /* '494' */
            return "T-Touching"
        }
        case 495: { /* '495' */
            return "Core Smart Home"
        }
        case 496: { /* '496' */
            return "GreenConnect Solutions SA"
        }
        case 497: { /* '497' */
            return "ELETTRONICA CONDUTTORI"
        }
        case 498: { /* '498' */
            return "MKFC"
        }
        case 499: { /* '499' */
            return "Automation+"
        }
        case 5: { /* '5' */
            return "Berker"
        }
        case 50: { /* '50' */
            return "Joh. Vaillant"
        }
        case 500: { /* '500' */
            return "blue and red"
        }
        case 501: { /* '501' */
            return "frogblue"
        }
        case 502: { /* '502' */
            return "SAVESOR"
        }
        case 503: { /* '503' */
            return "App Tech"
        }
        case 504: { /* '504' */
            return "sensortec AG"
        }
        case 505: { /* '505' */
            return "nysa technology & solutions"
        }
        case 506: { /* '506' */
            return "FARADITE"
        }
        case 507: { /* '507' */
            return "Optimus"
        }
        case 508: { /* '508' */
            return "KTS s.r.l."
        }
        case 509: { /* '509' */
            return "Ramcro SPA"
        }
        case 51: { /* '51' */
            return "AMP Deutschland"
        }
        case 510: { /* '510' */
            return "Wuhan WiseCreate Universe Technology Co., Ltd"
        }
        case 511: { /* '511' */
            return "BEMI Smart Home Ltd"
        }
        case 512: { /* '512' */
            return "Ardomus"
        }
        case 513: { /* '513' */
            return "ChangXing"
        }
        case 514: { /* '514' */
            return "E-Controls"
        }
        case 515: { /* '515' */
            return "AIB Technology"
        }
        case 516: { /* '516' */
            return "NVC"
        }
        case 517: { /* '517' */
            return "Kbox"
        }
        case 518: { /* '518' */
            return "CNS"
        }
        case 519: { /* '519' */
            return "Tyba"
        }
        case 52: { /* '52' */
            return "Bosch Thermotechnik GmbH"
        }
        case 520: { /* '520' */
            return "Atrel"
        }
        case 521: { /* '521' */
            return "Simon Electric (China) Co., LTD"
        }
        case 522: { /* '522' */
            return "Kordz Group"
        }
        case 523: { /* '523' */
            return "ND Electric"
        }
        case 524: { /* '524' */
            return "Controlium"
        }
        case 525: { /* '525' */
            return "FAMO GmbH & Co. KG"
        }
        case 526: { /* '526' */
            return "CDN Smart"
        }
        case 527: { /* '527' */
            return "Heston"
        }
        case 528: { /* '528' */
            return "ESLA CONEXIONES S.L."
        }
        case 529: { /* '529' */
            return "Weishaupt"
        }
        case 53: { /* '53' */
            return "SEF - ECOTEC"
        }
        case 530: { /* '530' */
            return "ASTRUM TECHNOLOGY"
        }
        case 531: { /* '531' */
            return "WUERTH ELEKTRONIK STELVIO KONTEK S.p.A."
        }
        case 532: { /* '532' */
            return "NANOTECO corporation"
        }
        case 533: { /* '533' */
            return "Nietian"
        }
        case 534: { /* '534' */
            return "Sumsir"
        }
        case 535: { /* '535' */
            return "ORBIS TECNOLOGIA ELECTRICA SA"
        }
        case 536: { /* '536' */
            return "ABB - reserved"
        }
        case 537: { /* '537' */
            return "Busch-Jaeger Elektro - reserved"
        }
        case 54: { /* '54' */
            return "DORMA GmbH + Co. KG"
        }
        case 55: { /* '55' */
            return "WindowMaster A/S"
        }
        case 56: { /* '56' */
            return "Walther Werke"
        }
        case 57: { /* '57' */
            return "ORAS"
        }
        case 58: { /* '58' */
            return "Dätwyler"
        }
        case 59: { /* '59' */
            return "Electrak"
        }
        case 6: { /* '6' */
            return "Busch-Jaeger Elektro"
        }
        case 60: { /* '60' */
            return "Techem"
        }
        case 61: { /* '61' */
            return "Schneider Electric Industries SAS"
        }
        case 62: { /* '62' */
            return "WHD Wilhelm Huber + Söhne"
        }
        case 63: { /* '63' */
            return "Bischoff Elektronik"
        }
        case 64: { /* '64' */
            return "JEPAZ"
        }
        case 65: { /* '65' */
            return "RTS Automation"
        }
        case 66: { /* '66' */
            return "EIBMARKT GmbH"
        }
        case 67: { /* '67' */
            return "WAREMA Renkhoff SE"
        }
        case 68: { /* '68' */
            return "Eelectron"
        }
        case 69: { /* '69' */
            return "Belden Wire & Cable B.V."
        }
        case 7: { /* '7' */
            return "GIRA Giersiepen"
        }
        case 70: { /* '70' */
            return "Becker-Antriebe GmbH"
        }
        case 71: { /* '71' */
            return "J.Stehle+Söhne GmbH"
        }
        case 72: { /* '72' */
            return "AGFEO"
        }
        case 73: { /* '73' */
            return "Zennio"
        }
        case 74: { /* '74' */
            return "TAPKO Technologies"
        }
        case 75: { /* '75' */
            return "HDL"
        }
        case 76: { /* '76' */
            return "Uponor"
        }
        case 77: { /* '77' */
            return "se Lightmanagement AG"
        }
        case 78: { /* '78' */
            return "Arcus-eds"
        }
        case 79: { /* '79' */
            return "Intesis"
        }
        case 8: { /* '8' */
            return "Hager Electro"
        }
        case 80: { /* '80' */
            return "Herholdt Controls srl"
        }
        case 81: { /* '81' */
            return "Niko-Zublin"
        }
        case 82: { /* '82' */
            return "Durable Technologies"
        }
        case 83: { /* '83' */
            return "Innoteam"
        }
        case 84: { /* '84' */
            return "ise GmbH"
        }
        case 85: { /* '85' */
            return "TEAM FOR TRONICS"
        }
        case 86: { /* '86' */
            return "CIAT"
        }
        case 87: { /* '87' */
            return "Remeha BV"
        }
        case 88: { /* '88' */
            return "ESYLUX"
        }
        case 89: { /* '89' */
            return "BASALTE"
        }
        case 9: { /* '9' */
            return "Insta GmbH"
        }
        case 90: { /* '90' */
            return "Vestamatic"
        }
        case 91: { /* '91' */
            return "MDT technologies"
        }
        case 92: { /* '92' */
            return "Warendorfer Küchen GmbH"
        }
        case 93: { /* '93' */
            return "Video-Star"
        }
        case 94: { /* '94' */
            return "Sitek"
        }
        case 95: { /* '95' */
            return "CONTROLtronic"
        }
        case 96: { /* '96' */
            return "function Technology"
        }
        case 97: { /* '97' */
            return "AMX"
        }
        case 98: { /* '98' */
            return "ELDAT"
        }
        case 99: { /* '99' */
            return "Panasonic"
        }
        default: {
            return ""
        }
    }
}
func KnxManufacturerByValue(value uint16) KnxManufacturer {
    switch value {
        case 0:
            return KnxManufacturer_M_UNKNOWN
        case 1:
            return KnxManufacturer_M_SIEMENS
        case 10:
            return KnxManufacturer_M_LEGRAND_APPAREILLAGE_ELECTRIQUE
        case 100:
            return KnxManufacturer_M_PULSE_TECHNOLOGIES
        case 101:
            return KnxManufacturer_M_CRESTRON
        case 102:
            return KnxManufacturer_M_STEINEL_PROFESSIONAL
        case 103:
            return KnxManufacturer_M_BILTON_LED_LIGHTING
        case 104:
            return KnxManufacturer_M_DENRO_AG
        case 105:
            return KnxManufacturer_M_GEPRO
        case 106:
            return KnxManufacturer_M_PREUSSEN_AUTOMATION
        case 107:
            return KnxManufacturer_M_ZOPPAS_INDUSTRIES
        case 108:
            return KnxManufacturer_M_MACTECH
        case 109:
            return KnxManufacturer_M_TECHNO_TREND
        case 11:
            return KnxManufacturer_M_MERTEN
        case 110:
            return KnxManufacturer_M_FS_CABLES
        case 111:
            return KnxManufacturer_M_DELTA_DORE
        case 112:
            return KnxManufacturer_M_EISSOUND
        case 113:
            return KnxManufacturer_M_CISCO
        case 114:
            return KnxManufacturer_M_DINUY
        case 115:
            return KnxManufacturer_M_IKNIX
        case 116:
            return KnxManufacturer_M_RADEMACHER_GERAETE_ELEKTRONIK_GMBH
        case 117:
            return KnxManufacturer_M_EGI_ELECTROACUSTICA_GENERAL_IBERICA
        case 118:
            return KnxManufacturer_M_BES___INGENIUM
        case 119:
            return KnxManufacturer_M_ELABNET
        case 12:
            return KnxManufacturer_M_ABB_SPA_SACE_DIVISION
        case 120:
            return KnxManufacturer_M_BLUMOTIX
        case 121:
            return KnxManufacturer_M_HUNTER_DOUGLAS
        case 122:
            return KnxManufacturer_M_APRICUM
        case 123:
            return KnxManufacturer_M_TIANSU_AUTOMATION
        case 124:
            return KnxManufacturer_M_BUBENDORFF
        case 125:
            return KnxManufacturer_M_MBS_GMBH
        case 126:
            return KnxManufacturer_M_ENERTEX_BAYERN_GMBH
        case 127:
            return KnxManufacturer_M_BMS
        case 128:
            return KnxManufacturer_M_SINAPSI
        case 129:
            return KnxManufacturer_M_EMBEDDED_SYSTEMS_SIA
        case 13:
            return KnxManufacturer_M_SIEDLE_AND_SOEHNE
        case 130:
            return KnxManufacturer_M_KNX1
        case 131:
            return KnxManufacturer_M_TOKKA
        case 132:
            return KnxManufacturer_M_NANOSENSE
        case 133:
            return KnxManufacturer_M_PEAR_AUTOMATION_GMBH
        case 134:
            return KnxManufacturer_M_DGA
        case 135:
            return KnxManufacturer_M_LUTRON
        case 136:
            return KnxManufacturer_M_AIRZONE___ALTRA
        case 137:
            return KnxManufacturer_M_LITHOSS_DESIGN_SWITCHES
        case 138:
            return KnxManufacturer_M_THREEATEL
        case 139:
            return KnxManufacturer_M_PHILIPS_CONTROLS
        case 14:
            return KnxManufacturer_M_EBERLE
        case 140:
            return KnxManufacturer_M_VELUX_AS
        case 141:
            return KnxManufacturer_M_LOYTEC
        case 142:
            return KnxManufacturer_M_EKINEX_S_P_A_
        case 143:
            return KnxManufacturer_M_SIRLAN_TECHNOLOGIES
        case 144:
            return KnxManufacturer_M_PROKNX_SAS
        case 145:
            return KnxManufacturer_M_IT_GMBH
        case 146:
            return KnxManufacturer_M_RENSON
        case 147:
            return KnxManufacturer_M_HEP_GROUP
        case 148:
            return KnxManufacturer_M_BALMART
        case 149:
            return KnxManufacturer_M_GFS_GMBH
        case 15:
            return KnxManufacturer_M_GEWISS
        case 150:
            return KnxManufacturer_M_SCHENKER_STOREN_AG
        case 151:
            return KnxManufacturer_M_ALGODUE_ELETTRONICA_S_R_L_
        case 152:
            return KnxManufacturer_M_ABB_FRANCE
        case 153:
            return KnxManufacturer_M_MAINTRONIC
        case 154:
            return KnxManufacturer_M_VANTAGE
        case 155:
            return KnxManufacturer_M_FORESIS
        case 156:
            return KnxManufacturer_M_RESEARCH_AND_PRODUCTION_ASSOCIATION_SEM
        case 157:
            return KnxManufacturer_M_WEINZIERL_ENGINEERING_GMBH
        case 158:
            return KnxManufacturer_M_MOEHLENHOFF_WAERMETECHNIK_GMBH
        case 159:
            return KnxManufacturer_M_PKC_GROUP_OYJ
        case 16:
            return KnxManufacturer_M_ALBERT_ACKERMANN
        case 160:
            return KnxManufacturer_M_B_E_G_
        case 161:
            return KnxManufacturer_M_ELSNER_ELEKTRONIK_GMBH
        case 162:
            return KnxManufacturer_M_SIEMENS_BUILDING_TECHNOLOGIES_HKCHINA_LTD_
        case 163:
            return KnxManufacturer_M_EUTRAC
        case 164:
            return KnxManufacturer_M_GUSTAV_HENSEL_GMBH_AND_CO__KG
        case 165:
            return KnxManufacturer_M_GARO_AB
        case 166:
            return KnxManufacturer_M_WALDMANN_LICHTTECHNIK
        case 167:
            return KnxManufacturer_M_SCHUECO
        case 168:
            return KnxManufacturer_M_EMU
        case 169:
            return KnxManufacturer_M_JNET_SYSTEMS_AG
        case 17:
            return KnxManufacturer_M_SCHUPA_GMBH
        case 170:
            return KnxManufacturer_M_TOTAL_SOLUTION_GMBH
        case 171:
            return KnxManufacturer_M_O_Y_L__ELECTRONICS
        case 172:
            return KnxManufacturer_M_GALAX_SYSTEM
        case 173:
            return KnxManufacturer_M_DISCH
        case 174:
            return KnxManufacturer_M_AUCOTEAM
        case 175:
            return KnxManufacturer_M_LUXMATE_CONTROLS
        case 176:
            return KnxManufacturer_M_DANFOSS
        case 177:
            return KnxManufacturer_M_AST_GMBH
        case 178:
            return KnxManufacturer_M_WILA_LEUCHTEN
        case 179:
            return KnxManufacturer_M_BPlusB_AUTOMATIONS__UND_STEUERUNGSTECHNIK
        case 18:
            return KnxManufacturer_M_ABB_SCHWEIZ
        case 180:
            return KnxManufacturer_M_LINGG_AND_JANKE
        case 181:
            return KnxManufacturer_M_SAUTER
        case 182:
            return KnxManufacturer_M_SIMU
        case 183:
            return KnxManufacturer_M_THEBEN_HTS_AG
        case 184:
            return KnxManufacturer_M_AMANN_GMBH
        case 185:
            return KnxManufacturer_M_BERG_ENERGIEKONTROLLSYSTEME_GMBH
        case 186:
            return KnxManufacturer_M_HUEPPE_FORM_SONNENSCHUTZSYSTEME_GMBH
        case 187:
            return KnxManufacturer_M_OVENTROP_KG
        case 188:
            return KnxManufacturer_M_GRIESSER_AG
        case 189:
            return KnxManufacturer_M_IPAS_GMBH
        case 19:
            return KnxManufacturer_M_FELLER
        case 190:
            return KnxManufacturer_M_ELERO_GMBH
        case 191:
            return KnxManufacturer_M_ARDAN_PRODUCTION_AND_INDUSTRIAL_CONTROLS_LTD_
        case 192:
            return KnxManufacturer_M_METEC_MESSTECHNIK_GMBH
        case 193:
            return KnxManufacturer_M_ELKA_ELEKTRONIK_GMBH
        case 194:
            return KnxManufacturer_M_ELEKTROANLAGEN_D__NAGEL
        case 195:
            return KnxManufacturer_M_TRIDONIC_BAUELEMENTE_GMBH
        case 196:
            return KnxManufacturer_M_STENGLER_GESELLSCHAFT
        case 197:
            return KnxManufacturer_M_SCHNEIDER_ELECTRIC_MG
        case 198:
            return KnxManufacturer_M_KNX_ASSOCIATION
        case 199:
            return KnxManufacturer_M_VIVO
        case 2:
            return KnxManufacturer_M_ABB
        case 20:
            return KnxManufacturer_M_GLAMOX_AS
        case 200:
            return KnxManufacturer_M_HUGO_MUELLER_GMBH_AND_CO_KG
        case 201:
            return KnxManufacturer_M_SIEMENS_HVAC
        case 202:
            return KnxManufacturer_M_APT
        case 203:
            return KnxManufacturer_M_HIGHDOM
        case 204:
            return KnxManufacturer_M_TOP_SERVICES
        case 205:
            return KnxManufacturer_M_AMBIHOME
        case 206:
            return KnxManufacturer_M_DATEC_ELECTRONIC_AG
        case 207:
            return KnxManufacturer_M_ABUS_SECURITY_CENTER
        case 208:
            return KnxManufacturer_M_LITE_PUTER
        case 209:
            return KnxManufacturer_M_TANTRON_ELECTRONIC
        case 21:
            return KnxManufacturer_M_DEHN_AND_SOEHNE
        case 210:
            return KnxManufacturer_M_INTERRA
        case 211:
            return KnxManufacturer_M_DKX_TECH
        case 212:
            return KnxManufacturer_M_VIATRON
        case 213:
            return KnxManufacturer_M_NAUTIBUS
        case 214:
            return KnxManufacturer_M_ON_SEMICONDUCTOR
        case 215:
            return KnxManufacturer_M_LONGCHUANG
        case 216:
            return KnxManufacturer_M_AIR_ON_AG
        case 217:
            return KnxManufacturer_M_IB_COMPANY_GMBH
        case 218:
            return KnxManufacturer_M_SATION_FACTORY
        case 219:
            return KnxManufacturer_M_AGENTILO_GMBH
        case 22:
            return KnxManufacturer_M_CRABTREE
        case 220:
            return KnxManufacturer_M_MAKEL_ELEKTRIK
        case 221:
            return KnxManufacturer_M_HELIOS_VENTILATOREN
        case 222:
            return KnxManufacturer_M_OTTO_SOLUTIONS_PTE_LTD
        case 223:
            return KnxManufacturer_M_AIRMASTER
        case 224:
            return KnxManufacturer_M_VALLOX_GMBH
        case 225:
            return KnxManufacturer_M_DALITEK
        case 226:
            return KnxManufacturer_M_ASIN
        case 227:
            return KnxManufacturer_M_BRIDGES_INTELLIGENCE_TECHNOLOGY_INC_
        case 228:
            return KnxManufacturer_M_ARBONIA
        case 229:
            return KnxManufacturer_M_KERMI
        case 23:
            return KnxManufacturer_M_EVOKNX
        case 230:
            return KnxManufacturer_M_PROLUX
        case 231:
            return KnxManufacturer_M_CLICHOME
        case 232:
            return KnxManufacturer_M_COMMAX
        case 233:
            return KnxManufacturer_M_EAE
        case 234:
            return KnxManufacturer_M_TENSE
        case 235:
            return KnxManufacturer_M_SEYOUNG_ELECTRONICS
        case 236:
            return KnxManufacturer_M_LIFEDOMUS
        case 237:
            return KnxManufacturer_M_EUROTRONIC_TECHNOLOGY_GMBH
        case 238:
            return KnxManufacturer_M_TCI
        case 239:
            return KnxManufacturer_M_RISHUN_ELECTRONIC
        case 24:
            return KnxManufacturer_M_PAUL_HOCHKOEPPER
        case 240:
            return KnxManufacturer_M_ZIPATO
        case 241:
            return KnxManufacturer_M_CM_SECURITY_GMBH_AND_CO_KG
        case 242:
            return KnxManufacturer_M_QING_CABLES
        case 243:
            return KnxManufacturer_M_LABIO
        case 244:
            return KnxManufacturer_M_COSTER_TECNOLOGIE_ELETTRONICHE_S_P_A_
        case 245:
            return KnxManufacturer_M_E_G_E
        case 246:
            return KnxManufacturer_M_NETXAUTOMATION
        case 247:
            return KnxManufacturer_M_TECALOR
        case 248:
            return KnxManufacturer_M_URMET_ELECTRONICS_HUIZHOU_LTD_
        case 249:
            return KnxManufacturer_M_PEIYING_BUILDING_CONTROL
        case 25:
            return KnxManufacturer_M_ALTENBURGER_ELECTRONIC
        case 250:
            return KnxManufacturer_M_BPT_S_P_A__A_SOCIO_UNICO
        case 251:
            return KnxManufacturer_M_KANONTEC___KANONBUS
        case 252:
            return KnxManufacturer_M_ISER_TECH
        case 253:
            return KnxManufacturer_M_FINELINE
        case 254:
            return KnxManufacturer_M_CP_ELECTRONICS_LTD
        case 255:
            return KnxManufacturer_M_NIKO_SERVODAN_AS
        case 256:
            return KnxManufacturer_M_SIMON_309
        case 257:
            return KnxManufacturer_M_GM_MODULAR_PVT__LTD_
        case 258:
            return KnxManufacturer_M_FU_CHENG_INTELLIGENCE
        case 259:
            return KnxManufacturer_M_NEXKON
        case 26:
            return KnxManufacturer_M_GRAESSLIN
        case 260:
            return KnxManufacturer_M_FEEL_S_R_L
        case 261:
            return KnxManufacturer_M_NOT_ASSIGNED_314
        case 262:
            return KnxManufacturer_M_SHENZHEN_FANHAI_SANJIANG_ELECTRONICS_CO___LTD_
        case 263:
            return KnxManufacturer_M_JIUZHOU_GREEBLE
        case 264:
            return KnxManufacturer_M_AUMUELLER_AUMATIC_GMBH
        case 265:
            return KnxManufacturer_M_ETMAN_ELECTRIC
        case 266:
            return KnxManufacturer_M_BLACK_NOVA
        case 267:
            return KnxManufacturer_M_ZIDATECH_AG
        case 268:
            return KnxManufacturer_M_IDGS_BVBA
        case 269:
            return KnxManufacturer_M_DAKANIMO
        case 27:
            return KnxManufacturer_M_SIMON_42
        case 270:
            return KnxManufacturer_M_TREBOR_AUTOMATION_AB
        case 271:
            return KnxManufacturer_M_SATEL_SP__Z_O_O_
        case 272:
            return KnxManufacturer_M_RUSSOUND__INC_
        case 273:
            return KnxManufacturer_M_MIDEA_HEATING_AND_VENTILATING_EQUIPMENT_CO_LTD
        case 274:
            return KnxManufacturer_M_CONSORZIO_TERRANUOVA
        case 275:
            return KnxManufacturer_M_WOLF_HEIZTECHNIK_GMBH
        case 276:
            return KnxManufacturer_M_SONTEC
        case 277:
            return KnxManufacturer_M_BELCOM_CABLES_LTD_
        case 278:
            return KnxManufacturer_M_GUANGZHOU_SEAWIN_ELECTRICAL_TECHNOLOGIES_CO___LTD_
        case 279:
            return KnxManufacturer_M_ACREL
        case 28:
            return KnxManufacturer_M_VIMAR
        case 280:
            return KnxManufacturer_M_FRANKE_AQUAROTTER_GMBH
        case 281:
            return KnxManufacturer_M_ORION_SYSTEMS
        case 282:
            return KnxManufacturer_M_SCHRACK_TECHNIK_GMBH
        case 283:
            return KnxManufacturer_M_INSPRID
        case 284:
            return KnxManufacturer_M_SUNRICHER
        case 285:
            return KnxManufacturer_M_MENRED_AUTOMATION_SYSTEMSHANGHAI_CO__LTD_
        case 286:
            return KnxManufacturer_M_AUREX
        case 287:
            return KnxManufacturer_M_JOSEF_BARTHELME_GMBH_AND_CO__KG
        case 288:
            return KnxManufacturer_M_ARCHITECTURE_NUMERIQUE
        case 289:
            return KnxManufacturer_M_UP_GROUP
        case 29:
            return KnxManufacturer_M_MOELLER_GEBAEUDEAUTOMATION_KG
        case 290:
            return KnxManufacturer_M_TEKNOS_AVINNO
        case 291:
            return KnxManufacturer_M_NINGBO_DOOYA_MECHANIC_AND_ELECTRONIC_TECHNOLOGY
        case 292:
            return KnxManufacturer_M_THERMOKON_SENSORTECHNIK_GMBH
        case 293:
            return KnxManufacturer_M_BELIMO_AUTOMATION_AG
        case 294:
            return KnxManufacturer_M_ZEHNDER_GROUP_INTERNATIONAL_AG
        case 295:
            return KnxManufacturer_M_SKS_KINKEL_ELEKTRONIK
        case 296:
            return KnxManufacturer_M_ECE_WURMITZER_GMBH
        case 297:
            return KnxManufacturer_M_LARS
        case 298:
            return KnxManufacturer_M_URC
        case 299:
            return KnxManufacturer_M_LIGHTCONTROL
        case 3:
            return KnxManufacturer_M_ALBRECHT_JUNG
        case 30:
            return KnxManufacturer_M_ELTAKO
        case 300:
            return KnxManufacturer_M_SHENZHEN_YM
        case 301:
            return KnxManufacturer_M_MEAN_WELL_ENTERPRISES_CO__LTD_
        case 302:
            return KnxManufacturer_M_OSIX
        case 303:
            return KnxManufacturer_M_AYPRO_TECHNOLOGY
        case 304:
            return KnxManufacturer_M_HEFEI_ECOLITE_SOFTWARE
        case 305:
            return KnxManufacturer_M_ENNO
        case 306:
            return KnxManufacturer_M_OHOSURE
        case 307:
            return KnxManufacturer_M_GAREFOWL
        case 308:
            return KnxManufacturer_M_GEZE
        case 309:
            return KnxManufacturer_M_LG_ELECTRONICS_INC_
        case 31:
            return KnxManufacturer_M_BOSCH_SIEMENS_HAUSHALTSGERAETE
        case 310:
            return KnxManufacturer_M_SMC_INTERIORS
        case 311:
            return KnxManufacturer_M_NOT_ASSIGNED_364
        case 312:
            return KnxManufacturer_M_SCS_CABLE
        case 313:
            return KnxManufacturer_M_HOVAL
        case 314:
            return KnxManufacturer_M_CANST
        case 315:
            return KnxManufacturer_M_HANGZHOU_BERLIN
        case 316:
            return KnxManufacturer_M_EVN_LICHTTECHNIK
        case 317:
            return KnxManufacturer_M_RUTEC
        case 318:
            return KnxManufacturer_M_FINDER
        case 319:
            return KnxManufacturer_M_FUJITSU_GENERAL_LIMITED
        case 32:
            return KnxManufacturer_M_RITTO_GMBHANDCO_KG
        case 320:
            return KnxManufacturer_M_ZF_FRIEDRICHSHAFEN_AG
        case 321:
            return KnxManufacturer_M_CREALED
        case 322:
            return KnxManufacturer_M_MILES_MAGIC_AUTOMATION_PRIVATE_LIMITED
        case 323:
            return KnxManufacturer_M_EPlus
        case 324:
            return KnxManufacturer_M_ITALCOND
        case 325:
            return KnxManufacturer_M_SATION
        case 326:
            return KnxManufacturer_M_NEWBEST
        case 327:
            return KnxManufacturer_M_GDS_DIGITAL_SYSTEMS
        case 328:
            return KnxManufacturer_M_IDDERO
        case 329:
            return KnxManufacturer_M_MBNLED
        case 33:
            return KnxManufacturer_M_POWER_CONTROLS
        case 330:
            return KnxManufacturer_M_VITRUM
        case 331:
            return KnxManufacturer_M_EKEY_BIOMETRIC_SYSTEMS_GMBH
        case 332:
            return KnxManufacturer_M_AMC
        case 333:
            return KnxManufacturer_M_TRILUX_GMBH_AND_CO__KG
        case 334:
            return KnxManufacturer_M_WEXCEDO
        case 335:
            return KnxManufacturer_M_VEMER_SPA
        case 336:
            return KnxManufacturer_M_ALEXANDER_BUERKLE_GMBH_AND_CO_KG
        case 337:
            return KnxManufacturer_M_CITRON
        case 338:
            return KnxManufacturer_M_SHENZHEN_HEGUANG
        case 339:
            return KnxManufacturer_M_NOT_ASSIGNED_392
        case 34:
            return KnxManufacturer_M_ZUMTOBEL
        case 340:
            return KnxManufacturer_M_TRANE_B_V_B_A
        case 341:
            return KnxManufacturer_M_CAREL
        case 342:
            return KnxManufacturer_M_PROLITE_CONTROLS
        case 343:
            return KnxManufacturer_M_BOSMER
        case 344:
            return KnxManufacturer_M_EUCHIPS
        case 345:
            return KnxManufacturer_M_CONNECT_THINKA_CONNECT
        case 346:
            return KnxManufacturer_M_PEAKNX_A_DOGAWIST_COMPANY
        case 347:
            return KnxManufacturer_M_ACEMATIC
        case 348:
            return KnxManufacturer_M_ELAUSYS
        case 349:
            return KnxManufacturer_M_ITK_ENGINEERING_AG
        case 35:
            return KnxManufacturer_M_PHOENIX_CONTACT
        case 350:
            return KnxManufacturer_M_INTEGRA_METERING_AG
        case 351:
            return KnxManufacturer_M_FMS_HOSPITALITY_PTE_LTD
        case 352:
            return KnxManufacturer_M_NUVO
        case 353:
            return KnxManufacturer_M_U__LUX_GMBH
        case 354:
            return KnxManufacturer_M_BRUMBERG_LEUCHTEN
        case 355:
            return KnxManufacturer_M_LIME
        case 356:
            return KnxManufacturer_M_GREAT_EMPIRE_INTERNATIONAL_GROUP_CO___LTD_
        case 357:
            return KnxManufacturer_M_KAVOSHPISHRO_ASIA
        case 358:
            return KnxManufacturer_M_V2_SPA
        case 359:
            return KnxManufacturer_M_JOHNSON_CONTROLS
        case 36:
            return KnxManufacturer_M_WAGO_KONTAKTTECHNIK
        case 360:
            return KnxManufacturer_M_ARKUD
        case 361:
            return KnxManufacturer_M_IRIDIUM_LTD_
        case 362:
            return KnxManufacturer_M_BSMART
        case 363:
            return KnxManufacturer_M_BAB_TECHNOLOGIE_GMBH
        case 364:
            return KnxManufacturer_M_NICE_SPA
        case 365:
            return KnxManufacturer_M_REDFISH_GROUP_PTY_LTD
        case 366:
            return KnxManufacturer_M_SABIANA_SPA
        case 367:
            return KnxManufacturer_M_UBEE_INTERACTIVE_EUROPE
        case 368:
            return KnxManufacturer_M_REXEL
        case 369:
            return KnxManufacturer_M_GES_TEKNIK_A_S_
        case 37:
            return KnxManufacturer_M_KNXPRESSO
        case 370:
            return KnxManufacturer_M_AVE_S_P_A_
        case 371:
            return KnxManufacturer_M_ZHUHAI_LTECH_TECHNOLOGY_CO___LTD_
        case 372:
            return KnxManufacturer_M_ARCOM
        case 373:
            return KnxManufacturer_M_VIA_TECHNOLOGIES__INC_
        case 374:
            return KnxManufacturer_M_FEELSMART_
        case 375:
            return KnxManufacturer_M_SUPCON
        case 376:
            return KnxManufacturer_M_MANIC
        case 377:
            return KnxManufacturer_M_TDE_GMBH
        case 378:
            return KnxManufacturer_M_NANJING_SHUFAN_INFORMATION_TECHNOLOGY_CO__LTD_
        case 379:
            return KnxManufacturer_M_EWTECH
        case 38:
            return KnxManufacturer_M_WIELAND_ELECTRIC
        case 380:
            return KnxManufacturer_M_KLUGER_AUTOMATION_GMBH
        case 381:
            return KnxManufacturer_M_JOONGANG_CONTROL
        case 382:
            return KnxManufacturer_M_GREENCONTROLS_TECHNOLOGY_SDN__BHD_
        case 383:
            return KnxManufacturer_M_IME_S_P_A_
        case 384:
            return KnxManufacturer_M_SICHUAN_HAODING
        case 385:
            return KnxManufacturer_M_MINDJAGA_LTD_
        case 386:
            return KnxManufacturer_M_RUILI_SMART_CONTROL
        case 387:
            return KnxManufacturer_M_CODESYS_GMBH
        case 388:
            return KnxManufacturer_M_MOORGEN_DEUTSCHLAND_GMBH
        case 389:
            return KnxManufacturer_M_CULLMANN_TECH
        case 39:
            return KnxManufacturer_M_HERMANN_KLEINHUIS
        case 390:
            return KnxManufacturer_M_MERCK_WINDOW_TECHNOLOGIES_B_V_
        case 391:
            return KnxManufacturer_M_ABEGO
        case 392:
            return KnxManufacturer_M_MYGEKKO
        case 393:
            return KnxManufacturer_M_ERGO3_SARL
        case 394:
            return KnxManufacturer_M_STMICROELECTRONICS_INTERNATIONAL_N_V_
        case 395:
            return KnxManufacturer_M_CJC_SYSTEMS
        case 396:
            return KnxManufacturer_M_SUDOKU
        case 397:
            return KnxManufacturer_M_AZ_E_LITE_PTE_LTD
        case 398:
            return KnxManufacturer_M_ARLIGHT
        case 399:
            return KnxManufacturer_M_GRUENBECK_WASSERAUFBEREITUNG_GMBH
        case 4:
            return KnxManufacturer_M_BTICINO
        case 40:
            return KnxManufacturer_M_STIEBEL_ELTRON
        case 400:
            return KnxManufacturer_M_MODULE_ELECTRONIC
        case 401:
            return KnxManufacturer_M_KOPLAT
        case 402:
            return KnxManufacturer_M_GUANGZHOU_LETOUR_LIFE_TECHNOLOGY_CO___LTD
        case 403:
            return KnxManufacturer_M_ILEVIA
        case 404:
            return KnxManufacturer_M_LN_SYSTEMTEQ
        case 405:
            return KnxManufacturer_M_HISENSE_SMARTHOME
        case 406:
            return KnxManufacturer_M_FLINK_AUTOMATION_SYSTEM
        case 407:
            return KnxManufacturer_M_XXTER_BV
        case 408:
            return KnxManufacturer_M_LYNXUS_TECHNOLOGY
        case 409:
            return KnxManufacturer_M_ROBOT_S_A_
        case 41:
            return KnxManufacturer_M_TEHALIT
        case 410:
            return KnxManufacturer_M_SHENZHEN_ATTE_SMART_LIFE_CO__LTD_
        case 411:
            return KnxManufacturer_M_NOBLESSE
        case 412:
            return KnxManufacturer_M_ADVANCED_DEVICES
        case 413:
            return KnxManufacturer_M_ATRINA_BUILDING_AUTOMATION_CO__LTD
        case 414:
            return KnxManufacturer_M_GUANGDONG_DAMING_LAFFEY_ELECTRIC_CO___LTD_
        case 415:
            return KnxManufacturer_M_WESTERSTRAND_URFABRIK_AB
        case 416:
            return KnxManufacturer_M_CONTROL4_CORPORATE
        case 417:
            return KnxManufacturer_M_ONTROL
        case 418:
            return KnxManufacturer_M_STARNET
        case 419:
            return KnxManufacturer_M_BETA_CAVI
        case 42:
            return KnxManufacturer_M_THEBEN_AG
        case 420:
            return KnxManufacturer_M_EASEMORE
        case 421:
            return KnxManufacturer_M_VIVALDI_SRL
        case 422:
            return KnxManufacturer_M_GREE_ELECTRIC_APPLIANCES_INC__OF_ZHUHAI
        case 423:
            return KnxManufacturer_M_HWISCON
        case 424:
            return KnxManufacturer_M_SHANGHAI_ELECON_INTELLIGENT_TECHNOLOGY_CO___LTD_
        case 425:
            return KnxManufacturer_M_KAMPMANN
        case 426:
            return KnxManufacturer_M_IMPOLUX_GMBH_LEDIMAX
        case 427:
            return KnxManufacturer_M_EVAUX
        case 428:
            return KnxManufacturer_M_WEBRO_CABLES_AND_CONNECTORS_LIMITED
        case 429:
            return KnxManufacturer_M_SHANGHAI_E_TECH_SOLUTION
        case 43:
            return KnxManufacturer_M_WILHELM_RUTENBECK
        case 430:
            return KnxManufacturer_M_GUANGZHOU_HOKO_ELECTRIC_CO__LTD_
        case 431:
            return KnxManufacturer_M_LAMMIN_HIGH_TECH_CO__LTD
        case 432:
            return KnxManufacturer_M_SHENZHEN_MERRYTEK_TECHNOLOGY_CO___LTD
        case 433:
            return KnxManufacturer_M_I_LUXUS
        case 434:
            return KnxManufacturer_M_ELMOS_SEMICONDUCTOR_AG
        case 435:
            return KnxManufacturer_M_EMCOM_TECHNOLOGY_INC
        case 436:
            return KnxManufacturer_M_PROJECT_INNOVATIONS_GMBH
        case 437:
            return KnxManufacturer_M_ITC
        case 438:
            return KnxManufacturer_M_ABB_LV_INSTALLATION_MATERIALS_COMPANY_LTD__BEIJING
        case 439:
            return KnxManufacturer_M_MAICO
        case 44:
            return KnxManufacturer_M_WINKHAUS
        case 440:
            return KnxManufacturer_M_ELAN_SRL
        case 441:
            return KnxManufacturer_M_MINHHA_TECHNOLOGY_CO__LTD
        case 442:
            return KnxManufacturer_M_ZHEJIANG_TIANJIE_INDUSTRIAL_CORP_
        case 443:
            return KnxManufacturer_M_IAUTOMATION_PTY_LIMITED
        case 444:
            return KnxManufacturer_M_EXTRON
        case 445:
            return KnxManufacturer_M_FREEDOMPRO
        case 446:
            return KnxManufacturer_M_ONEHOME
        case 447:
            return KnxManufacturer_M_EOS_SAUNATECHNIK_GMBH
        case 448:
            return KnxManufacturer_M_KUSATEK_GMBH
        case 449:
            return KnxManufacturer_M_EISBAER_SCADA
        case 45:
            return KnxManufacturer_M_ROBERT_BOSCH
        case 450:
            return KnxManufacturer_M_AUTOMATISMI_BENINCA_S_P_A_
        case 451:
            return KnxManufacturer_M_BLENDOM
        case 452:
            return KnxManufacturer_M_MADEL_AIR_TECHNICAL_DIFFUSION
        case 453:
            return KnxManufacturer_M_NIKO
        case 454:
            return KnxManufacturer_M_BOSCH_REXROTH_AG
        case 455:
            return KnxManufacturer_M_CANDM_PRODUCTS
        case 456:
            return KnxManufacturer_M_HOERMANN_KG_VERKAUFSGESELLSCHAFT
        case 457:
            return KnxManufacturer_M_SHANGHAI_RAJAYASA_CO__LTD
        case 458:
            return KnxManufacturer_M_SUZUKI
        case 459:
            return KnxManufacturer_M_SILENT_GLISS_INTERNATIONAL_LTD_
        case 46:
            return KnxManufacturer_M_SOMFY
        case 460:
            return KnxManufacturer_M_BEE_CONTROLS_ADGSC_GROUP
        case 461:
            return KnxManufacturer_M_XDTECGMBH
        case 462:
            return KnxManufacturer_M_OSRAM
        case 463:
            return KnxManufacturer_M_LEBENOR
        case 464:
            return KnxManufacturer_M_AUTOMANENG
        case 465:
            return KnxManufacturer_M_HONEYWELL_AUTOMATION_SOLUTION_CONTROLCHINA
        case 466:
            return KnxManufacturer_M_HANGZHOU_BINTHEN_INTELLIGENCE_TECHNOLOGY_CO__LTD
        case 467:
            return KnxManufacturer_M_ETA_HEIZTECHNIK
        case 468:
            return KnxManufacturer_M_DIVUS_GMBH
        case 469:
            return KnxManufacturer_M_NANJING_TAIJIESAI_INTELLIGENT_TECHNOLOGY_CO__LTD_
        case 47:
            return KnxManufacturer_M_WOERTZ
        case 470:
            return KnxManufacturer_M_LUNATONE
        case 471:
            return KnxManufacturer_M_ZHEJIANG_SCTECH_BUILDING_INTELLIGENT
        case 472:
            return KnxManufacturer_M_FOSHAN_QITE_TECHNOLOGY_CO___LTD_
        case 473:
            return KnxManufacturer_M_NOKE
        case 474:
            return KnxManufacturer_M_LANDCOM
        case 475:
            return KnxManufacturer_M_STORK_AS
        case 476:
            return KnxManufacturer_M_HANGZHOU_SHENDU_TECHNOLOGY_CO___LTD_
        case 477:
            return KnxManufacturer_M_COOLAUTOMATION
        case 478:
            return KnxManufacturer_M_APRSTERN
        case 479:
            return KnxManufacturer_M_SONNEN
        case 48:
            return KnxManufacturer_M_VIESSMANN_WERKE
        case 480:
            return KnxManufacturer_M_DNAKE
        case 481:
            return KnxManufacturer_M_NEUBERGER_GEBAEUDEAUTOMATION_GMBH
        case 482:
            return KnxManufacturer_M_STILIGER
        case 483:
            return KnxManufacturer_M_BERGHOF_AUTOMATION_GMBH
        case 484:
            return KnxManufacturer_M_TOTAL_AUTOMATION_AND_CONTROLS_GMBH
        case 485:
            return KnxManufacturer_M_DOVIT
        case 486:
            return KnxManufacturer_M_INSTALIGHTING_GMBH
        case 487:
            return KnxManufacturer_M_UNI_TEC
        case 488:
            return KnxManufacturer_M_CASATUNES
        case 489:
            return KnxManufacturer_M_EMT
        case 49:
            return KnxManufacturer_M_IMI_HYDRONIC_ENGINEERING
        case 490:
            return KnxManufacturer_M_SENFFICIENT
        case 491:
            return KnxManufacturer_M_AUROLITE_ELECTRICAL_PANYU_GUANGZHOU_LIMITED
        case 492:
            return KnxManufacturer_M_ABB_XIAMEN_SMART_TECHNOLOGY_CO___LTD_
        case 493:
            return KnxManufacturer_M_SAMSON_ELECTRIC_WIRE
        case 494:
            return KnxManufacturer_M_T_TOUCHING
        case 495:
            return KnxManufacturer_M_CORE_SMART_HOME
        case 496:
            return KnxManufacturer_M_GREENCONNECT_SOLUTIONS_SA
        case 497:
            return KnxManufacturer_M_ELETTRONICA_CONDUTTORI
        case 498:
            return KnxManufacturer_M_MKFC
        case 499:
            return KnxManufacturer_M_AUTOMATIONPlus
        case 5:
            return KnxManufacturer_M_BERKER
        case 50:
            return KnxManufacturer_M_JOH__VAILLANT
        case 500:
            return KnxManufacturer_M_BLUE_AND_RED
        case 501:
            return KnxManufacturer_M_FROGBLUE
        case 502:
            return KnxManufacturer_M_SAVESOR
        case 503:
            return KnxManufacturer_M_APP_TECH
        case 504:
            return KnxManufacturer_M_SENSORTEC_AG
        case 505:
            return KnxManufacturer_M_NYSA_TECHNOLOGY_AND_SOLUTIONS
        case 506:
            return KnxManufacturer_M_FARADITE
        case 507:
            return KnxManufacturer_M_OPTIMUS
        case 508:
            return KnxManufacturer_M_KTS_S_R_L_
        case 509:
            return KnxManufacturer_M_RAMCRO_SPA
        case 51:
            return KnxManufacturer_M_AMP_DEUTSCHLAND
        case 510:
            return KnxManufacturer_M_WUHAN_WISECREATE_UNIVERSE_TECHNOLOGY_CO___LTD
        case 511:
            return KnxManufacturer_M_BEMI_SMART_HOME_LTD
        case 512:
            return KnxManufacturer_M_ARDOMUS
        case 513:
            return KnxManufacturer_M_CHANGXING
        case 514:
            return KnxManufacturer_M_E_CONTROLS
        case 515:
            return KnxManufacturer_M_AIB_TECHNOLOGY
        case 516:
            return KnxManufacturer_M_NVC
        case 517:
            return KnxManufacturer_M_KBOX
        case 518:
            return KnxManufacturer_M_CNS
        case 519:
            return KnxManufacturer_M_TYBA
        case 52:
            return KnxManufacturer_M_BOSCH_THERMOTECHNIK_GMBH
        case 520:
            return KnxManufacturer_M_ATREL
        case 521:
            return KnxManufacturer_M_SIMON_ELECTRIC_CHINA_CO___LTD
        case 522:
            return KnxManufacturer_M_KORDZ_GROUP
        case 523:
            return KnxManufacturer_M_ND_ELECTRIC
        case 524:
            return KnxManufacturer_M_CONTROLIUM
        case 525:
            return KnxManufacturer_M_FAMO_GMBH_AND_CO__KG
        case 526:
            return KnxManufacturer_M_CDN_SMART
        case 527:
            return KnxManufacturer_M_HESTON
        case 528:
            return KnxManufacturer_M_ESLA_CONEXIONES_S_L_
        case 529:
            return KnxManufacturer_M_WEISHAUPT
        case 53:
            return KnxManufacturer_M_SEF___ECOTEC
        case 530:
            return KnxManufacturer_M_ASTRUM_TECHNOLOGY
        case 531:
            return KnxManufacturer_M_WUERTH_ELEKTRONIK_STELVIO_KONTEK_S_P_A_
        case 532:
            return KnxManufacturer_M_NANOTECO_CORPORATION
        case 533:
            return KnxManufacturer_M_NIETIAN
        case 534:
            return KnxManufacturer_M_SUMSIR
        case 535:
            return KnxManufacturer_M_ORBIS_TECNOLOGIA_ELECTRICA_SA
        case 536:
            return KnxManufacturer_M_ABB___RESERVED
        case 537:
            return KnxManufacturer_M_BUSCH_JAEGER_ELEKTRO___RESERVED
        case 54:
            return KnxManufacturer_M_DORMA_GMBH_Plus_CO__KG
        case 55:
            return KnxManufacturer_M_WINDOWMASTER_AS
        case 56:
            return KnxManufacturer_M_WALTHER_WERKE
        case 57:
            return KnxManufacturer_M_ORAS
        case 58:
            return KnxManufacturer_M_DAETWYLER
        case 59:
            return KnxManufacturer_M_ELECTRAK
        case 6:
            return KnxManufacturer_M_BUSCH_JAEGER_ELEKTRO
        case 60:
            return KnxManufacturer_M_TECHEM
        case 61:
            return KnxManufacturer_M_SCHNEIDER_ELECTRIC_INDUSTRIES_SAS
        case 62:
            return KnxManufacturer_M_WHD_WILHELM_HUBER_Plus_SOEHNE
        case 63:
            return KnxManufacturer_M_BISCHOFF_ELEKTRONIK
        case 64:
            return KnxManufacturer_M_JEPAZ
        case 65:
            return KnxManufacturer_M_RTS_AUTOMATION
        case 66:
            return KnxManufacturer_M_EIBMARKT_GMBH
        case 67:
            return KnxManufacturer_M_WAREMA_RENKHOFF_SE
        case 68:
            return KnxManufacturer_M_EELECTRON
        case 69:
            return KnxManufacturer_M_BELDEN_WIRE_AND_CABLE_B_V_
        case 7:
            return KnxManufacturer_M_GIRA_GIERSIEPEN
        case 70:
            return KnxManufacturer_M_BECKER_ANTRIEBE_GMBH
        case 71:
            return KnxManufacturer_M_J_STEHLEPlusSOEHNE_GMBH
        case 72:
            return KnxManufacturer_M_AGFEO
        case 73:
            return KnxManufacturer_M_ZENNIO
        case 74:
            return KnxManufacturer_M_TAPKO_TECHNOLOGIES
        case 75:
            return KnxManufacturer_M_HDL
        case 76:
            return KnxManufacturer_M_UPONOR
        case 77:
            return KnxManufacturer_M_SE_LIGHTMANAGEMENT_AG
        case 78:
            return KnxManufacturer_M_ARCUS_EDS
        case 79:
            return KnxManufacturer_M_INTESIS
        case 8:
            return KnxManufacturer_M_HAGER_ELECTRO
        case 80:
            return KnxManufacturer_M_HERHOLDT_CONTROLS_SRL
        case 81:
            return KnxManufacturer_M_NIKO_ZUBLIN
        case 82:
            return KnxManufacturer_M_DURABLE_TECHNOLOGIES
        case 83:
            return KnxManufacturer_M_INNOTEAM
        case 84:
            return KnxManufacturer_M_ISE_GMBH
        case 85:
            return KnxManufacturer_M_TEAM_FOR_TRONICS
        case 86:
            return KnxManufacturer_M_CIAT
        case 87:
            return KnxManufacturer_M_REMEHA_BV
        case 88:
            return KnxManufacturer_M_ESYLUX
        case 89:
            return KnxManufacturer_M_BASALTE
        case 9:
            return KnxManufacturer_M_INSTA_GMBH
        case 90:
            return KnxManufacturer_M_VESTAMATIC
        case 91:
            return KnxManufacturer_M_MDT_TECHNOLOGIES
        case 92:
            return KnxManufacturer_M_WARENDORFER_KUECHEN_GMBH
        case 93:
            return KnxManufacturer_M_VIDEO_STAR
        case 94:
            return KnxManufacturer_M_SITEK
        case 95:
            return KnxManufacturer_M_CONTROLTRONIC
        case 96:
            return KnxManufacturer_M_FUNCTION_TECHNOLOGY
        case 97:
            return KnxManufacturer_M_AMX
        case 98:
            return KnxManufacturer_M_ELDAT
        case 99:
            return KnxManufacturer_M_PANASONIC
    }
    return 0
}

func KnxManufacturerByName(value string) KnxManufacturer {
    switch value {
    case "M_UNKNOWN":
        return KnxManufacturer_M_UNKNOWN
    case "M_SIEMENS":
        return KnxManufacturer_M_SIEMENS
    case "M_LEGRAND_APPAREILLAGE_ELECTRIQUE":
        return KnxManufacturer_M_LEGRAND_APPAREILLAGE_ELECTRIQUE
    case "M_PULSE_TECHNOLOGIES":
        return KnxManufacturer_M_PULSE_TECHNOLOGIES
    case "M_CRESTRON":
        return KnxManufacturer_M_CRESTRON
    case "M_STEINEL_PROFESSIONAL":
        return KnxManufacturer_M_STEINEL_PROFESSIONAL
    case "M_BILTON_LED_LIGHTING":
        return KnxManufacturer_M_BILTON_LED_LIGHTING
    case "M_DENRO_AG":
        return KnxManufacturer_M_DENRO_AG
    case "M_GEPRO":
        return KnxManufacturer_M_GEPRO
    case "M_PREUSSEN_AUTOMATION":
        return KnxManufacturer_M_PREUSSEN_AUTOMATION
    case "M_ZOPPAS_INDUSTRIES":
        return KnxManufacturer_M_ZOPPAS_INDUSTRIES
    case "M_MACTECH":
        return KnxManufacturer_M_MACTECH
    case "M_TECHNO_TREND":
        return KnxManufacturer_M_TECHNO_TREND
    case "M_MERTEN":
        return KnxManufacturer_M_MERTEN
    case "M_FS_CABLES":
        return KnxManufacturer_M_FS_CABLES
    case "M_DELTA_DORE":
        return KnxManufacturer_M_DELTA_DORE
    case "M_EISSOUND":
        return KnxManufacturer_M_EISSOUND
    case "M_CISCO":
        return KnxManufacturer_M_CISCO
    case "M_DINUY":
        return KnxManufacturer_M_DINUY
    case "M_IKNIX":
        return KnxManufacturer_M_IKNIX
    case "M_RADEMACHER_GERAETE_ELEKTRONIK_GMBH":
        return KnxManufacturer_M_RADEMACHER_GERAETE_ELEKTRONIK_GMBH
    case "M_EGI_ELECTROACUSTICA_GENERAL_IBERICA":
        return KnxManufacturer_M_EGI_ELECTROACUSTICA_GENERAL_IBERICA
    case "M_BES___INGENIUM":
        return KnxManufacturer_M_BES___INGENIUM
    case "M_ELABNET":
        return KnxManufacturer_M_ELABNET
    case "M_ABB_SPA_SACE_DIVISION":
        return KnxManufacturer_M_ABB_SPA_SACE_DIVISION
    case "M_BLUMOTIX":
        return KnxManufacturer_M_BLUMOTIX
    case "M_HUNTER_DOUGLAS":
        return KnxManufacturer_M_HUNTER_DOUGLAS
    case "M_APRICUM":
        return KnxManufacturer_M_APRICUM
    case "M_TIANSU_AUTOMATION":
        return KnxManufacturer_M_TIANSU_AUTOMATION
    case "M_BUBENDORFF":
        return KnxManufacturer_M_BUBENDORFF
    case "M_MBS_GMBH":
        return KnxManufacturer_M_MBS_GMBH
    case "M_ENERTEX_BAYERN_GMBH":
        return KnxManufacturer_M_ENERTEX_BAYERN_GMBH
    case "M_BMS":
        return KnxManufacturer_M_BMS
    case "M_SINAPSI":
        return KnxManufacturer_M_SINAPSI
    case "M_EMBEDDED_SYSTEMS_SIA":
        return KnxManufacturer_M_EMBEDDED_SYSTEMS_SIA
    case "M_SIEDLE_AND_SOEHNE":
        return KnxManufacturer_M_SIEDLE_AND_SOEHNE
    case "M_KNX1":
        return KnxManufacturer_M_KNX1
    case "M_TOKKA":
        return KnxManufacturer_M_TOKKA
    case "M_NANOSENSE":
        return KnxManufacturer_M_NANOSENSE
    case "M_PEAR_AUTOMATION_GMBH":
        return KnxManufacturer_M_PEAR_AUTOMATION_GMBH
    case "M_DGA":
        return KnxManufacturer_M_DGA
    case "M_LUTRON":
        return KnxManufacturer_M_LUTRON
    case "M_AIRZONE___ALTRA":
        return KnxManufacturer_M_AIRZONE___ALTRA
    case "M_LITHOSS_DESIGN_SWITCHES":
        return KnxManufacturer_M_LITHOSS_DESIGN_SWITCHES
    case "M_THREEATEL":
        return KnxManufacturer_M_THREEATEL
    case "M_PHILIPS_CONTROLS":
        return KnxManufacturer_M_PHILIPS_CONTROLS
    case "M_EBERLE":
        return KnxManufacturer_M_EBERLE
    case "M_VELUX_AS":
        return KnxManufacturer_M_VELUX_AS
    case "M_LOYTEC":
        return KnxManufacturer_M_LOYTEC
    case "M_EKINEX_S_P_A_":
        return KnxManufacturer_M_EKINEX_S_P_A_
    case "M_SIRLAN_TECHNOLOGIES":
        return KnxManufacturer_M_SIRLAN_TECHNOLOGIES
    case "M_PROKNX_SAS":
        return KnxManufacturer_M_PROKNX_SAS
    case "M_IT_GMBH":
        return KnxManufacturer_M_IT_GMBH
    case "M_RENSON":
        return KnxManufacturer_M_RENSON
    case "M_HEP_GROUP":
        return KnxManufacturer_M_HEP_GROUP
    case "M_BALMART":
        return KnxManufacturer_M_BALMART
    case "M_GFS_GMBH":
        return KnxManufacturer_M_GFS_GMBH
    case "M_GEWISS":
        return KnxManufacturer_M_GEWISS
    case "M_SCHENKER_STOREN_AG":
        return KnxManufacturer_M_SCHENKER_STOREN_AG
    case "M_ALGODUE_ELETTRONICA_S_R_L_":
        return KnxManufacturer_M_ALGODUE_ELETTRONICA_S_R_L_
    case "M_ABB_FRANCE":
        return KnxManufacturer_M_ABB_FRANCE
    case "M_MAINTRONIC":
        return KnxManufacturer_M_MAINTRONIC
    case "M_VANTAGE":
        return KnxManufacturer_M_VANTAGE
    case "M_FORESIS":
        return KnxManufacturer_M_FORESIS
    case "M_RESEARCH_AND_PRODUCTION_ASSOCIATION_SEM":
        return KnxManufacturer_M_RESEARCH_AND_PRODUCTION_ASSOCIATION_SEM
    case "M_WEINZIERL_ENGINEERING_GMBH":
        return KnxManufacturer_M_WEINZIERL_ENGINEERING_GMBH
    case "M_MOEHLENHOFF_WAERMETECHNIK_GMBH":
        return KnxManufacturer_M_MOEHLENHOFF_WAERMETECHNIK_GMBH
    case "M_PKC_GROUP_OYJ":
        return KnxManufacturer_M_PKC_GROUP_OYJ
    case "M_ALBERT_ACKERMANN":
        return KnxManufacturer_M_ALBERT_ACKERMANN
    case "M_B_E_G_":
        return KnxManufacturer_M_B_E_G_
    case "M_ELSNER_ELEKTRONIK_GMBH":
        return KnxManufacturer_M_ELSNER_ELEKTRONIK_GMBH
    case "M_SIEMENS_BUILDING_TECHNOLOGIES_HKCHINA_LTD_":
        return KnxManufacturer_M_SIEMENS_BUILDING_TECHNOLOGIES_HKCHINA_LTD_
    case "M_EUTRAC":
        return KnxManufacturer_M_EUTRAC
    case "M_GUSTAV_HENSEL_GMBH_AND_CO__KG":
        return KnxManufacturer_M_GUSTAV_HENSEL_GMBH_AND_CO__KG
    case "M_GARO_AB":
        return KnxManufacturer_M_GARO_AB
    case "M_WALDMANN_LICHTTECHNIK":
        return KnxManufacturer_M_WALDMANN_LICHTTECHNIK
    case "M_SCHUECO":
        return KnxManufacturer_M_SCHUECO
    case "M_EMU":
        return KnxManufacturer_M_EMU
    case "M_JNET_SYSTEMS_AG":
        return KnxManufacturer_M_JNET_SYSTEMS_AG
    case "M_SCHUPA_GMBH":
        return KnxManufacturer_M_SCHUPA_GMBH
    case "M_TOTAL_SOLUTION_GMBH":
        return KnxManufacturer_M_TOTAL_SOLUTION_GMBH
    case "M_O_Y_L__ELECTRONICS":
        return KnxManufacturer_M_O_Y_L__ELECTRONICS
    case "M_GALAX_SYSTEM":
        return KnxManufacturer_M_GALAX_SYSTEM
    case "M_DISCH":
        return KnxManufacturer_M_DISCH
    case "M_AUCOTEAM":
        return KnxManufacturer_M_AUCOTEAM
    case "M_LUXMATE_CONTROLS":
        return KnxManufacturer_M_LUXMATE_CONTROLS
    case "M_DANFOSS":
        return KnxManufacturer_M_DANFOSS
    case "M_AST_GMBH":
        return KnxManufacturer_M_AST_GMBH
    case "M_WILA_LEUCHTEN":
        return KnxManufacturer_M_WILA_LEUCHTEN
    case "M_BPlusB_AUTOMATIONS__UND_STEUERUNGSTECHNIK":
        return KnxManufacturer_M_BPlusB_AUTOMATIONS__UND_STEUERUNGSTECHNIK
    case "M_ABB_SCHWEIZ":
        return KnxManufacturer_M_ABB_SCHWEIZ
    case "M_LINGG_AND_JANKE":
        return KnxManufacturer_M_LINGG_AND_JANKE
    case "M_SAUTER":
        return KnxManufacturer_M_SAUTER
    case "M_SIMU":
        return KnxManufacturer_M_SIMU
    case "M_THEBEN_HTS_AG":
        return KnxManufacturer_M_THEBEN_HTS_AG
    case "M_AMANN_GMBH":
        return KnxManufacturer_M_AMANN_GMBH
    case "M_BERG_ENERGIEKONTROLLSYSTEME_GMBH":
        return KnxManufacturer_M_BERG_ENERGIEKONTROLLSYSTEME_GMBH
    case "M_HUEPPE_FORM_SONNENSCHUTZSYSTEME_GMBH":
        return KnxManufacturer_M_HUEPPE_FORM_SONNENSCHUTZSYSTEME_GMBH
    case "M_OVENTROP_KG":
        return KnxManufacturer_M_OVENTROP_KG
    case "M_GRIESSER_AG":
        return KnxManufacturer_M_GRIESSER_AG
    case "M_IPAS_GMBH":
        return KnxManufacturer_M_IPAS_GMBH
    case "M_FELLER":
        return KnxManufacturer_M_FELLER
    case "M_ELERO_GMBH":
        return KnxManufacturer_M_ELERO_GMBH
    case "M_ARDAN_PRODUCTION_AND_INDUSTRIAL_CONTROLS_LTD_":
        return KnxManufacturer_M_ARDAN_PRODUCTION_AND_INDUSTRIAL_CONTROLS_LTD_
    case "M_METEC_MESSTECHNIK_GMBH":
        return KnxManufacturer_M_METEC_MESSTECHNIK_GMBH
    case "M_ELKA_ELEKTRONIK_GMBH":
        return KnxManufacturer_M_ELKA_ELEKTRONIK_GMBH
    case "M_ELEKTROANLAGEN_D__NAGEL":
        return KnxManufacturer_M_ELEKTROANLAGEN_D__NAGEL
    case "M_TRIDONIC_BAUELEMENTE_GMBH":
        return KnxManufacturer_M_TRIDONIC_BAUELEMENTE_GMBH
    case "M_STENGLER_GESELLSCHAFT":
        return KnxManufacturer_M_STENGLER_GESELLSCHAFT
    case "M_SCHNEIDER_ELECTRIC_MG":
        return KnxManufacturer_M_SCHNEIDER_ELECTRIC_MG
    case "M_KNX_ASSOCIATION":
        return KnxManufacturer_M_KNX_ASSOCIATION
    case "M_VIVO":
        return KnxManufacturer_M_VIVO
    case "M_ABB":
        return KnxManufacturer_M_ABB
    case "M_GLAMOX_AS":
        return KnxManufacturer_M_GLAMOX_AS
    case "M_HUGO_MUELLER_GMBH_AND_CO_KG":
        return KnxManufacturer_M_HUGO_MUELLER_GMBH_AND_CO_KG
    case "M_SIEMENS_HVAC":
        return KnxManufacturer_M_SIEMENS_HVAC
    case "M_APT":
        return KnxManufacturer_M_APT
    case "M_HIGHDOM":
        return KnxManufacturer_M_HIGHDOM
    case "M_TOP_SERVICES":
        return KnxManufacturer_M_TOP_SERVICES
    case "M_AMBIHOME":
        return KnxManufacturer_M_AMBIHOME
    case "M_DATEC_ELECTRONIC_AG":
        return KnxManufacturer_M_DATEC_ELECTRONIC_AG
    case "M_ABUS_SECURITY_CENTER":
        return KnxManufacturer_M_ABUS_SECURITY_CENTER
    case "M_LITE_PUTER":
        return KnxManufacturer_M_LITE_PUTER
    case "M_TANTRON_ELECTRONIC":
        return KnxManufacturer_M_TANTRON_ELECTRONIC
    case "M_DEHN_AND_SOEHNE":
        return KnxManufacturer_M_DEHN_AND_SOEHNE
    case "M_INTERRA":
        return KnxManufacturer_M_INTERRA
    case "M_DKX_TECH":
        return KnxManufacturer_M_DKX_TECH
    case "M_VIATRON":
        return KnxManufacturer_M_VIATRON
    case "M_NAUTIBUS":
        return KnxManufacturer_M_NAUTIBUS
    case "M_ON_SEMICONDUCTOR":
        return KnxManufacturer_M_ON_SEMICONDUCTOR
    case "M_LONGCHUANG":
        return KnxManufacturer_M_LONGCHUANG
    case "M_AIR_ON_AG":
        return KnxManufacturer_M_AIR_ON_AG
    case "M_IB_COMPANY_GMBH":
        return KnxManufacturer_M_IB_COMPANY_GMBH
    case "M_SATION_FACTORY":
        return KnxManufacturer_M_SATION_FACTORY
    case "M_AGENTILO_GMBH":
        return KnxManufacturer_M_AGENTILO_GMBH
    case "M_CRABTREE":
        return KnxManufacturer_M_CRABTREE
    case "M_MAKEL_ELEKTRIK":
        return KnxManufacturer_M_MAKEL_ELEKTRIK
    case "M_HELIOS_VENTILATOREN":
        return KnxManufacturer_M_HELIOS_VENTILATOREN
    case "M_OTTO_SOLUTIONS_PTE_LTD":
        return KnxManufacturer_M_OTTO_SOLUTIONS_PTE_LTD
    case "M_AIRMASTER":
        return KnxManufacturer_M_AIRMASTER
    case "M_VALLOX_GMBH":
        return KnxManufacturer_M_VALLOX_GMBH
    case "M_DALITEK":
        return KnxManufacturer_M_DALITEK
    case "M_ASIN":
        return KnxManufacturer_M_ASIN
    case "M_BRIDGES_INTELLIGENCE_TECHNOLOGY_INC_":
        return KnxManufacturer_M_BRIDGES_INTELLIGENCE_TECHNOLOGY_INC_
    case "M_ARBONIA":
        return KnxManufacturer_M_ARBONIA
    case "M_KERMI":
        return KnxManufacturer_M_KERMI
    case "M_EVOKNX":
        return KnxManufacturer_M_EVOKNX
    case "M_PROLUX":
        return KnxManufacturer_M_PROLUX
    case "M_CLICHOME":
        return KnxManufacturer_M_CLICHOME
    case "M_COMMAX":
        return KnxManufacturer_M_COMMAX
    case "M_EAE":
        return KnxManufacturer_M_EAE
    case "M_TENSE":
        return KnxManufacturer_M_TENSE
    case "M_SEYOUNG_ELECTRONICS":
        return KnxManufacturer_M_SEYOUNG_ELECTRONICS
    case "M_LIFEDOMUS":
        return KnxManufacturer_M_LIFEDOMUS
    case "M_EUROTRONIC_TECHNOLOGY_GMBH":
        return KnxManufacturer_M_EUROTRONIC_TECHNOLOGY_GMBH
    case "M_TCI":
        return KnxManufacturer_M_TCI
    case "M_RISHUN_ELECTRONIC":
        return KnxManufacturer_M_RISHUN_ELECTRONIC
    case "M_PAUL_HOCHKOEPPER":
        return KnxManufacturer_M_PAUL_HOCHKOEPPER
    case "M_ZIPATO":
        return KnxManufacturer_M_ZIPATO
    case "M_CM_SECURITY_GMBH_AND_CO_KG":
        return KnxManufacturer_M_CM_SECURITY_GMBH_AND_CO_KG
    case "M_QING_CABLES":
        return KnxManufacturer_M_QING_CABLES
    case "M_LABIO":
        return KnxManufacturer_M_LABIO
    case "M_COSTER_TECNOLOGIE_ELETTRONICHE_S_P_A_":
        return KnxManufacturer_M_COSTER_TECNOLOGIE_ELETTRONICHE_S_P_A_
    case "M_E_G_E":
        return KnxManufacturer_M_E_G_E
    case "M_NETXAUTOMATION":
        return KnxManufacturer_M_NETXAUTOMATION
    case "M_TECALOR":
        return KnxManufacturer_M_TECALOR
    case "M_URMET_ELECTRONICS_HUIZHOU_LTD_":
        return KnxManufacturer_M_URMET_ELECTRONICS_HUIZHOU_LTD_
    case "M_PEIYING_BUILDING_CONTROL":
        return KnxManufacturer_M_PEIYING_BUILDING_CONTROL
    case "M_ALTENBURGER_ELECTRONIC":
        return KnxManufacturer_M_ALTENBURGER_ELECTRONIC
    case "M_BPT_S_P_A__A_SOCIO_UNICO":
        return KnxManufacturer_M_BPT_S_P_A__A_SOCIO_UNICO
    case "M_KANONTEC___KANONBUS":
        return KnxManufacturer_M_KANONTEC___KANONBUS
    case "M_ISER_TECH":
        return KnxManufacturer_M_ISER_TECH
    case "M_FINELINE":
        return KnxManufacturer_M_FINELINE
    case "M_CP_ELECTRONICS_LTD":
        return KnxManufacturer_M_CP_ELECTRONICS_LTD
    case "M_NIKO_SERVODAN_AS":
        return KnxManufacturer_M_NIKO_SERVODAN_AS
    case "M_SIMON_309":
        return KnxManufacturer_M_SIMON_309
    case "M_GM_MODULAR_PVT__LTD_":
        return KnxManufacturer_M_GM_MODULAR_PVT__LTD_
    case "M_FU_CHENG_INTELLIGENCE":
        return KnxManufacturer_M_FU_CHENG_INTELLIGENCE
    case "M_NEXKON":
        return KnxManufacturer_M_NEXKON
    case "M_GRAESSLIN":
        return KnxManufacturer_M_GRAESSLIN
    case "M_FEEL_S_R_L":
        return KnxManufacturer_M_FEEL_S_R_L
    case "M_NOT_ASSIGNED_314":
        return KnxManufacturer_M_NOT_ASSIGNED_314
    case "M_SHENZHEN_FANHAI_SANJIANG_ELECTRONICS_CO___LTD_":
        return KnxManufacturer_M_SHENZHEN_FANHAI_SANJIANG_ELECTRONICS_CO___LTD_
    case "M_JIUZHOU_GREEBLE":
        return KnxManufacturer_M_JIUZHOU_GREEBLE
    case "M_AUMUELLER_AUMATIC_GMBH":
        return KnxManufacturer_M_AUMUELLER_AUMATIC_GMBH
    case "M_ETMAN_ELECTRIC":
        return KnxManufacturer_M_ETMAN_ELECTRIC
    case "M_BLACK_NOVA":
        return KnxManufacturer_M_BLACK_NOVA
    case "M_ZIDATECH_AG":
        return KnxManufacturer_M_ZIDATECH_AG
    case "M_IDGS_BVBA":
        return KnxManufacturer_M_IDGS_BVBA
    case "M_DAKANIMO":
        return KnxManufacturer_M_DAKANIMO
    case "M_SIMON_42":
        return KnxManufacturer_M_SIMON_42
    case "M_TREBOR_AUTOMATION_AB":
        return KnxManufacturer_M_TREBOR_AUTOMATION_AB
    case "M_SATEL_SP__Z_O_O_":
        return KnxManufacturer_M_SATEL_SP__Z_O_O_
    case "M_RUSSOUND__INC_":
        return KnxManufacturer_M_RUSSOUND__INC_
    case "M_MIDEA_HEATING_AND_VENTILATING_EQUIPMENT_CO_LTD":
        return KnxManufacturer_M_MIDEA_HEATING_AND_VENTILATING_EQUIPMENT_CO_LTD
    case "M_CONSORZIO_TERRANUOVA":
        return KnxManufacturer_M_CONSORZIO_TERRANUOVA
    case "M_WOLF_HEIZTECHNIK_GMBH":
        return KnxManufacturer_M_WOLF_HEIZTECHNIK_GMBH
    case "M_SONTEC":
        return KnxManufacturer_M_SONTEC
    case "M_BELCOM_CABLES_LTD_":
        return KnxManufacturer_M_BELCOM_CABLES_LTD_
    case "M_GUANGZHOU_SEAWIN_ELECTRICAL_TECHNOLOGIES_CO___LTD_":
        return KnxManufacturer_M_GUANGZHOU_SEAWIN_ELECTRICAL_TECHNOLOGIES_CO___LTD_
    case "M_ACREL":
        return KnxManufacturer_M_ACREL
    case "M_VIMAR":
        return KnxManufacturer_M_VIMAR
    case "M_FRANKE_AQUAROTTER_GMBH":
        return KnxManufacturer_M_FRANKE_AQUAROTTER_GMBH
    case "M_ORION_SYSTEMS":
        return KnxManufacturer_M_ORION_SYSTEMS
    case "M_SCHRACK_TECHNIK_GMBH":
        return KnxManufacturer_M_SCHRACK_TECHNIK_GMBH
    case "M_INSPRID":
        return KnxManufacturer_M_INSPRID
    case "M_SUNRICHER":
        return KnxManufacturer_M_SUNRICHER
    case "M_MENRED_AUTOMATION_SYSTEMSHANGHAI_CO__LTD_":
        return KnxManufacturer_M_MENRED_AUTOMATION_SYSTEMSHANGHAI_CO__LTD_
    case "M_AUREX":
        return KnxManufacturer_M_AUREX
    case "M_JOSEF_BARTHELME_GMBH_AND_CO__KG":
        return KnxManufacturer_M_JOSEF_BARTHELME_GMBH_AND_CO__KG
    case "M_ARCHITECTURE_NUMERIQUE":
        return KnxManufacturer_M_ARCHITECTURE_NUMERIQUE
    case "M_UP_GROUP":
        return KnxManufacturer_M_UP_GROUP
    case "M_MOELLER_GEBAEUDEAUTOMATION_KG":
        return KnxManufacturer_M_MOELLER_GEBAEUDEAUTOMATION_KG
    case "M_TEKNOS_AVINNO":
        return KnxManufacturer_M_TEKNOS_AVINNO
    case "M_NINGBO_DOOYA_MECHANIC_AND_ELECTRONIC_TECHNOLOGY":
        return KnxManufacturer_M_NINGBO_DOOYA_MECHANIC_AND_ELECTRONIC_TECHNOLOGY
    case "M_THERMOKON_SENSORTECHNIK_GMBH":
        return KnxManufacturer_M_THERMOKON_SENSORTECHNIK_GMBH
    case "M_BELIMO_AUTOMATION_AG":
        return KnxManufacturer_M_BELIMO_AUTOMATION_AG
    case "M_ZEHNDER_GROUP_INTERNATIONAL_AG":
        return KnxManufacturer_M_ZEHNDER_GROUP_INTERNATIONAL_AG
    case "M_SKS_KINKEL_ELEKTRONIK":
        return KnxManufacturer_M_SKS_KINKEL_ELEKTRONIK
    case "M_ECE_WURMITZER_GMBH":
        return KnxManufacturer_M_ECE_WURMITZER_GMBH
    case "M_LARS":
        return KnxManufacturer_M_LARS
    case "M_URC":
        return KnxManufacturer_M_URC
    case "M_LIGHTCONTROL":
        return KnxManufacturer_M_LIGHTCONTROL
    case "M_ALBRECHT_JUNG":
        return KnxManufacturer_M_ALBRECHT_JUNG
    case "M_ELTAKO":
        return KnxManufacturer_M_ELTAKO
    case "M_SHENZHEN_YM":
        return KnxManufacturer_M_SHENZHEN_YM
    case "M_MEAN_WELL_ENTERPRISES_CO__LTD_":
        return KnxManufacturer_M_MEAN_WELL_ENTERPRISES_CO__LTD_
    case "M_OSIX":
        return KnxManufacturer_M_OSIX
    case "M_AYPRO_TECHNOLOGY":
        return KnxManufacturer_M_AYPRO_TECHNOLOGY
    case "M_HEFEI_ECOLITE_SOFTWARE":
        return KnxManufacturer_M_HEFEI_ECOLITE_SOFTWARE
    case "M_ENNO":
        return KnxManufacturer_M_ENNO
    case "M_OHOSURE":
        return KnxManufacturer_M_OHOSURE
    case "M_GAREFOWL":
        return KnxManufacturer_M_GAREFOWL
    case "M_GEZE":
        return KnxManufacturer_M_GEZE
    case "M_LG_ELECTRONICS_INC_":
        return KnxManufacturer_M_LG_ELECTRONICS_INC_
    case "M_BOSCH_SIEMENS_HAUSHALTSGERAETE":
        return KnxManufacturer_M_BOSCH_SIEMENS_HAUSHALTSGERAETE
    case "M_SMC_INTERIORS":
        return KnxManufacturer_M_SMC_INTERIORS
    case "M_NOT_ASSIGNED_364":
        return KnxManufacturer_M_NOT_ASSIGNED_364
    case "M_SCS_CABLE":
        return KnxManufacturer_M_SCS_CABLE
    case "M_HOVAL":
        return KnxManufacturer_M_HOVAL
    case "M_CANST":
        return KnxManufacturer_M_CANST
    case "M_HANGZHOU_BERLIN":
        return KnxManufacturer_M_HANGZHOU_BERLIN
    case "M_EVN_LICHTTECHNIK":
        return KnxManufacturer_M_EVN_LICHTTECHNIK
    case "M_RUTEC":
        return KnxManufacturer_M_RUTEC
    case "M_FINDER":
        return KnxManufacturer_M_FINDER
    case "M_FUJITSU_GENERAL_LIMITED":
        return KnxManufacturer_M_FUJITSU_GENERAL_LIMITED
    case "M_RITTO_GMBHANDCO_KG":
        return KnxManufacturer_M_RITTO_GMBHANDCO_KG
    case "M_ZF_FRIEDRICHSHAFEN_AG":
        return KnxManufacturer_M_ZF_FRIEDRICHSHAFEN_AG
    case "M_CREALED":
        return KnxManufacturer_M_CREALED
    case "M_MILES_MAGIC_AUTOMATION_PRIVATE_LIMITED":
        return KnxManufacturer_M_MILES_MAGIC_AUTOMATION_PRIVATE_LIMITED
    case "M_EPlus":
        return KnxManufacturer_M_EPlus
    case "M_ITALCOND":
        return KnxManufacturer_M_ITALCOND
    case "M_SATION":
        return KnxManufacturer_M_SATION
    case "M_NEWBEST":
        return KnxManufacturer_M_NEWBEST
    case "M_GDS_DIGITAL_SYSTEMS":
        return KnxManufacturer_M_GDS_DIGITAL_SYSTEMS
    case "M_IDDERO":
        return KnxManufacturer_M_IDDERO
    case "M_MBNLED":
        return KnxManufacturer_M_MBNLED
    case "M_POWER_CONTROLS":
        return KnxManufacturer_M_POWER_CONTROLS
    case "M_VITRUM":
        return KnxManufacturer_M_VITRUM
    case "M_EKEY_BIOMETRIC_SYSTEMS_GMBH":
        return KnxManufacturer_M_EKEY_BIOMETRIC_SYSTEMS_GMBH
    case "M_AMC":
        return KnxManufacturer_M_AMC
    case "M_TRILUX_GMBH_AND_CO__KG":
        return KnxManufacturer_M_TRILUX_GMBH_AND_CO__KG
    case "M_WEXCEDO":
        return KnxManufacturer_M_WEXCEDO
    case "M_VEMER_SPA":
        return KnxManufacturer_M_VEMER_SPA
    case "M_ALEXANDER_BUERKLE_GMBH_AND_CO_KG":
        return KnxManufacturer_M_ALEXANDER_BUERKLE_GMBH_AND_CO_KG
    case "M_CITRON":
        return KnxManufacturer_M_CITRON
    case "M_SHENZHEN_HEGUANG":
        return KnxManufacturer_M_SHENZHEN_HEGUANG
    case "M_NOT_ASSIGNED_392":
        return KnxManufacturer_M_NOT_ASSIGNED_392
    case "M_ZUMTOBEL":
        return KnxManufacturer_M_ZUMTOBEL
    case "M_TRANE_B_V_B_A":
        return KnxManufacturer_M_TRANE_B_V_B_A
    case "M_CAREL":
        return KnxManufacturer_M_CAREL
    case "M_PROLITE_CONTROLS":
        return KnxManufacturer_M_PROLITE_CONTROLS
    case "M_BOSMER":
        return KnxManufacturer_M_BOSMER
    case "M_EUCHIPS":
        return KnxManufacturer_M_EUCHIPS
    case "M_CONNECT_THINKA_CONNECT":
        return KnxManufacturer_M_CONNECT_THINKA_CONNECT
    case "M_PEAKNX_A_DOGAWIST_COMPANY":
        return KnxManufacturer_M_PEAKNX_A_DOGAWIST_COMPANY
    case "M_ACEMATIC":
        return KnxManufacturer_M_ACEMATIC
    case "M_ELAUSYS":
        return KnxManufacturer_M_ELAUSYS
    case "M_ITK_ENGINEERING_AG":
        return KnxManufacturer_M_ITK_ENGINEERING_AG
    case "M_PHOENIX_CONTACT":
        return KnxManufacturer_M_PHOENIX_CONTACT
    case "M_INTEGRA_METERING_AG":
        return KnxManufacturer_M_INTEGRA_METERING_AG
    case "M_FMS_HOSPITALITY_PTE_LTD":
        return KnxManufacturer_M_FMS_HOSPITALITY_PTE_LTD
    case "M_NUVO":
        return KnxManufacturer_M_NUVO
    case "M_U__LUX_GMBH":
        return KnxManufacturer_M_U__LUX_GMBH
    case "M_BRUMBERG_LEUCHTEN":
        return KnxManufacturer_M_BRUMBERG_LEUCHTEN
    case "M_LIME":
        return KnxManufacturer_M_LIME
    case "M_GREAT_EMPIRE_INTERNATIONAL_GROUP_CO___LTD_":
        return KnxManufacturer_M_GREAT_EMPIRE_INTERNATIONAL_GROUP_CO___LTD_
    case "M_KAVOSHPISHRO_ASIA":
        return KnxManufacturer_M_KAVOSHPISHRO_ASIA
    case "M_V2_SPA":
        return KnxManufacturer_M_V2_SPA
    case "M_JOHNSON_CONTROLS":
        return KnxManufacturer_M_JOHNSON_CONTROLS
    case "M_WAGO_KONTAKTTECHNIK":
        return KnxManufacturer_M_WAGO_KONTAKTTECHNIK
    case "M_ARKUD":
        return KnxManufacturer_M_ARKUD
    case "M_IRIDIUM_LTD_":
        return KnxManufacturer_M_IRIDIUM_LTD_
    case "M_BSMART":
        return KnxManufacturer_M_BSMART
    case "M_BAB_TECHNOLOGIE_GMBH":
        return KnxManufacturer_M_BAB_TECHNOLOGIE_GMBH
    case "M_NICE_SPA":
        return KnxManufacturer_M_NICE_SPA
    case "M_REDFISH_GROUP_PTY_LTD":
        return KnxManufacturer_M_REDFISH_GROUP_PTY_LTD
    case "M_SABIANA_SPA":
        return KnxManufacturer_M_SABIANA_SPA
    case "M_UBEE_INTERACTIVE_EUROPE":
        return KnxManufacturer_M_UBEE_INTERACTIVE_EUROPE
    case "M_REXEL":
        return KnxManufacturer_M_REXEL
    case "M_GES_TEKNIK_A_S_":
        return KnxManufacturer_M_GES_TEKNIK_A_S_
    case "M_KNXPRESSO":
        return KnxManufacturer_M_KNXPRESSO
    case "M_AVE_S_P_A_":
        return KnxManufacturer_M_AVE_S_P_A_
    case "M_ZHUHAI_LTECH_TECHNOLOGY_CO___LTD_":
        return KnxManufacturer_M_ZHUHAI_LTECH_TECHNOLOGY_CO___LTD_
    case "M_ARCOM":
        return KnxManufacturer_M_ARCOM
    case "M_VIA_TECHNOLOGIES__INC_":
        return KnxManufacturer_M_VIA_TECHNOLOGIES__INC_
    case "M_FEELSMART_":
        return KnxManufacturer_M_FEELSMART_
    case "M_SUPCON":
        return KnxManufacturer_M_SUPCON
    case "M_MANIC":
        return KnxManufacturer_M_MANIC
    case "M_TDE_GMBH":
        return KnxManufacturer_M_TDE_GMBH
    case "M_NANJING_SHUFAN_INFORMATION_TECHNOLOGY_CO__LTD_":
        return KnxManufacturer_M_NANJING_SHUFAN_INFORMATION_TECHNOLOGY_CO__LTD_
    case "M_EWTECH":
        return KnxManufacturer_M_EWTECH
    case "M_WIELAND_ELECTRIC":
        return KnxManufacturer_M_WIELAND_ELECTRIC
    case "M_KLUGER_AUTOMATION_GMBH":
        return KnxManufacturer_M_KLUGER_AUTOMATION_GMBH
    case "M_JOONGANG_CONTROL":
        return KnxManufacturer_M_JOONGANG_CONTROL
    case "M_GREENCONTROLS_TECHNOLOGY_SDN__BHD_":
        return KnxManufacturer_M_GREENCONTROLS_TECHNOLOGY_SDN__BHD_
    case "M_IME_S_P_A_":
        return KnxManufacturer_M_IME_S_P_A_
    case "M_SICHUAN_HAODING":
        return KnxManufacturer_M_SICHUAN_HAODING
    case "M_MINDJAGA_LTD_":
        return KnxManufacturer_M_MINDJAGA_LTD_
    case "M_RUILI_SMART_CONTROL":
        return KnxManufacturer_M_RUILI_SMART_CONTROL
    case "M_CODESYS_GMBH":
        return KnxManufacturer_M_CODESYS_GMBH
    case "M_MOORGEN_DEUTSCHLAND_GMBH":
        return KnxManufacturer_M_MOORGEN_DEUTSCHLAND_GMBH
    case "M_CULLMANN_TECH":
        return KnxManufacturer_M_CULLMANN_TECH
    case "M_HERMANN_KLEINHUIS":
        return KnxManufacturer_M_HERMANN_KLEINHUIS
    case "M_MERCK_WINDOW_TECHNOLOGIES_B_V_":
        return KnxManufacturer_M_MERCK_WINDOW_TECHNOLOGIES_B_V_
    case "M_ABEGO":
        return KnxManufacturer_M_ABEGO
    case "M_MYGEKKO":
        return KnxManufacturer_M_MYGEKKO
    case "M_ERGO3_SARL":
        return KnxManufacturer_M_ERGO3_SARL
    case "M_STMICROELECTRONICS_INTERNATIONAL_N_V_":
        return KnxManufacturer_M_STMICROELECTRONICS_INTERNATIONAL_N_V_
    case "M_CJC_SYSTEMS":
        return KnxManufacturer_M_CJC_SYSTEMS
    case "M_SUDOKU":
        return KnxManufacturer_M_SUDOKU
    case "M_AZ_E_LITE_PTE_LTD":
        return KnxManufacturer_M_AZ_E_LITE_PTE_LTD
    case "M_ARLIGHT":
        return KnxManufacturer_M_ARLIGHT
    case "M_GRUENBECK_WASSERAUFBEREITUNG_GMBH":
        return KnxManufacturer_M_GRUENBECK_WASSERAUFBEREITUNG_GMBH
    case "M_BTICINO":
        return KnxManufacturer_M_BTICINO
    case "M_STIEBEL_ELTRON":
        return KnxManufacturer_M_STIEBEL_ELTRON
    case "M_MODULE_ELECTRONIC":
        return KnxManufacturer_M_MODULE_ELECTRONIC
    case "M_KOPLAT":
        return KnxManufacturer_M_KOPLAT
    case "M_GUANGZHOU_LETOUR_LIFE_TECHNOLOGY_CO___LTD":
        return KnxManufacturer_M_GUANGZHOU_LETOUR_LIFE_TECHNOLOGY_CO___LTD
    case "M_ILEVIA":
        return KnxManufacturer_M_ILEVIA
    case "M_LN_SYSTEMTEQ":
        return KnxManufacturer_M_LN_SYSTEMTEQ
    case "M_HISENSE_SMARTHOME":
        return KnxManufacturer_M_HISENSE_SMARTHOME
    case "M_FLINK_AUTOMATION_SYSTEM":
        return KnxManufacturer_M_FLINK_AUTOMATION_SYSTEM
    case "M_XXTER_BV":
        return KnxManufacturer_M_XXTER_BV
    case "M_LYNXUS_TECHNOLOGY":
        return KnxManufacturer_M_LYNXUS_TECHNOLOGY
    case "M_ROBOT_S_A_":
        return KnxManufacturer_M_ROBOT_S_A_
    case "M_TEHALIT":
        return KnxManufacturer_M_TEHALIT
    case "M_SHENZHEN_ATTE_SMART_LIFE_CO__LTD_":
        return KnxManufacturer_M_SHENZHEN_ATTE_SMART_LIFE_CO__LTD_
    case "M_NOBLESSE":
        return KnxManufacturer_M_NOBLESSE
    case "M_ADVANCED_DEVICES":
        return KnxManufacturer_M_ADVANCED_DEVICES
    case "M_ATRINA_BUILDING_AUTOMATION_CO__LTD":
        return KnxManufacturer_M_ATRINA_BUILDING_AUTOMATION_CO__LTD
    case "M_GUANGDONG_DAMING_LAFFEY_ELECTRIC_CO___LTD_":
        return KnxManufacturer_M_GUANGDONG_DAMING_LAFFEY_ELECTRIC_CO___LTD_
    case "M_WESTERSTRAND_URFABRIK_AB":
        return KnxManufacturer_M_WESTERSTRAND_URFABRIK_AB
    case "M_CONTROL4_CORPORATE":
        return KnxManufacturer_M_CONTROL4_CORPORATE
    case "M_ONTROL":
        return KnxManufacturer_M_ONTROL
    case "M_STARNET":
        return KnxManufacturer_M_STARNET
    case "M_BETA_CAVI":
        return KnxManufacturer_M_BETA_CAVI
    case "M_THEBEN_AG":
        return KnxManufacturer_M_THEBEN_AG
    case "M_EASEMORE":
        return KnxManufacturer_M_EASEMORE
    case "M_VIVALDI_SRL":
        return KnxManufacturer_M_VIVALDI_SRL
    case "M_GREE_ELECTRIC_APPLIANCES_INC__OF_ZHUHAI":
        return KnxManufacturer_M_GREE_ELECTRIC_APPLIANCES_INC__OF_ZHUHAI
    case "M_HWISCON":
        return KnxManufacturer_M_HWISCON
    case "M_SHANGHAI_ELECON_INTELLIGENT_TECHNOLOGY_CO___LTD_":
        return KnxManufacturer_M_SHANGHAI_ELECON_INTELLIGENT_TECHNOLOGY_CO___LTD_
    case "M_KAMPMANN":
        return KnxManufacturer_M_KAMPMANN
    case "M_IMPOLUX_GMBH_LEDIMAX":
        return KnxManufacturer_M_IMPOLUX_GMBH_LEDIMAX
    case "M_EVAUX":
        return KnxManufacturer_M_EVAUX
    case "M_WEBRO_CABLES_AND_CONNECTORS_LIMITED":
        return KnxManufacturer_M_WEBRO_CABLES_AND_CONNECTORS_LIMITED
    case "M_SHANGHAI_E_TECH_SOLUTION":
        return KnxManufacturer_M_SHANGHAI_E_TECH_SOLUTION
    case "M_WILHELM_RUTENBECK":
        return KnxManufacturer_M_WILHELM_RUTENBECK
    case "M_GUANGZHOU_HOKO_ELECTRIC_CO__LTD_":
        return KnxManufacturer_M_GUANGZHOU_HOKO_ELECTRIC_CO__LTD_
    case "M_LAMMIN_HIGH_TECH_CO__LTD":
        return KnxManufacturer_M_LAMMIN_HIGH_TECH_CO__LTD
    case "M_SHENZHEN_MERRYTEK_TECHNOLOGY_CO___LTD":
        return KnxManufacturer_M_SHENZHEN_MERRYTEK_TECHNOLOGY_CO___LTD
    case "M_I_LUXUS":
        return KnxManufacturer_M_I_LUXUS
    case "M_ELMOS_SEMICONDUCTOR_AG":
        return KnxManufacturer_M_ELMOS_SEMICONDUCTOR_AG
    case "M_EMCOM_TECHNOLOGY_INC":
        return KnxManufacturer_M_EMCOM_TECHNOLOGY_INC
    case "M_PROJECT_INNOVATIONS_GMBH":
        return KnxManufacturer_M_PROJECT_INNOVATIONS_GMBH
    case "M_ITC":
        return KnxManufacturer_M_ITC
    case "M_ABB_LV_INSTALLATION_MATERIALS_COMPANY_LTD__BEIJING":
        return KnxManufacturer_M_ABB_LV_INSTALLATION_MATERIALS_COMPANY_LTD__BEIJING
    case "M_MAICO":
        return KnxManufacturer_M_MAICO
    case "M_WINKHAUS":
        return KnxManufacturer_M_WINKHAUS
    case "M_ELAN_SRL":
        return KnxManufacturer_M_ELAN_SRL
    case "M_MINHHA_TECHNOLOGY_CO__LTD":
        return KnxManufacturer_M_MINHHA_TECHNOLOGY_CO__LTD
    case "M_ZHEJIANG_TIANJIE_INDUSTRIAL_CORP_":
        return KnxManufacturer_M_ZHEJIANG_TIANJIE_INDUSTRIAL_CORP_
    case "M_IAUTOMATION_PTY_LIMITED":
        return KnxManufacturer_M_IAUTOMATION_PTY_LIMITED
    case "M_EXTRON":
        return KnxManufacturer_M_EXTRON
    case "M_FREEDOMPRO":
        return KnxManufacturer_M_FREEDOMPRO
    case "M_ONEHOME":
        return KnxManufacturer_M_ONEHOME
    case "M_EOS_SAUNATECHNIK_GMBH":
        return KnxManufacturer_M_EOS_SAUNATECHNIK_GMBH
    case "M_KUSATEK_GMBH":
        return KnxManufacturer_M_KUSATEK_GMBH
    case "M_EISBAER_SCADA":
        return KnxManufacturer_M_EISBAER_SCADA
    case "M_ROBERT_BOSCH":
        return KnxManufacturer_M_ROBERT_BOSCH
    case "M_AUTOMATISMI_BENINCA_S_P_A_":
        return KnxManufacturer_M_AUTOMATISMI_BENINCA_S_P_A_
    case "M_BLENDOM":
        return KnxManufacturer_M_BLENDOM
    case "M_MADEL_AIR_TECHNICAL_DIFFUSION":
        return KnxManufacturer_M_MADEL_AIR_TECHNICAL_DIFFUSION
    case "M_NIKO":
        return KnxManufacturer_M_NIKO
    case "M_BOSCH_REXROTH_AG":
        return KnxManufacturer_M_BOSCH_REXROTH_AG
    case "M_CANDM_PRODUCTS":
        return KnxManufacturer_M_CANDM_PRODUCTS
    case "M_HOERMANN_KG_VERKAUFSGESELLSCHAFT":
        return KnxManufacturer_M_HOERMANN_KG_VERKAUFSGESELLSCHAFT
    case "M_SHANGHAI_RAJAYASA_CO__LTD":
        return KnxManufacturer_M_SHANGHAI_RAJAYASA_CO__LTD
    case "M_SUZUKI":
        return KnxManufacturer_M_SUZUKI
    case "M_SILENT_GLISS_INTERNATIONAL_LTD_":
        return KnxManufacturer_M_SILENT_GLISS_INTERNATIONAL_LTD_
    case "M_SOMFY":
        return KnxManufacturer_M_SOMFY
    case "M_BEE_CONTROLS_ADGSC_GROUP":
        return KnxManufacturer_M_BEE_CONTROLS_ADGSC_GROUP
    case "M_XDTECGMBH":
        return KnxManufacturer_M_XDTECGMBH
    case "M_OSRAM":
        return KnxManufacturer_M_OSRAM
    case "M_LEBENOR":
        return KnxManufacturer_M_LEBENOR
    case "M_AUTOMANENG":
        return KnxManufacturer_M_AUTOMANENG
    case "M_HONEYWELL_AUTOMATION_SOLUTION_CONTROLCHINA":
        return KnxManufacturer_M_HONEYWELL_AUTOMATION_SOLUTION_CONTROLCHINA
    case "M_HANGZHOU_BINTHEN_INTELLIGENCE_TECHNOLOGY_CO__LTD":
        return KnxManufacturer_M_HANGZHOU_BINTHEN_INTELLIGENCE_TECHNOLOGY_CO__LTD
    case "M_ETA_HEIZTECHNIK":
        return KnxManufacturer_M_ETA_HEIZTECHNIK
    case "M_DIVUS_GMBH":
        return KnxManufacturer_M_DIVUS_GMBH
    case "M_NANJING_TAIJIESAI_INTELLIGENT_TECHNOLOGY_CO__LTD_":
        return KnxManufacturer_M_NANJING_TAIJIESAI_INTELLIGENT_TECHNOLOGY_CO__LTD_
    case "M_WOERTZ":
        return KnxManufacturer_M_WOERTZ
    case "M_LUNATONE":
        return KnxManufacturer_M_LUNATONE
    case "M_ZHEJIANG_SCTECH_BUILDING_INTELLIGENT":
        return KnxManufacturer_M_ZHEJIANG_SCTECH_BUILDING_INTELLIGENT
    case "M_FOSHAN_QITE_TECHNOLOGY_CO___LTD_":
        return KnxManufacturer_M_FOSHAN_QITE_TECHNOLOGY_CO___LTD_
    case "M_NOKE":
        return KnxManufacturer_M_NOKE
    case "M_LANDCOM":
        return KnxManufacturer_M_LANDCOM
    case "M_STORK_AS":
        return KnxManufacturer_M_STORK_AS
    case "M_HANGZHOU_SHENDU_TECHNOLOGY_CO___LTD_":
        return KnxManufacturer_M_HANGZHOU_SHENDU_TECHNOLOGY_CO___LTD_
    case "M_COOLAUTOMATION":
        return KnxManufacturer_M_COOLAUTOMATION
    case "M_APRSTERN":
        return KnxManufacturer_M_APRSTERN
    case "M_SONNEN":
        return KnxManufacturer_M_SONNEN
    case "M_VIESSMANN_WERKE":
        return KnxManufacturer_M_VIESSMANN_WERKE
    case "M_DNAKE":
        return KnxManufacturer_M_DNAKE
    case "M_NEUBERGER_GEBAEUDEAUTOMATION_GMBH":
        return KnxManufacturer_M_NEUBERGER_GEBAEUDEAUTOMATION_GMBH
    case "M_STILIGER":
        return KnxManufacturer_M_STILIGER
    case "M_BERGHOF_AUTOMATION_GMBH":
        return KnxManufacturer_M_BERGHOF_AUTOMATION_GMBH
    case "M_TOTAL_AUTOMATION_AND_CONTROLS_GMBH":
        return KnxManufacturer_M_TOTAL_AUTOMATION_AND_CONTROLS_GMBH
    case "M_DOVIT":
        return KnxManufacturer_M_DOVIT
    case "M_INSTALIGHTING_GMBH":
        return KnxManufacturer_M_INSTALIGHTING_GMBH
    case "M_UNI_TEC":
        return KnxManufacturer_M_UNI_TEC
    case "M_CASATUNES":
        return KnxManufacturer_M_CASATUNES
    case "M_EMT":
        return KnxManufacturer_M_EMT
    case "M_IMI_HYDRONIC_ENGINEERING":
        return KnxManufacturer_M_IMI_HYDRONIC_ENGINEERING
    case "M_SENFFICIENT":
        return KnxManufacturer_M_SENFFICIENT
    case "M_AUROLITE_ELECTRICAL_PANYU_GUANGZHOU_LIMITED":
        return KnxManufacturer_M_AUROLITE_ELECTRICAL_PANYU_GUANGZHOU_LIMITED
    case "M_ABB_XIAMEN_SMART_TECHNOLOGY_CO___LTD_":
        return KnxManufacturer_M_ABB_XIAMEN_SMART_TECHNOLOGY_CO___LTD_
    case "M_SAMSON_ELECTRIC_WIRE":
        return KnxManufacturer_M_SAMSON_ELECTRIC_WIRE
    case "M_T_TOUCHING":
        return KnxManufacturer_M_T_TOUCHING
    case "M_CORE_SMART_HOME":
        return KnxManufacturer_M_CORE_SMART_HOME
    case "M_GREENCONNECT_SOLUTIONS_SA":
        return KnxManufacturer_M_GREENCONNECT_SOLUTIONS_SA
    case "M_ELETTRONICA_CONDUTTORI":
        return KnxManufacturer_M_ELETTRONICA_CONDUTTORI
    case "M_MKFC":
        return KnxManufacturer_M_MKFC
    case "M_AUTOMATIONPlus":
        return KnxManufacturer_M_AUTOMATIONPlus
    case "M_BERKER":
        return KnxManufacturer_M_BERKER
    case "M_JOH__VAILLANT":
        return KnxManufacturer_M_JOH__VAILLANT
    case "M_BLUE_AND_RED":
        return KnxManufacturer_M_BLUE_AND_RED
    case "M_FROGBLUE":
        return KnxManufacturer_M_FROGBLUE
    case "M_SAVESOR":
        return KnxManufacturer_M_SAVESOR
    case "M_APP_TECH":
        return KnxManufacturer_M_APP_TECH
    case "M_SENSORTEC_AG":
        return KnxManufacturer_M_SENSORTEC_AG
    case "M_NYSA_TECHNOLOGY_AND_SOLUTIONS":
        return KnxManufacturer_M_NYSA_TECHNOLOGY_AND_SOLUTIONS
    case "M_FARADITE":
        return KnxManufacturer_M_FARADITE
    case "M_OPTIMUS":
        return KnxManufacturer_M_OPTIMUS
    case "M_KTS_S_R_L_":
        return KnxManufacturer_M_KTS_S_R_L_
    case "M_RAMCRO_SPA":
        return KnxManufacturer_M_RAMCRO_SPA
    case "M_AMP_DEUTSCHLAND":
        return KnxManufacturer_M_AMP_DEUTSCHLAND
    case "M_WUHAN_WISECREATE_UNIVERSE_TECHNOLOGY_CO___LTD":
        return KnxManufacturer_M_WUHAN_WISECREATE_UNIVERSE_TECHNOLOGY_CO___LTD
    case "M_BEMI_SMART_HOME_LTD":
        return KnxManufacturer_M_BEMI_SMART_HOME_LTD
    case "M_ARDOMUS":
        return KnxManufacturer_M_ARDOMUS
    case "M_CHANGXING":
        return KnxManufacturer_M_CHANGXING
    case "M_E_CONTROLS":
        return KnxManufacturer_M_E_CONTROLS
    case "M_AIB_TECHNOLOGY":
        return KnxManufacturer_M_AIB_TECHNOLOGY
    case "M_NVC":
        return KnxManufacturer_M_NVC
    case "M_KBOX":
        return KnxManufacturer_M_KBOX
    case "M_CNS":
        return KnxManufacturer_M_CNS
    case "M_TYBA":
        return KnxManufacturer_M_TYBA
    case "M_BOSCH_THERMOTECHNIK_GMBH":
        return KnxManufacturer_M_BOSCH_THERMOTECHNIK_GMBH
    case "M_ATREL":
        return KnxManufacturer_M_ATREL
    case "M_SIMON_ELECTRIC_CHINA_CO___LTD":
        return KnxManufacturer_M_SIMON_ELECTRIC_CHINA_CO___LTD
    case "M_KORDZ_GROUP":
        return KnxManufacturer_M_KORDZ_GROUP
    case "M_ND_ELECTRIC":
        return KnxManufacturer_M_ND_ELECTRIC
    case "M_CONTROLIUM":
        return KnxManufacturer_M_CONTROLIUM
    case "M_FAMO_GMBH_AND_CO__KG":
        return KnxManufacturer_M_FAMO_GMBH_AND_CO__KG
    case "M_CDN_SMART":
        return KnxManufacturer_M_CDN_SMART
    case "M_HESTON":
        return KnxManufacturer_M_HESTON
    case "M_ESLA_CONEXIONES_S_L_":
        return KnxManufacturer_M_ESLA_CONEXIONES_S_L_
    case "M_WEISHAUPT":
        return KnxManufacturer_M_WEISHAUPT
    case "M_SEF___ECOTEC":
        return KnxManufacturer_M_SEF___ECOTEC
    case "M_ASTRUM_TECHNOLOGY":
        return KnxManufacturer_M_ASTRUM_TECHNOLOGY
    case "M_WUERTH_ELEKTRONIK_STELVIO_KONTEK_S_P_A_":
        return KnxManufacturer_M_WUERTH_ELEKTRONIK_STELVIO_KONTEK_S_P_A_
    case "M_NANOTECO_CORPORATION":
        return KnxManufacturer_M_NANOTECO_CORPORATION
    case "M_NIETIAN":
        return KnxManufacturer_M_NIETIAN
    case "M_SUMSIR":
        return KnxManufacturer_M_SUMSIR
    case "M_ORBIS_TECNOLOGIA_ELECTRICA_SA":
        return KnxManufacturer_M_ORBIS_TECNOLOGIA_ELECTRICA_SA
    case "M_ABB___RESERVED":
        return KnxManufacturer_M_ABB___RESERVED
    case "M_BUSCH_JAEGER_ELEKTRO___RESERVED":
        return KnxManufacturer_M_BUSCH_JAEGER_ELEKTRO___RESERVED
    case "M_DORMA_GMBH_Plus_CO__KG":
        return KnxManufacturer_M_DORMA_GMBH_Plus_CO__KG
    case "M_WINDOWMASTER_AS":
        return KnxManufacturer_M_WINDOWMASTER_AS
    case "M_WALTHER_WERKE":
        return KnxManufacturer_M_WALTHER_WERKE
    case "M_ORAS":
        return KnxManufacturer_M_ORAS
    case "M_DAETWYLER":
        return KnxManufacturer_M_DAETWYLER
    case "M_ELECTRAK":
        return KnxManufacturer_M_ELECTRAK
    case "M_BUSCH_JAEGER_ELEKTRO":
        return KnxManufacturer_M_BUSCH_JAEGER_ELEKTRO
    case "M_TECHEM":
        return KnxManufacturer_M_TECHEM
    case "M_SCHNEIDER_ELECTRIC_INDUSTRIES_SAS":
        return KnxManufacturer_M_SCHNEIDER_ELECTRIC_INDUSTRIES_SAS
    case "M_WHD_WILHELM_HUBER_Plus_SOEHNE":
        return KnxManufacturer_M_WHD_WILHELM_HUBER_Plus_SOEHNE
    case "M_BISCHOFF_ELEKTRONIK":
        return KnxManufacturer_M_BISCHOFF_ELEKTRONIK
    case "M_JEPAZ":
        return KnxManufacturer_M_JEPAZ
    case "M_RTS_AUTOMATION":
        return KnxManufacturer_M_RTS_AUTOMATION
    case "M_EIBMARKT_GMBH":
        return KnxManufacturer_M_EIBMARKT_GMBH
    case "M_WAREMA_RENKHOFF_SE":
        return KnxManufacturer_M_WAREMA_RENKHOFF_SE
    case "M_EELECTRON":
        return KnxManufacturer_M_EELECTRON
    case "M_BELDEN_WIRE_AND_CABLE_B_V_":
        return KnxManufacturer_M_BELDEN_WIRE_AND_CABLE_B_V_
    case "M_GIRA_GIERSIEPEN":
        return KnxManufacturer_M_GIRA_GIERSIEPEN
    case "M_BECKER_ANTRIEBE_GMBH":
        return KnxManufacturer_M_BECKER_ANTRIEBE_GMBH
    case "M_J_STEHLEPlusSOEHNE_GMBH":
        return KnxManufacturer_M_J_STEHLEPlusSOEHNE_GMBH
    case "M_AGFEO":
        return KnxManufacturer_M_AGFEO
    case "M_ZENNIO":
        return KnxManufacturer_M_ZENNIO
    case "M_TAPKO_TECHNOLOGIES":
        return KnxManufacturer_M_TAPKO_TECHNOLOGIES
    case "M_HDL":
        return KnxManufacturer_M_HDL
    case "M_UPONOR":
        return KnxManufacturer_M_UPONOR
    case "M_SE_LIGHTMANAGEMENT_AG":
        return KnxManufacturer_M_SE_LIGHTMANAGEMENT_AG
    case "M_ARCUS_EDS":
        return KnxManufacturer_M_ARCUS_EDS
    case "M_INTESIS":
        return KnxManufacturer_M_INTESIS
    case "M_HAGER_ELECTRO":
        return KnxManufacturer_M_HAGER_ELECTRO
    case "M_HERHOLDT_CONTROLS_SRL":
        return KnxManufacturer_M_HERHOLDT_CONTROLS_SRL
    case "M_NIKO_ZUBLIN":
        return KnxManufacturer_M_NIKO_ZUBLIN
    case "M_DURABLE_TECHNOLOGIES":
        return KnxManufacturer_M_DURABLE_TECHNOLOGIES
    case "M_INNOTEAM":
        return KnxManufacturer_M_INNOTEAM
    case "M_ISE_GMBH":
        return KnxManufacturer_M_ISE_GMBH
    case "M_TEAM_FOR_TRONICS":
        return KnxManufacturer_M_TEAM_FOR_TRONICS
    case "M_CIAT":
        return KnxManufacturer_M_CIAT
    case "M_REMEHA_BV":
        return KnxManufacturer_M_REMEHA_BV
    case "M_ESYLUX":
        return KnxManufacturer_M_ESYLUX
    case "M_BASALTE":
        return KnxManufacturer_M_BASALTE
    case "M_INSTA_GMBH":
        return KnxManufacturer_M_INSTA_GMBH
    case "M_VESTAMATIC":
        return KnxManufacturer_M_VESTAMATIC
    case "M_MDT_TECHNOLOGIES":
        return KnxManufacturer_M_MDT_TECHNOLOGIES
    case "M_WARENDORFER_KUECHEN_GMBH":
        return KnxManufacturer_M_WARENDORFER_KUECHEN_GMBH
    case "M_VIDEO_STAR":
        return KnxManufacturer_M_VIDEO_STAR
    case "M_SITEK":
        return KnxManufacturer_M_SITEK
    case "M_CONTROLTRONIC":
        return KnxManufacturer_M_CONTROLTRONIC
    case "M_FUNCTION_TECHNOLOGY":
        return KnxManufacturer_M_FUNCTION_TECHNOLOGY
    case "M_AMX":
        return KnxManufacturer_M_AMX
    case "M_ELDAT":
        return KnxManufacturer_M_ELDAT
    case "M_PANASONIC":
        return KnxManufacturer_M_PANASONIC
    }
    return 0
}

func CastKnxManufacturer(structType interface{}) KnxManufacturer {
    castFunc := func(typ interface{}) KnxManufacturer {
        if sKnxManufacturer, ok := typ.(KnxManufacturer); ok {
            return sKnxManufacturer
        }
        return 0
    }
    return castFunc(structType)
}

func (m KnxManufacturer) LengthInBits() uint16 {
    return 16
}

func (m KnxManufacturer) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func KnxManufacturerParse(io *utils.ReadBuffer) (KnxManufacturer, error) {
    val, err := io.ReadUint16(16)
    if err != nil {
        return 0, nil
    }
    return KnxManufacturerByValue(val), nil
}

func (e KnxManufacturer) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint16(16, uint16(e))
    return err
}

func (e KnxManufacturer) String() string {
    switch e {
    case KnxManufacturer_M_UNKNOWN:
        return "M_UNKNOWN"
    case KnxManufacturer_M_SIEMENS:
        return "M_SIEMENS"
    case KnxManufacturer_M_LEGRAND_APPAREILLAGE_ELECTRIQUE:
        return "M_LEGRAND_APPAREILLAGE_ELECTRIQUE"
    case KnxManufacturer_M_PULSE_TECHNOLOGIES:
        return "M_PULSE_TECHNOLOGIES"
    case KnxManufacturer_M_CRESTRON:
        return "M_CRESTRON"
    case KnxManufacturer_M_STEINEL_PROFESSIONAL:
        return "M_STEINEL_PROFESSIONAL"
    case KnxManufacturer_M_BILTON_LED_LIGHTING:
        return "M_BILTON_LED_LIGHTING"
    case KnxManufacturer_M_DENRO_AG:
        return "M_DENRO_AG"
    case KnxManufacturer_M_GEPRO:
        return "M_GEPRO"
    case KnxManufacturer_M_PREUSSEN_AUTOMATION:
        return "M_PREUSSEN_AUTOMATION"
    case KnxManufacturer_M_ZOPPAS_INDUSTRIES:
        return "M_ZOPPAS_INDUSTRIES"
    case KnxManufacturer_M_MACTECH:
        return "M_MACTECH"
    case KnxManufacturer_M_TECHNO_TREND:
        return "M_TECHNO_TREND"
    case KnxManufacturer_M_MERTEN:
        return "M_MERTEN"
    case KnxManufacturer_M_FS_CABLES:
        return "M_FS_CABLES"
    case KnxManufacturer_M_DELTA_DORE:
        return "M_DELTA_DORE"
    case KnxManufacturer_M_EISSOUND:
        return "M_EISSOUND"
    case KnxManufacturer_M_CISCO:
        return "M_CISCO"
    case KnxManufacturer_M_DINUY:
        return "M_DINUY"
    case KnxManufacturer_M_IKNIX:
        return "M_IKNIX"
    case KnxManufacturer_M_RADEMACHER_GERAETE_ELEKTRONIK_GMBH:
        return "M_RADEMACHER_GERAETE_ELEKTRONIK_GMBH"
    case KnxManufacturer_M_EGI_ELECTROACUSTICA_GENERAL_IBERICA:
        return "M_EGI_ELECTROACUSTICA_GENERAL_IBERICA"
    case KnxManufacturer_M_BES___INGENIUM:
        return "M_BES___INGENIUM"
    case KnxManufacturer_M_ELABNET:
        return "M_ELABNET"
    case KnxManufacturer_M_ABB_SPA_SACE_DIVISION:
        return "M_ABB_SPA_SACE_DIVISION"
    case KnxManufacturer_M_BLUMOTIX:
        return "M_BLUMOTIX"
    case KnxManufacturer_M_HUNTER_DOUGLAS:
        return "M_HUNTER_DOUGLAS"
    case KnxManufacturer_M_APRICUM:
        return "M_APRICUM"
    case KnxManufacturer_M_TIANSU_AUTOMATION:
        return "M_TIANSU_AUTOMATION"
    case KnxManufacturer_M_BUBENDORFF:
        return "M_BUBENDORFF"
    case KnxManufacturer_M_MBS_GMBH:
        return "M_MBS_GMBH"
    case KnxManufacturer_M_ENERTEX_BAYERN_GMBH:
        return "M_ENERTEX_BAYERN_GMBH"
    case KnxManufacturer_M_BMS:
        return "M_BMS"
    case KnxManufacturer_M_SINAPSI:
        return "M_SINAPSI"
    case KnxManufacturer_M_EMBEDDED_SYSTEMS_SIA:
        return "M_EMBEDDED_SYSTEMS_SIA"
    case KnxManufacturer_M_SIEDLE_AND_SOEHNE:
        return "M_SIEDLE_AND_SOEHNE"
    case KnxManufacturer_M_KNX1:
        return "M_KNX1"
    case KnxManufacturer_M_TOKKA:
        return "M_TOKKA"
    case KnxManufacturer_M_NANOSENSE:
        return "M_NANOSENSE"
    case KnxManufacturer_M_PEAR_AUTOMATION_GMBH:
        return "M_PEAR_AUTOMATION_GMBH"
    case KnxManufacturer_M_DGA:
        return "M_DGA"
    case KnxManufacturer_M_LUTRON:
        return "M_LUTRON"
    case KnxManufacturer_M_AIRZONE___ALTRA:
        return "M_AIRZONE___ALTRA"
    case KnxManufacturer_M_LITHOSS_DESIGN_SWITCHES:
        return "M_LITHOSS_DESIGN_SWITCHES"
    case KnxManufacturer_M_THREEATEL:
        return "M_THREEATEL"
    case KnxManufacturer_M_PHILIPS_CONTROLS:
        return "M_PHILIPS_CONTROLS"
    case KnxManufacturer_M_EBERLE:
        return "M_EBERLE"
    case KnxManufacturer_M_VELUX_AS:
        return "M_VELUX_AS"
    case KnxManufacturer_M_LOYTEC:
        return "M_LOYTEC"
    case KnxManufacturer_M_EKINEX_S_P_A_:
        return "M_EKINEX_S_P_A_"
    case KnxManufacturer_M_SIRLAN_TECHNOLOGIES:
        return "M_SIRLAN_TECHNOLOGIES"
    case KnxManufacturer_M_PROKNX_SAS:
        return "M_PROKNX_SAS"
    case KnxManufacturer_M_IT_GMBH:
        return "M_IT_GMBH"
    case KnxManufacturer_M_RENSON:
        return "M_RENSON"
    case KnxManufacturer_M_HEP_GROUP:
        return "M_HEP_GROUP"
    case KnxManufacturer_M_BALMART:
        return "M_BALMART"
    case KnxManufacturer_M_GFS_GMBH:
        return "M_GFS_GMBH"
    case KnxManufacturer_M_GEWISS:
        return "M_GEWISS"
    case KnxManufacturer_M_SCHENKER_STOREN_AG:
        return "M_SCHENKER_STOREN_AG"
    case KnxManufacturer_M_ALGODUE_ELETTRONICA_S_R_L_:
        return "M_ALGODUE_ELETTRONICA_S_R_L_"
    case KnxManufacturer_M_ABB_FRANCE:
        return "M_ABB_FRANCE"
    case KnxManufacturer_M_MAINTRONIC:
        return "M_MAINTRONIC"
    case KnxManufacturer_M_VANTAGE:
        return "M_VANTAGE"
    case KnxManufacturer_M_FORESIS:
        return "M_FORESIS"
    case KnxManufacturer_M_RESEARCH_AND_PRODUCTION_ASSOCIATION_SEM:
        return "M_RESEARCH_AND_PRODUCTION_ASSOCIATION_SEM"
    case KnxManufacturer_M_WEINZIERL_ENGINEERING_GMBH:
        return "M_WEINZIERL_ENGINEERING_GMBH"
    case KnxManufacturer_M_MOEHLENHOFF_WAERMETECHNIK_GMBH:
        return "M_MOEHLENHOFF_WAERMETECHNIK_GMBH"
    case KnxManufacturer_M_PKC_GROUP_OYJ:
        return "M_PKC_GROUP_OYJ"
    case KnxManufacturer_M_ALBERT_ACKERMANN:
        return "M_ALBERT_ACKERMANN"
    case KnxManufacturer_M_B_E_G_:
        return "M_B_E_G_"
    case KnxManufacturer_M_ELSNER_ELEKTRONIK_GMBH:
        return "M_ELSNER_ELEKTRONIK_GMBH"
    case KnxManufacturer_M_SIEMENS_BUILDING_TECHNOLOGIES_HKCHINA_LTD_:
        return "M_SIEMENS_BUILDING_TECHNOLOGIES_HKCHINA_LTD_"
    case KnxManufacturer_M_EUTRAC:
        return "M_EUTRAC"
    case KnxManufacturer_M_GUSTAV_HENSEL_GMBH_AND_CO__KG:
        return "M_GUSTAV_HENSEL_GMBH_AND_CO__KG"
    case KnxManufacturer_M_GARO_AB:
        return "M_GARO_AB"
    case KnxManufacturer_M_WALDMANN_LICHTTECHNIK:
        return "M_WALDMANN_LICHTTECHNIK"
    case KnxManufacturer_M_SCHUECO:
        return "M_SCHUECO"
    case KnxManufacturer_M_EMU:
        return "M_EMU"
    case KnxManufacturer_M_JNET_SYSTEMS_AG:
        return "M_JNET_SYSTEMS_AG"
    case KnxManufacturer_M_SCHUPA_GMBH:
        return "M_SCHUPA_GMBH"
    case KnxManufacturer_M_TOTAL_SOLUTION_GMBH:
        return "M_TOTAL_SOLUTION_GMBH"
    case KnxManufacturer_M_O_Y_L__ELECTRONICS:
        return "M_O_Y_L__ELECTRONICS"
    case KnxManufacturer_M_GALAX_SYSTEM:
        return "M_GALAX_SYSTEM"
    case KnxManufacturer_M_DISCH:
        return "M_DISCH"
    case KnxManufacturer_M_AUCOTEAM:
        return "M_AUCOTEAM"
    case KnxManufacturer_M_LUXMATE_CONTROLS:
        return "M_LUXMATE_CONTROLS"
    case KnxManufacturer_M_DANFOSS:
        return "M_DANFOSS"
    case KnxManufacturer_M_AST_GMBH:
        return "M_AST_GMBH"
    case KnxManufacturer_M_WILA_LEUCHTEN:
        return "M_WILA_LEUCHTEN"
    case KnxManufacturer_M_BPlusB_AUTOMATIONS__UND_STEUERUNGSTECHNIK:
        return "M_BPlusB_AUTOMATIONS__UND_STEUERUNGSTECHNIK"
    case KnxManufacturer_M_ABB_SCHWEIZ:
        return "M_ABB_SCHWEIZ"
    case KnxManufacturer_M_LINGG_AND_JANKE:
        return "M_LINGG_AND_JANKE"
    case KnxManufacturer_M_SAUTER:
        return "M_SAUTER"
    case KnxManufacturer_M_SIMU:
        return "M_SIMU"
    case KnxManufacturer_M_THEBEN_HTS_AG:
        return "M_THEBEN_HTS_AG"
    case KnxManufacturer_M_AMANN_GMBH:
        return "M_AMANN_GMBH"
    case KnxManufacturer_M_BERG_ENERGIEKONTROLLSYSTEME_GMBH:
        return "M_BERG_ENERGIEKONTROLLSYSTEME_GMBH"
    case KnxManufacturer_M_HUEPPE_FORM_SONNENSCHUTZSYSTEME_GMBH:
        return "M_HUEPPE_FORM_SONNENSCHUTZSYSTEME_GMBH"
    case KnxManufacturer_M_OVENTROP_KG:
        return "M_OVENTROP_KG"
    case KnxManufacturer_M_GRIESSER_AG:
        return "M_GRIESSER_AG"
    case KnxManufacturer_M_IPAS_GMBH:
        return "M_IPAS_GMBH"
    case KnxManufacturer_M_FELLER:
        return "M_FELLER"
    case KnxManufacturer_M_ELERO_GMBH:
        return "M_ELERO_GMBH"
    case KnxManufacturer_M_ARDAN_PRODUCTION_AND_INDUSTRIAL_CONTROLS_LTD_:
        return "M_ARDAN_PRODUCTION_AND_INDUSTRIAL_CONTROLS_LTD_"
    case KnxManufacturer_M_METEC_MESSTECHNIK_GMBH:
        return "M_METEC_MESSTECHNIK_GMBH"
    case KnxManufacturer_M_ELKA_ELEKTRONIK_GMBH:
        return "M_ELKA_ELEKTRONIK_GMBH"
    case KnxManufacturer_M_ELEKTROANLAGEN_D__NAGEL:
        return "M_ELEKTROANLAGEN_D__NAGEL"
    case KnxManufacturer_M_TRIDONIC_BAUELEMENTE_GMBH:
        return "M_TRIDONIC_BAUELEMENTE_GMBH"
    case KnxManufacturer_M_STENGLER_GESELLSCHAFT:
        return "M_STENGLER_GESELLSCHAFT"
    case KnxManufacturer_M_SCHNEIDER_ELECTRIC_MG:
        return "M_SCHNEIDER_ELECTRIC_MG"
    case KnxManufacturer_M_KNX_ASSOCIATION:
        return "M_KNX_ASSOCIATION"
    case KnxManufacturer_M_VIVO:
        return "M_VIVO"
    case KnxManufacturer_M_ABB:
        return "M_ABB"
    case KnxManufacturer_M_GLAMOX_AS:
        return "M_GLAMOX_AS"
    case KnxManufacturer_M_HUGO_MUELLER_GMBH_AND_CO_KG:
        return "M_HUGO_MUELLER_GMBH_AND_CO_KG"
    case KnxManufacturer_M_SIEMENS_HVAC:
        return "M_SIEMENS_HVAC"
    case KnxManufacturer_M_APT:
        return "M_APT"
    case KnxManufacturer_M_HIGHDOM:
        return "M_HIGHDOM"
    case KnxManufacturer_M_TOP_SERVICES:
        return "M_TOP_SERVICES"
    case KnxManufacturer_M_AMBIHOME:
        return "M_AMBIHOME"
    case KnxManufacturer_M_DATEC_ELECTRONIC_AG:
        return "M_DATEC_ELECTRONIC_AG"
    case KnxManufacturer_M_ABUS_SECURITY_CENTER:
        return "M_ABUS_SECURITY_CENTER"
    case KnxManufacturer_M_LITE_PUTER:
        return "M_LITE_PUTER"
    case KnxManufacturer_M_TANTRON_ELECTRONIC:
        return "M_TANTRON_ELECTRONIC"
    case KnxManufacturer_M_DEHN_AND_SOEHNE:
        return "M_DEHN_AND_SOEHNE"
    case KnxManufacturer_M_INTERRA:
        return "M_INTERRA"
    case KnxManufacturer_M_DKX_TECH:
        return "M_DKX_TECH"
    case KnxManufacturer_M_VIATRON:
        return "M_VIATRON"
    case KnxManufacturer_M_NAUTIBUS:
        return "M_NAUTIBUS"
    case KnxManufacturer_M_ON_SEMICONDUCTOR:
        return "M_ON_SEMICONDUCTOR"
    case KnxManufacturer_M_LONGCHUANG:
        return "M_LONGCHUANG"
    case KnxManufacturer_M_AIR_ON_AG:
        return "M_AIR_ON_AG"
    case KnxManufacturer_M_IB_COMPANY_GMBH:
        return "M_IB_COMPANY_GMBH"
    case KnxManufacturer_M_SATION_FACTORY:
        return "M_SATION_FACTORY"
    case KnxManufacturer_M_AGENTILO_GMBH:
        return "M_AGENTILO_GMBH"
    case KnxManufacturer_M_CRABTREE:
        return "M_CRABTREE"
    case KnxManufacturer_M_MAKEL_ELEKTRIK:
        return "M_MAKEL_ELEKTRIK"
    case KnxManufacturer_M_HELIOS_VENTILATOREN:
        return "M_HELIOS_VENTILATOREN"
    case KnxManufacturer_M_OTTO_SOLUTIONS_PTE_LTD:
        return "M_OTTO_SOLUTIONS_PTE_LTD"
    case KnxManufacturer_M_AIRMASTER:
        return "M_AIRMASTER"
    case KnxManufacturer_M_VALLOX_GMBH:
        return "M_VALLOX_GMBH"
    case KnxManufacturer_M_DALITEK:
        return "M_DALITEK"
    case KnxManufacturer_M_ASIN:
        return "M_ASIN"
    case KnxManufacturer_M_BRIDGES_INTELLIGENCE_TECHNOLOGY_INC_:
        return "M_BRIDGES_INTELLIGENCE_TECHNOLOGY_INC_"
    case KnxManufacturer_M_ARBONIA:
        return "M_ARBONIA"
    case KnxManufacturer_M_KERMI:
        return "M_KERMI"
    case KnxManufacturer_M_EVOKNX:
        return "M_EVOKNX"
    case KnxManufacturer_M_PROLUX:
        return "M_PROLUX"
    case KnxManufacturer_M_CLICHOME:
        return "M_CLICHOME"
    case KnxManufacturer_M_COMMAX:
        return "M_COMMAX"
    case KnxManufacturer_M_EAE:
        return "M_EAE"
    case KnxManufacturer_M_TENSE:
        return "M_TENSE"
    case KnxManufacturer_M_SEYOUNG_ELECTRONICS:
        return "M_SEYOUNG_ELECTRONICS"
    case KnxManufacturer_M_LIFEDOMUS:
        return "M_LIFEDOMUS"
    case KnxManufacturer_M_EUROTRONIC_TECHNOLOGY_GMBH:
        return "M_EUROTRONIC_TECHNOLOGY_GMBH"
    case KnxManufacturer_M_TCI:
        return "M_TCI"
    case KnxManufacturer_M_RISHUN_ELECTRONIC:
        return "M_RISHUN_ELECTRONIC"
    case KnxManufacturer_M_PAUL_HOCHKOEPPER:
        return "M_PAUL_HOCHKOEPPER"
    case KnxManufacturer_M_ZIPATO:
        return "M_ZIPATO"
    case KnxManufacturer_M_CM_SECURITY_GMBH_AND_CO_KG:
        return "M_CM_SECURITY_GMBH_AND_CO_KG"
    case KnxManufacturer_M_QING_CABLES:
        return "M_QING_CABLES"
    case KnxManufacturer_M_LABIO:
        return "M_LABIO"
    case KnxManufacturer_M_COSTER_TECNOLOGIE_ELETTRONICHE_S_P_A_:
        return "M_COSTER_TECNOLOGIE_ELETTRONICHE_S_P_A_"
    case KnxManufacturer_M_E_G_E:
        return "M_E_G_E"
    case KnxManufacturer_M_NETXAUTOMATION:
        return "M_NETXAUTOMATION"
    case KnxManufacturer_M_TECALOR:
        return "M_TECALOR"
    case KnxManufacturer_M_URMET_ELECTRONICS_HUIZHOU_LTD_:
        return "M_URMET_ELECTRONICS_HUIZHOU_LTD_"
    case KnxManufacturer_M_PEIYING_BUILDING_CONTROL:
        return "M_PEIYING_BUILDING_CONTROL"
    case KnxManufacturer_M_ALTENBURGER_ELECTRONIC:
        return "M_ALTENBURGER_ELECTRONIC"
    case KnxManufacturer_M_BPT_S_P_A__A_SOCIO_UNICO:
        return "M_BPT_S_P_A__A_SOCIO_UNICO"
    case KnxManufacturer_M_KANONTEC___KANONBUS:
        return "M_KANONTEC___KANONBUS"
    case KnxManufacturer_M_ISER_TECH:
        return "M_ISER_TECH"
    case KnxManufacturer_M_FINELINE:
        return "M_FINELINE"
    case KnxManufacturer_M_CP_ELECTRONICS_LTD:
        return "M_CP_ELECTRONICS_LTD"
    case KnxManufacturer_M_NIKO_SERVODAN_AS:
        return "M_NIKO_SERVODAN_AS"
    case KnxManufacturer_M_SIMON_309:
        return "M_SIMON_309"
    case KnxManufacturer_M_GM_MODULAR_PVT__LTD_:
        return "M_GM_MODULAR_PVT__LTD_"
    case KnxManufacturer_M_FU_CHENG_INTELLIGENCE:
        return "M_FU_CHENG_INTELLIGENCE"
    case KnxManufacturer_M_NEXKON:
        return "M_NEXKON"
    case KnxManufacturer_M_GRAESSLIN:
        return "M_GRAESSLIN"
    case KnxManufacturer_M_FEEL_S_R_L:
        return "M_FEEL_S_R_L"
    case KnxManufacturer_M_NOT_ASSIGNED_314:
        return "M_NOT_ASSIGNED_314"
    case KnxManufacturer_M_SHENZHEN_FANHAI_SANJIANG_ELECTRONICS_CO___LTD_:
        return "M_SHENZHEN_FANHAI_SANJIANG_ELECTRONICS_CO___LTD_"
    case KnxManufacturer_M_JIUZHOU_GREEBLE:
        return "M_JIUZHOU_GREEBLE"
    case KnxManufacturer_M_AUMUELLER_AUMATIC_GMBH:
        return "M_AUMUELLER_AUMATIC_GMBH"
    case KnxManufacturer_M_ETMAN_ELECTRIC:
        return "M_ETMAN_ELECTRIC"
    case KnxManufacturer_M_BLACK_NOVA:
        return "M_BLACK_NOVA"
    case KnxManufacturer_M_ZIDATECH_AG:
        return "M_ZIDATECH_AG"
    case KnxManufacturer_M_IDGS_BVBA:
        return "M_IDGS_BVBA"
    case KnxManufacturer_M_DAKANIMO:
        return "M_DAKANIMO"
    case KnxManufacturer_M_SIMON_42:
        return "M_SIMON_42"
    case KnxManufacturer_M_TREBOR_AUTOMATION_AB:
        return "M_TREBOR_AUTOMATION_AB"
    case KnxManufacturer_M_SATEL_SP__Z_O_O_:
        return "M_SATEL_SP__Z_O_O_"
    case KnxManufacturer_M_RUSSOUND__INC_:
        return "M_RUSSOUND__INC_"
    case KnxManufacturer_M_MIDEA_HEATING_AND_VENTILATING_EQUIPMENT_CO_LTD:
        return "M_MIDEA_HEATING_AND_VENTILATING_EQUIPMENT_CO_LTD"
    case KnxManufacturer_M_CONSORZIO_TERRANUOVA:
        return "M_CONSORZIO_TERRANUOVA"
    case KnxManufacturer_M_WOLF_HEIZTECHNIK_GMBH:
        return "M_WOLF_HEIZTECHNIK_GMBH"
    case KnxManufacturer_M_SONTEC:
        return "M_SONTEC"
    case KnxManufacturer_M_BELCOM_CABLES_LTD_:
        return "M_BELCOM_CABLES_LTD_"
    case KnxManufacturer_M_GUANGZHOU_SEAWIN_ELECTRICAL_TECHNOLOGIES_CO___LTD_:
        return "M_GUANGZHOU_SEAWIN_ELECTRICAL_TECHNOLOGIES_CO___LTD_"
    case KnxManufacturer_M_ACREL:
        return "M_ACREL"
    case KnxManufacturer_M_VIMAR:
        return "M_VIMAR"
    case KnxManufacturer_M_FRANKE_AQUAROTTER_GMBH:
        return "M_FRANKE_AQUAROTTER_GMBH"
    case KnxManufacturer_M_ORION_SYSTEMS:
        return "M_ORION_SYSTEMS"
    case KnxManufacturer_M_SCHRACK_TECHNIK_GMBH:
        return "M_SCHRACK_TECHNIK_GMBH"
    case KnxManufacturer_M_INSPRID:
        return "M_INSPRID"
    case KnxManufacturer_M_SUNRICHER:
        return "M_SUNRICHER"
    case KnxManufacturer_M_MENRED_AUTOMATION_SYSTEMSHANGHAI_CO__LTD_:
        return "M_MENRED_AUTOMATION_SYSTEMSHANGHAI_CO__LTD_"
    case KnxManufacturer_M_AUREX:
        return "M_AUREX"
    case KnxManufacturer_M_JOSEF_BARTHELME_GMBH_AND_CO__KG:
        return "M_JOSEF_BARTHELME_GMBH_AND_CO__KG"
    case KnxManufacturer_M_ARCHITECTURE_NUMERIQUE:
        return "M_ARCHITECTURE_NUMERIQUE"
    case KnxManufacturer_M_UP_GROUP:
        return "M_UP_GROUP"
    case KnxManufacturer_M_MOELLER_GEBAEUDEAUTOMATION_KG:
        return "M_MOELLER_GEBAEUDEAUTOMATION_KG"
    case KnxManufacturer_M_TEKNOS_AVINNO:
        return "M_TEKNOS_AVINNO"
    case KnxManufacturer_M_NINGBO_DOOYA_MECHANIC_AND_ELECTRONIC_TECHNOLOGY:
        return "M_NINGBO_DOOYA_MECHANIC_AND_ELECTRONIC_TECHNOLOGY"
    case KnxManufacturer_M_THERMOKON_SENSORTECHNIK_GMBH:
        return "M_THERMOKON_SENSORTECHNIK_GMBH"
    case KnxManufacturer_M_BELIMO_AUTOMATION_AG:
        return "M_BELIMO_AUTOMATION_AG"
    case KnxManufacturer_M_ZEHNDER_GROUP_INTERNATIONAL_AG:
        return "M_ZEHNDER_GROUP_INTERNATIONAL_AG"
    case KnxManufacturer_M_SKS_KINKEL_ELEKTRONIK:
        return "M_SKS_KINKEL_ELEKTRONIK"
    case KnxManufacturer_M_ECE_WURMITZER_GMBH:
        return "M_ECE_WURMITZER_GMBH"
    case KnxManufacturer_M_LARS:
        return "M_LARS"
    case KnxManufacturer_M_URC:
        return "M_URC"
    case KnxManufacturer_M_LIGHTCONTROL:
        return "M_LIGHTCONTROL"
    case KnxManufacturer_M_ALBRECHT_JUNG:
        return "M_ALBRECHT_JUNG"
    case KnxManufacturer_M_ELTAKO:
        return "M_ELTAKO"
    case KnxManufacturer_M_SHENZHEN_YM:
        return "M_SHENZHEN_YM"
    case KnxManufacturer_M_MEAN_WELL_ENTERPRISES_CO__LTD_:
        return "M_MEAN_WELL_ENTERPRISES_CO__LTD_"
    case KnxManufacturer_M_OSIX:
        return "M_OSIX"
    case KnxManufacturer_M_AYPRO_TECHNOLOGY:
        return "M_AYPRO_TECHNOLOGY"
    case KnxManufacturer_M_HEFEI_ECOLITE_SOFTWARE:
        return "M_HEFEI_ECOLITE_SOFTWARE"
    case KnxManufacturer_M_ENNO:
        return "M_ENNO"
    case KnxManufacturer_M_OHOSURE:
        return "M_OHOSURE"
    case KnxManufacturer_M_GAREFOWL:
        return "M_GAREFOWL"
    case KnxManufacturer_M_GEZE:
        return "M_GEZE"
    case KnxManufacturer_M_LG_ELECTRONICS_INC_:
        return "M_LG_ELECTRONICS_INC_"
    case KnxManufacturer_M_BOSCH_SIEMENS_HAUSHALTSGERAETE:
        return "M_BOSCH_SIEMENS_HAUSHALTSGERAETE"
    case KnxManufacturer_M_SMC_INTERIORS:
        return "M_SMC_INTERIORS"
    case KnxManufacturer_M_NOT_ASSIGNED_364:
        return "M_NOT_ASSIGNED_364"
    case KnxManufacturer_M_SCS_CABLE:
        return "M_SCS_CABLE"
    case KnxManufacturer_M_HOVAL:
        return "M_HOVAL"
    case KnxManufacturer_M_CANST:
        return "M_CANST"
    case KnxManufacturer_M_HANGZHOU_BERLIN:
        return "M_HANGZHOU_BERLIN"
    case KnxManufacturer_M_EVN_LICHTTECHNIK:
        return "M_EVN_LICHTTECHNIK"
    case KnxManufacturer_M_RUTEC:
        return "M_RUTEC"
    case KnxManufacturer_M_FINDER:
        return "M_FINDER"
    case KnxManufacturer_M_FUJITSU_GENERAL_LIMITED:
        return "M_FUJITSU_GENERAL_LIMITED"
    case KnxManufacturer_M_RITTO_GMBHANDCO_KG:
        return "M_RITTO_GMBHANDCO_KG"
    case KnxManufacturer_M_ZF_FRIEDRICHSHAFEN_AG:
        return "M_ZF_FRIEDRICHSHAFEN_AG"
    case KnxManufacturer_M_CREALED:
        return "M_CREALED"
    case KnxManufacturer_M_MILES_MAGIC_AUTOMATION_PRIVATE_LIMITED:
        return "M_MILES_MAGIC_AUTOMATION_PRIVATE_LIMITED"
    case KnxManufacturer_M_EPlus:
        return "M_EPlus"
    case KnxManufacturer_M_ITALCOND:
        return "M_ITALCOND"
    case KnxManufacturer_M_SATION:
        return "M_SATION"
    case KnxManufacturer_M_NEWBEST:
        return "M_NEWBEST"
    case KnxManufacturer_M_GDS_DIGITAL_SYSTEMS:
        return "M_GDS_DIGITAL_SYSTEMS"
    case KnxManufacturer_M_IDDERO:
        return "M_IDDERO"
    case KnxManufacturer_M_MBNLED:
        return "M_MBNLED"
    case KnxManufacturer_M_POWER_CONTROLS:
        return "M_POWER_CONTROLS"
    case KnxManufacturer_M_VITRUM:
        return "M_VITRUM"
    case KnxManufacturer_M_EKEY_BIOMETRIC_SYSTEMS_GMBH:
        return "M_EKEY_BIOMETRIC_SYSTEMS_GMBH"
    case KnxManufacturer_M_AMC:
        return "M_AMC"
    case KnxManufacturer_M_TRILUX_GMBH_AND_CO__KG:
        return "M_TRILUX_GMBH_AND_CO__KG"
    case KnxManufacturer_M_WEXCEDO:
        return "M_WEXCEDO"
    case KnxManufacturer_M_VEMER_SPA:
        return "M_VEMER_SPA"
    case KnxManufacturer_M_ALEXANDER_BUERKLE_GMBH_AND_CO_KG:
        return "M_ALEXANDER_BUERKLE_GMBH_AND_CO_KG"
    case KnxManufacturer_M_CITRON:
        return "M_CITRON"
    case KnxManufacturer_M_SHENZHEN_HEGUANG:
        return "M_SHENZHEN_HEGUANG"
    case KnxManufacturer_M_NOT_ASSIGNED_392:
        return "M_NOT_ASSIGNED_392"
    case KnxManufacturer_M_ZUMTOBEL:
        return "M_ZUMTOBEL"
    case KnxManufacturer_M_TRANE_B_V_B_A:
        return "M_TRANE_B_V_B_A"
    case KnxManufacturer_M_CAREL:
        return "M_CAREL"
    case KnxManufacturer_M_PROLITE_CONTROLS:
        return "M_PROLITE_CONTROLS"
    case KnxManufacturer_M_BOSMER:
        return "M_BOSMER"
    case KnxManufacturer_M_EUCHIPS:
        return "M_EUCHIPS"
    case KnxManufacturer_M_CONNECT_THINKA_CONNECT:
        return "M_CONNECT_THINKA_CONNECT"
    case KnxManufacturer_M_PEAKNX_A_DOGAWIST_COMPANY:
        return "M_PEAKNX_A_DOGAWIST_COMPANY"
    case KnxManufacturer_M_ACEMATIC:
        return "M_ACEMATIC"
    case KnxManufacturer_M_ELAUSYS:
        return "M_ELAUSYS"
    case KnxManufacturer_M_ITK_ENGINEERING_AG:
        return "M_ITK_ENGINEERING_AG"
    case KnxManufacturer_M_PHOENIX_CONTACT:
        return "M_PHOENIX_CONTACT"
    case KnxManufacturer_M_INTEGRA_METERING_AG:
        return "M_INTEGRA_METERING_AG"
    case KnxManufacturer_M_FMS_HOSPITALITY_PTE_LTD:
        return "M_FMS_HOSPITALITY_PTE_LTD"
    case KnxManufacturer_M_NUVO:
        return "M_NUVO"
    case KnxManufacturer_M_U__LUX_GMBH:
        return "M_U__LUX_GMBH"
    case KnxManufacturer_M_BRUMBERG_LEUCHTEN:
        return "M_BRUMBERG_LEUCHTEN"
    case KnxManufacturer_M_LIME:
        return "M_LIME"
    case KnxManufacturer_M_GREAT_EMPIRE_INTERNATIONAL_GROUP_CO___LTD_:
        return "M_GREAT_EMPIRE_INTERNATIONAL_GROUP_CO___LTD_"
    case KnxManufacturer_M_KAVOSHPISHRO_ASIA:
        return "M_KAVOSHPISHRO_ASIA"
    case KnxManufacturer_M_V2_SPA:
        return "M_V2_SPA"
    case KnxManufacturer_M_JOHNSON_CONTROLS:
        return "M_JOHNSON_CONTROLS"
    case KnxManufacturer_M_WAGO_KONTAKTTECHNIK:
        return "M_WAGO_KONTAKTTECHNIK"
    case KnxManufacturer_M_ARKUD:
        return "M_ARKUD"
    case KnxManufacturer_M_IRIDIUM_LTD_:
        return "M_IRIDIUM_LTD_"
    case KnxManufacturer_M_BSMART:
        return "M_BSMART"
    case KnxManufacturer_M_BAB_TECHNOLOGIE_GMBH:
        return "M_BAB_TECHNOLOGIE_GMBH"
    case KnxManufacturer_M_NICE_SPA:
        return "M_NICE_SPA"
    case KnxManufacturer_M_REDFISH_GROUP_PTY_LTD:
        return "M_REDFISH_GROUP_PTY_LTD"
    case KnxManufacturer_M_SABIANA_SPA:
        return "M_SABIANA_SPA"
    case KnxManufacturer_M_UBEE_INTERACTIVE_EUROPE:
        return "M_UBEE_INTERACTIVE_EUROPE"
    case KnxManufacturer_M_REXEL:
        return "M_REXEL"
    case KnxManufacturer_M_GES_TEKNIK_A_S_:
        return "M_GES_TEKNIK_A_S_"
    case KnxManufacturer_M_KNXPRESSO:
        return "M_KNXPRESSO"
    case KnxManufacturer_M_AVE_S_P_A_:
        return "M_AVE_S_P_A_"
    case KnxManufacturer_M_ZHUHAI_LTECH_TECHNOLOGY_CO___LTD_:
        return "M_ZHUHAI_LTECH_TECHNOLOGY_CO___LTD_"
    case KnxManufacturer_M_ARCOM:
        return "M_ARCOM"
    case KnxManufacturer_M_VIA_TECHNOLOGIES__INC_:
        return "M_VIA_TECHNOLOGIES__INC_"
    case KnxManufacturer_M_FEELSMART_:
        return "M_FEELSMART_"
    case KnxManufacturer_M_SUPCON:
        return "M_SUPCON"
    case KnxManufacturer_M_MANIC:
        return "M_MANIC"
    case KnxManufacturer_M_TDE_GMBH:
        return "M_TDE_GMBH"
    case KnxManufacturer_M_NANJING_SHUFAN_INFORMATION_TECHNOLOGY_CO__LTD_:
        return "M_NANJING_SHUFAN_INFORMATION_TECHNOLOGY_CO__LTD_"
    case KnxManufacturer_M_EWTECH:
        return "M_EWTECH"
    case KnxManufacturer_M_WIELAND_ELECTRIC:
        return "M_WIELAND_ELECTRIC"
    case KnxManufacturer_M_KLUGER_AUTOMATION_GMBH:
        return "M_KLUGER_AUTOMATION_GMBH"
    case KnxManufacturer_M_JOONGANG_CONTROL:
        return "M_JOONGANG_CONTROL"
    case KnxManufacturer_M_GREENCONTROLS_TECHNOLOGY_SDN__BHD_:
        return "M_GREENCONTROLS_TECHNOLOGY_SDN__BHD_"
    case KnxManufacturer_M_IME_S_P_A_:
        return "M_IME_S_P_A_"
    case KnxManufacturer_M_SICHUAN_HAODING:
        return "M_SICHUAN_HAODING"
    case KnxManufacturer_M_MINDJAGA_LTD_:
        return "M_MINDJAGA_LTD_"
    case KnxManufacturer_M_RUILI_SMART_CONTROL:
        return "M_RUILI_SMART_CONTROL"
    case KnxManufacturer_M_CODESYS_GMBH:
        return "M_CODESYS_GMBH"
    case KnxManufacturer_M_MOORGEN_DEUTSCHLAND_GMBH:
        return "M_MOORGEN_DEUTSCHLAND_GMBH"
    case KnxManufacturer_M_CULLMANN_TECH:
        return "M_CULLMANN_TECH"
    case KnxManufacturer_M_HERMANN_KLEINHUIS:
        return "M_HERMANN_KLEINHUIS"
    case KnxManufacturer_M_MERCK_WINDOW_TECHNOLOGIES_B_V_:
        return "M_MERCK_WINDOW_TECHNOLOGIES_B_V_"
    case KnxManufacturer_M_ABEGO:
        return "M_ABEGO"
    case KnxManufacturer_M_MYGEKKO:
        return "M_MYGEKKO"
    case KnxManufacturer_M_ERGO3_SARL:
        return "M_ERGO3_SARL"
    case KnxManufacturer_M_STMICROELECTRONICS_INTERNATIONAL_N_V_:
        return "M_STMICROELECTRONICS_INTERNATIONAL_N_V_"
    case KnxManufacturer_M_CJC_SYSTEMS:
        return "M_CJC_SYSTEMS"
    case KnxManufacturer_M_SUDOKU:
        return "M_SUDOKU"
    case KnxManufacturer_M_AZ_E_LITE_PTE_LTD:
        return "M_AZ_E_LITE_PTE_LTD"
    case KnxManufacturer_M_ARLIGHT:
        return "M_ARLIGHT"
    case KnxManufacturer_M_GRUENBECK_WASSERAUFBEREITUNG_GMBH:
        return "M_GRUENBECK_WASSERAUFBEREITUNG_GMBH"
    case KnxManufacturer_M_BTICINO:
        return "M_BTICINO"
    case KnxManufacturer_M_STIEBEL_ELTRON:
        return "M_STIEBEL_ELTRON"
    case KnxManufacturer_M_MODULE_ELECTRONIC:
        return "M_MODULE_ELECTRONIC"
    case KnxManufacturer_M_KOPLAT:
        return "M_KOPLAT"
    case KnxManufacturer_M_GUANGZHOU_LETOUR_LIFE_TECHNOLOGY_CO___LTD:
        return "M_GUANGZHOU_LETOUR_LIFE_TECHNOLOGY_CO___LTD"
    case KnxManufacturer_M_ILEVIA:
        return "M_ILEVIA"
    case KnxManufacturer_M_LN_SYSTEMTEQ:
        return "M_LN_SYSTEMTEQ"
    case KnxManufacturer_M_HISENSE_SMARTHOME:
        return "M_HISENSE_SMARTHOME"
    case KnxManufacturer_M_FLINK_AUTOMATION_SYSTEM:
        return "M_FLINK_AUTOMATION_SYSTEM"
    case KnxManufacturer_M_XXTER_BV:
        return "M_XXTER_BV"
    case KnxManufacturer_M_LYNXUS_TECHNOLOGY:
        return "M_LYNXUS_TECHNOLOGY"
    case KnxManufacturer_M_ROBOT_S_A_:
        return "M_ROBOT_S_A_"
    case KnxManufacturer_M_TEHALIT:
        return "M_TEHALIT"
    case KnxManufacturer_M_SHENZHEN_ATTE_SMART_LIFE_CO__LTD_:
        return "M_SHENZHEN_ATTE_SMART_LIFE_CO__LTD_"
    case KnxManufacturer_M_NOBLESSE:
        return "M_NOBLESSE"
    case KnxManufacturer_M_ADVANCED_DEVICES:
        return "M_ADVANCED_DEVICES"
    case KnxManufacturer_M_ATRINA_BUILDING_AUTOMATION_CO__LTD:
        return "M_ATRINA_BUILDING_AUTOMATION_CO__LTD"
    case KnxManufacturer_M_GUANGDONG_DAMING_LAFFEY_ELECTRIC_CO___LTD_:
        return "M_GUANGDONG_DAMING_LAFFEY_ELECTRIC_CO___LTD_"
    case KnxManufacturer_M_WESTERSTRAND_URFABRIK_AB:
        return "M_WESTERSTRAND_URFABRIK_AB"
    case KnxManufacturer_M_CONTROL4_CORPORATE:
        return "M_CONTROL4_CORPORATE"
    case KnxManufacturer_M_ONTROL:
        return "M_ONTROL"
    case KnxManufacturer_M_STARNET:
        return "M_STARNET"
    case KnxManufacturer_M_BETA_CAVI:
        return "M_BETA_CAVI"
    case KnxManufacturer_M_THEBEN_AG:
        return "M_THEBEN_AG"
    case KnxManufacturer_M_EASEMORE:
        return "M_EASEMORE"
    case KnxManufacturer_M_VIVALDI_SRL:
        return "M_VIVALDI_SRL"
    case KnxManufacturer_M_GREE_ELECTRIC_APPLIANCES_INC__OF_ZHUHAI:
        return "M_GREE_ELECTRIC_APPLIANCES_INC__OF_ZHUHAI"
    case KnxManufacturer_M_HWISCON:
        return "M_HWISCON"
    case KnxManufacturer_M_SHANGHAI_ELECON_INTELLIGENT_TECHNOLOGY_CO___LTD_:
        return "M_SHANGHAI_ELECON_INTELLIGENT_TECHNOLOGY_CO___LTD_"
    case KnxManufacturer_M_KAMPMANN:
        return "M_KAMPMANN"
    case KnxManufacturer_M_IMPOLUX_GMBH_LEDIMAX:
        return "M_IMPOLUX_GMBH_LEDIMAX"
    case KnxManufacturer_M_EVAUX:
        return "M_EVAUX"
    case KnxManufacturer_M_WEBRO_CABLES_AND_CONNECTORS_LIMITED:
        return "M_WEBRO_CABLES_AND_CONNECTORS_LIMITED"
    case KnxManufacturer_M_SHANGHAI_E_TECH_SOLUTION:
        return "M_SHANGHAI_E_TECH_SOLUTION"
    case KnxManufacturer_M_WILHELM_RUTENBECK:
        return "M_WILHELM_RUTENBECK"
    case KnxManufacturer_M_GUANGZHOU_HOKO_ELECTRIC_CO__LTD_:
        return "M_GUANGZHOU_HOKO_ELECTRIC_CO__LTD_"
    case KnxManufacturer_M_LAMMIN_HIGH_TECH_CO__LTD:
        return "M_LAMMIN_HIGH_TECH_CO__LTD"
    case KnxManufacturer_M_SHENZHEN_MERRYTEK_TECHNOLOGY_CO___LTD:
        return "M_SHENZHEN_MERRYTEK_TECHNOLOGY_CO___LTD"
    case KnxManufacturer_M_I_LUXUS:
        return "M_I_LUXUS"
    case KnxManufacturer_M_ELMOS_SEMICONDUCTOR_AG:
        return "M_ELMOS_SEMICONDUCTOR_AG"
    case KnxManufacturer_M_EMCOM_TECHNOLOGY_INC:
        return "M_EMCOM_TECHNOLOGY_INC"
    case KnxManufacturer_M_PROJECT_INNOVATIONS_GMBH:
        return "M_PROJECT_INNOVATIONS_GMBH"
    case KnxManufacturer_M_ITC:
        return "M_ITC"
    case KnxManufacturer_M_ABB_LV_INSTALLATION_MATERIALS_COMPANY_LTD__BEIJING:
        return "M_ABB_LV_INSTALLATION_MATERIALS_COMPANY_LTD__BEIJING"
    case KnxManufacturer_M_MAICO:
        return "M_MAICO"
    case KnxManufacturer_M_WINKHAUS:
        return "M_WINKHAUS"
    case KnxManufacturer_M_ELAN_SRL:
        return "M_ELAN_SRL"
    case KnxManufacturer_M_MINHHA_TECHNOLOGY_CO__LTD:
        return "M_MINHHA_TECHNOLOGY_CO__LTD"
    case KnxManufacturer_M_ZHEJIANG_TIANJIE_INDUSTRIAL_CORP_:
        return "M_ZHEJIANG_TIANJIE_INDUSTRIAL_CORP_"
    case KnxManufacturer_M_IAUTOMATION_PTY_LIMITED:
        return "M_IAUTOMATION_PTY_LIMITED"
    case KnxManufacturer_M_EXTRON:
        return "M_EXTRON"
    case KnxManufacturer_M_FREEDOMPRO:
        return "M_FREEDOMPRO"
    case KnxManufacturer_M_ONEHOME:
        return "M_ONEHOME"
    case KnxManufacturer_M_EOS_SAUNATECHNIK_GMBH:
        return "M_EOS_SAUNATECHNIK_GMBH"
    case KnxManufacturer_M_KUSATEK_GMBH:
        return "M_KUSATEK_GMBH"
    case KnxManufacturer_M_EISBAER_SCADA:
        return "M_EISBAER_SCADA"
    case KnxManufacturer_M_ROBERT_BOSCH:
        return "M_ROBERT_BOSCH"
    case KnxManufacturer_M_AUTOMATISMI_BENINCA_S_P_A_:
        return "M_AUTOMATISMI_BENINCA_S_P_A_"
    case KnxManufacturer_M_BLENDOM:
        return "M_BLENDOM"
    case KnxManufacturer_M_MADEL_AIR_TECHNICAL_DIFFUSION:
        return "M_MADEL_AIR_TECHNICAL_DIFFUSION"
    case KnxManufacturer_M_NIKO:
        return "M_NIKO"
    case KnxManufacturer_M_BOSCH_REXROTH_AG:
        return "M_BOSCH_REXROTH_AG"
    case KnxManufacturer_M_CANDM_PRODUCTS:
        return "M_CANDM_PRODUCTS"
    case KnxManufacturer_M_HOERMANN_KG_VERKAUFSGESELLSCHAFT:
        return "M_HOERMANN_KG_VERKAUFSGESELLSCHAFT"
    case KnxManufacturer_M_SHANGHAI_RAJAYASA_CO__LTD:
        return "M_SHANGHAI_RAJAYASA_CO__LTD"
    case KnxManufacturer_M_SUZUKI:
        return "M_SUZUKI"
    case KnxManufacturer_M_SILENT_GLISS_INTERNATIONAL_LTD_:
        return "M_SILENT_GLISS_INTERNATIONAL_LTD_"
    case KnxManufacturer_M_SOMFY:
        return "M_SOMFY"
    case KnxManufacturer_M_BEE_CONTROLS_ADGSC_GROUP:
        return "M_BEE_CONTROLS_ADGSC_GROUP"
    case KnxManufacturer_M_XDTECGMBH:
        return "M_XDTECGMBH"
    case KnxManufacturer_M_OSRAM:
        return "M_OSRAM"
    case KnxManufacturer_M_LEBENOR:
        return "M_LEBENOR"
    case KnxManufacturer_M_AUTOMANENG:
        return "M_AUTOMANENG"
    case KnxManufacturer_M_HONEYWELL_AUTOMATION_SOLUTION_CONTROLCHINA:
        return "M_HONEYWELL_AUTOMATION_SOLUTION_CONTROLCHINA"
    case KnxManufacturer_M_HANGZHOU_BINTHEN_INTELLIGENCE_TECHNOLOGY_CO__LTD:
        return "M_HANGZHOU_BINTHEN_INTELLIGENCE_TECHNOLOGY_CO__LTD"
    case KnxManufacturer_M_ETA_HEIZTECHNIK:
        return "M_ETA_HEIZTECHNIK"
    case KnxManufacturer_M_DIVUS_GMBH:
        return "M_DIVUS_GMBH"
    case KnxManufacturer_M_NANJING_TAIJIESAI_INTELLIGENT_TECHNOLOGY_CO__LTD_:
        return "M_NANJING_TAIJIESAI_INTELLIGENT_TECHNOLOGY_CO__LTD_"
    case KnxManufacturer_M_WOERTZ:
        return "M_WOERTZ"
    case KnxManufacturer_M_LUNATONE:
        return "M_LUNATONE"
    case KnxManufacturer_M_ZHEJIANG_SCTECH_BUILDING_INTELLIGENT:
        return "M_ZHEJIANG_SCTECH_BUILDING_INTELLIGENT"
    case KnxManufacturer_M_FOSHAN_QITE_TECHNOLOGY_CO___LTD_:
        return "M_FOSHAN_QITE_TECHNOLOGY_CO___LTD_"
    case KnxManufacturer_M_NOKE:
        return "M_NOKE"
    case KnxManufacturer_M_LANDCOM:
        return "M_LANDCOM"
    case KnxManufacturer_M_STORK_AS:
        return "M_STORK_AS"
    case KnxManufacturer_M_HANGZHOU_SHENDU_TECHNOLOGY_CO___LTD_:
        return "M_HANGZHOU_SHENDU_TECHNOLOGY_CO___LTD_"
    case KnxManufacturer_M_COOLAUTOMATION:
        return "M_COOLAUTOMATION"
    case KnxManufacturer_M_APRSTERN:
        return "M_APRSTERN"
    case KnxManufacturer_M_SONNEN:
        return "M_SONNEN"
    case KnxManufacturer_M_VIESSMANN_WERKE:
        return "M_VIESSMANN_WERKE"
    case KnxManufacturer_M_DNAKE:
        return "M_DNAKE"
    case KnxManufacturer_M_NEUBERGER_GEBAEUDEAUTOMATION_GMBH:
        return "M_NEUBERGER_GEBAEUDEAUTOMATION_GMBH"
    case KnxManufacturer_M_STILIGER:
        return "M_STILIGER"
    case KnxManufacturer_M_BERGHOF_AUTOMATION_GMBH:
        return "M_BERGHOF_AUTOMATION_GMBH"
    case KnxManufacturer_M_TOTAL_AUTOMATION_AND_CONTROLS_GMBH:
        return "M_TOTAL_AUTOMATION_AND_CONTROLS_GMBH"
    case KnxManufacturer_M_DOVIT:
        return "M_DOVIT"
    case KnxManufacturer_M_INSTALIGHTING_GMBH:
        return "M_INSTALIGHTING_GMBH"
    case KnxManufacturer_M_UNI_TEC:
        return "M_UNI_TEC"
    case KnxManufacturer_M_CASATUNES:
        return "M_CASATUNES"
    case KnxManufacturer_M_EMT:
        return "M_EMT"
    case KnxManufacturer_M_IMI_HYDRONIC_ENGINEERING:
        return "M_IMI_HYDRONIC_ENGINEERING"
    case KnxManufacturer_M_SENFFICIENT:
        return "M_SENFFICIENT"
    case KnxManufacturer_M_AUROLITE_ELECTRICAL_PANYU_GUANGZHOU_LIMITED:
        return "M_AUROLITE_ELECTRICAL_PANYU_GUANGZHOU_LIMITED"
    case KnxManufacturer_M_ABB_XIAMEN_SMART_TECHNOLOGY_CO___LTD_:
        return "M_ABB_XIAMEN_SMART_TECHNOLOGY_CO___LTD_"
    case KnxManufacturer_M_SAMSON_ELECTRIC_WIRE:
        return "M_SAMSON_ELECTRIC_WIRE"
    case KnxManufacturer_M_T_TOUCHING:
        return "M_T_TOUCHING"
    case KnxManufacturer_M_CORE_SMART_HOME:
        return "M_CORE_SMART_HOME"
    case KnxManufacturer_M_GREENCONNECT_SOLUTIONS_SA:
        return "M_GREENCONNECT_SOLUTIONS_SA"
    case KnxManufacturer_M_ELETTRONICA_CONDUTTORI:
        return "M_ELETTRONICA_CONDUTTORI"
    case KnxManufacturer_M_MKFC:
        return "M_MKFC"
    case KnxManufacturer_M_AUTOMATIONPlus:
        return "M_AUTOMATIONPlus"
    case KnxManufacturer_M_BERKER:
        return "M_BERKER"
    case KnxManufacturer_M_JOH__VAILLANT:
        return "M_JOH__VAILLANT"
    case KnxManufacturer_M_BLUE_AND_RED:
        return "M_BLUE_AND_RED"
    case KnxManufacturer_M_FROGBLUE:
        return "M_FROGBLUE"
    case KnxManufacturer_M_SAVESOR:
        return "M_SAVESOR"
    case KnxManufacturer_M_APP_TECH:
        return "M_APP_TECH"
    case KnxManufacturer_M_SENSORTEC_AG:
        return "M_SENSORTEC_AG"
    case KnxManufacturer_M_NYSA_TECHNOLOGY_AND_SOLUTIONS:
        return "M_NYSA_TECHNOLOGY_AND_SOLUTIONS"
    case KnxManufacturer_M_FARADITE:
        return "M_FARADITE"
    case KnxManufacturer_M_OPTIMUS:
        return "M_OPTIMUS"
    case KnxManufacturer_M_KTS_S_R_L_:
        return "M_KTS_S_R_L_"
    case KnxManufacturer_M_RAMCRO_SPA:
        return "M_RAMCRO_SPA"
    case KnxManufacturer_M_AMP_DEUTSCHLAND:
        return "M_AMP_DEUTSCHLAND"
    case KnxManufacturer_M_WUHAN_WISECREATE_UNIVERSE_TECHNOLOGY_CO___LTD:
        return "M_WUHAN_WISECREATE_UNIVERSE_TECHNOLOGY_CO___LTD"
    case KnxManufacturer_M_BEMI_SMART_HOME_LTD:
        return "M_BEMI_SMART_HOME_LTD"
    case KnxManufacturer_M_ARDOMUS:
        return "M_ARDOMUS"
    case KnxManufacturer_M_CHANGXING:
        return "M_CHANGXING"
    case KnxManufacturer_M_E_CONTROLS:
        return "M_E_CONTROLS"
    case KnxManufacturer_M_AIB_TECHNOLOGY:
        return "M_AIB_TECHNOLOGY"
    case KnxManufacturer_M_NVC:
        return "M_NVC"
    case KnxManufacturer_M_KBOX:
        return "M_KBOX"
    case KnxManufacturer_M_CNS:
        return "M_CNS"
    case KnxManufacturer_M_TYBA:
        return "M_TYBA"
    case KnxManufacturer_M_BOSCH_THERMOTECHNIK_GMBH:
        return "M_BOSCH_THERMOTECHNIK_GMBH"
    case KnxManufacturer_M_ATREL:
        return "M_ATREL"
    case KnxManufacturer_M_SIMON_ELECTRIC_CHINA_CO___LTD:
        return "M_SIMON_ELECTRIC_CHINA_CO___LTD"
    case KnxManufacturer_M_KORDZ_GROUP:
        return "M_KORDZ_GROUP"
    case KnxManufacturer_M_ND_ELECTRIC:
        return "M_ND_ELECTRIC"
    case KnxManufacturer_M_CONTROLIUM:
        return "M_CONTROLIUM"
    case KnxManufacturer_M_FAMO_GMBH_AND_CO__KG:
        return "M_FAMO_GMBH_AND_CO__KG"
    case KnxManufacturer_M_CDN_SMART:
        return "M_CDN_SMART"
    case KnxManufacturer_M_HESTON:
        return "M_HESTON"
    case KnxManufacturer_M_ESLA_CONEXIONES_S_L_:
        return "M_ESLA_CONEXIONES_S_L_"
    case KnxManufacturer_M_WEISHAUPT:
        return "M_WEISHAUPT"
    case KnxManufacturer_M_SEF___ECOTEC:
        return "M_SEF___ECOTEC"
    case KnxManufacturer_M_ASTRUM_TECHNOLOGY:
        return "M_ASTRUM_TECHNOLOGY"
    case KnxManufacturer_M_WUERTH_ELEKTRONIK_STELVIO_KONTEK_S_P_A_:
        return "M_WUERTH_ELEKTRONIK_STELVIO_KONTEK_S_P_A_"
    case KnxManufacturer_M_NANOTECO_CORPORATION:
        return "M_NANOTECO_CORPORATION"
    case KnxManufacturer_M_NIETIAN:
        return "M_NIETIAN"
    case KnxManufacturer_M_SUMSIR:
        return "M_SUMSIR"
    case KnxManufacturer_M_ORBIS_TECNOLOGIA_ELECTRICA_SA:
        return "M_ORBIS_TECNOLOGIA_ELECTRICA_SA"
    case KnxManufacturer_M_ABB___RESERVED:
        return "M_ABB___RESERVED"
    case KnxManufacturer_M_BUSCH_JAEGER_ELEKTRO___RESERVED:
        return "M_BUSCH_JAEGER_ELEKTRO___RESERVED"
    case KnxManufacturer_M_DORMA_GMBH_Plus_CO__KG:
        return "M_DORMA_GMBH_Plus_CO__KG"
    case KnxManufacturer_M_WINDOWMASTER_AS:
        return "M_WINDOWMASTER_AS"
    case KnxManufacturer_M_WALTHER_WERKE:
        return "M_WALTHER_WERKE"
    case KnxManufacturer_M_ORAS:
        return "M_ORAS"
    case KnxManufacturer_M_DAETWYLER:
        return "M_DAETWYLER"
    case KnxManufacturer_M_ELECTRAK:
        return "M_ELECTRAK"
    case KnxManufacturer_M_BUSCH_JAEGER_ELEKTRO:
        return "M_BUSCH_JAEGER_ELEKTRO"
    case KnxManufacturer_M_TECHEM:
        return "M_TECHEM"
    case KnxManufacturer_M_SCHNEIDER_ELECTRIC_INDUSTRIES_SAS:
        return "M_SCHNEIDER_ELECTRIC_INDUSTRIES_SAS"
    case KnxManufacturer_M_WHD_WILHELM_HUBER_Plus_SOEHNE:
        return "M_WHD_WILHELM_HUBER_Plus_SOEHNE"
    case KnxManufacturer_M_BISCHOFF_ELEKTRONIK:
        return "M_BISCHOFF_ELEKTRONIK"
    case KnxManufacturer_M_JEPAZ:
        return "M_JEPAZ"
    case KnxManufacturer_M_RTS_AUTOMATION:
        return "M_RTS_AUTOMATION"
    case KnxManufacturer_M_EIBMARKT_GMBH:
        return "M_EIBMARKT_GMBH"
    case KnxManufacturer_M_WAREMA_RENKHOFF_SE:
        return "M_WAREMA_RENKHOFF_SE"
    case KnxManufacturer_M_EELECTRON:
        return "M_EELECTRON"
    case KnxManufacturer_M_BELDEN_WIRE_AND_CABLE_B_V_:
        return "M_BELDEN_WIRE_AND_CABLE_B_V_"
    case KnxManufacturer_M_GIRA_GIERSIEPEN:
        return "M_GIRA_GIERSIEPEN"
    case KnxManufacturer_M_BECKER_ANTRIEBE_GMBH:
        return "M_BECKER_ANTRIEBE_GMBH"
    case KnxManufacturer_M_J_STEHLEPlusSOEHNE_GMBH:
        return "M_J_STEHLEPlusSOEHNE_GMBH"
    case KnxManufacturer_M_AGFEO:
        return "M_AGFEO"
    case KnxManufacturer_M_ZENNIO:
        return "M_ZENNIO"
    case KnxManufacturer_M_TAPKO_TECHNOLOGIES:
        return "M_TAPKO_TECHNOLOGIES"
    case KnxManufacturer_M_HDL:
        return "M_HDL"
    case KnxManufacturer_M_UPONOR:
        return "M_UPONOR"
    case KnxManufacturer_M_SE_LIGHTMANAGEMENT_AG:
        return "M_SE_LIGHTMANAGEMENT_AG"
    case KnxManufacturer_M_ARCUS_EDS:
        return "M_ARCUS_EDS"
    case KnxManufacturer_M_INTESIS:
        return "M_INTESIS"
    case KnxManufacturer_M_HAGER_ELECTRO:
        return "M_HAGER_ELECTRO"
    case KnxManufacturer_M_HERHOLDT_CONTROLS_SRL:
        return "M_HERHOLDT_CONTROLS_SRL"
    case KnxManufacturer_M_NIKO_ZUBLIN:
        return "M_NIKO_ZUBLIN"
    case KnxManufacturer_M_DURABLE_TECHNOLOGIES:
        return "M_DURABLE_TECHNOLOGIES"
    case KnxManufacturer_M_INNOTEAM:
        return "M_INNOTEAM"
    case KnxManufacturer_M_ISE_GMBH:
        return "M_ISE_GMBH"
    case KnxManufacturer_M_TEAM_FOR_TRONICS:
        return "M_TEAM_FOR_TRONICS"
    case KnxManufacturer_M_CIAT:
        return "M_CIAT"
    case KnxManufacturer_M_REMEHA_BV:
        return "M_REMEHA_BV"
    case KnxManufacturer_M_ESYLUX:
        return "M_ESYLUX"
    case KnxManufacturer_M_BASALTE:
        return "M_BASALTE"
    case KnxManufacturer_M_INSTA_GMBH:
        return "M_INSTA_GMBH"
    case KnxManufacturer_M_VESTAMATIC:
        return "M_VESTAMATIC"
    case KnxManufacturer_M_MDT_TECHNOLOGIES:
        return "M_MDT_TECHNOLOGIES"
    case KnxManufacturer_M_WARENDORFER_KUECHEN_GMBH:
        return "M_WARENDORFER_KUECHEN_GMBH"
    case KnxManufacturer_M_VIDEO_STAR:
        return "M_VIDEO_STAR"
    case KnxManufacturer_M_SITEK:
        return "M_SITEK"
    case KnxManufacturer_M_CONTROLTRONIC:
        return "M_CONTROLTRONIC"
    case KnxManufacturer_M_FUNCTION_TECHNOLOGY:
        return "M_FUNCTION_TECHNOLOGY"
    case KnxManufacturer_M_AMX:
        return "M_AMX"
    case KnxManufacturer_M_ELDAT:
        return "M_ELDAT"
    case KnxManufacturer_M_PANASONIC:
        return "M_PANASONIC"
    }
    return ""
}
