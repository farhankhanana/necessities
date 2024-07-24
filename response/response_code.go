package response

// HTTP response code
const (
	HTTPOk                   = 200
	HTTPCreated              = 201
	HTTPAccepted             = 202
	HTTPNoContent            = 204
	HTTPBadRequest           = 400
	HTTPUnauthorized         = 401
	HTTPForbidden            = 403
	HTTPNotFound             = 404
	HTTPMethodNotAllowed     = 405
	HTTPNotAcceptable        = 406
	HTTPRequestTimeout       = 408
	HTTPConflict             = 409
	HTTPUnsupportedMediaType = 415
	HTTPInternalServerError  = 500
	HTTPNotImplemented       = 501
	HTTPServiceUnavailable   = 503
)

// Internal response Code
const (
	RCSuccess                      = "0000"
	RCNoDataUpdated                = "0001"
	RCUnknownError                 = "1001"
	RCSystemError                  = "1002"
	RCDatabaseError                = "1003"
	RCFileSystemError              = "1004"
	RCThirdPartySystemError        = "1005"
	RCConnectionTimeout            = "1006"
	RCDataNotFound                 = "1007"
	RCDuplicateData                = "1008"
	RCImmutableData                = "1009"
	RCNotAuthorizedAccess          = "1010"
	RCIInvalidCredential           = "1011"
	RCUserIsLoggedIn               = "1012"
	RCInvalidLoginSession          = "1013"
	RCUnsupportedBodyType          = "1014"
	RCMissingParameter             = "1015"
	RCInvalidInputFormat           = "1016"
	RCUploadFileFailed             = "1017"
	RCOTPHasBeenSent               = "1018"
	RCOTPHasExpired                = "1019"
	RCOTPInvalid                   = "1020"
	RCPINHasExpired                = "1021"
	RCPINInvalid                   = "1022"
	RCAccountNotFullySetup         = "1023"
	RCAccountDisabled              = "1024"
	RCBadRequest                   = "1025"
	RCInvalidSignature             = "1026"
	RCOTPAttemptBlocked            = "1027"
	RCOTPMaxAttempt                = "1028"
	RCInvalidCode                  = "1029"
	RCCodeHasExpired               = "1030"
	RCInvalidInputData             = "1031"
	RCValidationProcessNotComplete = "1032"
	RCUnregisteredDeviceID         = "1033"
	RCCommunicateWithIbridge       = "1034"
	RCOTPAttemptBlockedPermanent   = "1035"
	RCOTPAttemptBlockedByNumber    = "1036"
	RCInternalServerError          = "1037"
	RCAuthMaxAttempt               = "1038"
	RCAuthAttemptBlockedPermanent  = "1039"
	RCOTPMaxSend                   = "1059"
	RCOTPMaxResend                 = "1060"
	RCOTPMaxVerify                 = "1061"
	RCRedisConnection              = "1062"
	RCErrorGateway                 = "1040"
	RCCardRegisteredDifference     = "1041"
	RCUserBlacklisted              = "1042"
	RCInvalidPersonalData          = "1043"
	RCNeedManualVerification       = "1044"
	RCFaceUndetected               = "1045"
	RCTransactionMaxDailyLimit     = "1046"
	RCTransactionMaxLimit          = "1047"
	RCInvalidExpiredCard           = "1048"
	RCInvalidATMCardNumber         = "1049"
	RCDataOnProcess                = "1050"
	RCServiceNotFound              = "1051"
	RCInvalidIPassport             = "1052"
	RCFaceDoesntMatch              = "1053"
	RCCantExecuteTask              = "1072"
	RCUseOriginalKTP               = "1056"
	RCCantRegisterMB               = "1054"
	RCIDNumberAlreadyRegistered    = "1057"
	RCImageCantBeRead              = "1058"
	RCIdentityNotMatchKTP          = "1055"
	RCConfirmationLinkSent         = "1063"
	RCInsufficientBalance          = "1064"
	RCTransactionNotAllowed        = "1065"
	RCTransactionBelowMinimal      = "1066"
	RCTransactionAmountDifferent   = "1067"
)
