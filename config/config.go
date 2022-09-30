package config

type Config struct {
	Username                string   // `optional will prompt`
	Passwd                  string   // `optional will prompt`
	SplunkUrl               string   // `required max length 64 characters`
	SleepTimeBetweenQueries int      // `required (ms between queries)`
	TimeFromNowToQuery      string   // `required could be in m,h,d from now backwards (##hours,##minutes,##days)`
	Handlers                []string // `required at least 1 Maximum characters 50`
	SplunkBaseUrl           string   // `required at least 1 Maximum characters 50`
	SplunkLoginUrl          string   // `required at least 1 Maximum characters 50`
	SplunkQueryUrl          string   // `required at least 1 Maximum characters 50`
	Data                    string   // `required at least 1 Maximum characters 50`
}

func CreateConfig() (c Config, err error) {
	c = Config{
		Username:                "",
		Passwd:                  "",
		SplunkBaseUrl:           "https://applepay.splunk.pie.apple.com:8089",
		SplunkLoginUrl:          "/services/auth/login",
		SplunkQueryUrl:          "/services/search/jobs",
		SleepTimeBetweenQueries: 2500,
		TimeFromNowToQuery:      "-4h",
		Handlers: []string{
			"CheckMerchantStatusHandler",
			"CompleteSessionHandler",
			"PaymentSessionHandler",
			"RegisterMerchantHandler",
			"NonceGeneratorHandler",
			"InAppRewrapHandler",
			"AugmentSessionHandler",
			"PerformPosTransactionHandler",
			"PerformPaymentHandler",
			"GetRegisteredMerchantHandler",
			"RedeemAccountTokenClaimHandler",
			"WebCommerceSessionShareHandler",
			"PrepareTransactionInstructionsHandler",
			"PaylaterCompleteSessionHandler",
			"UpdatePurchaseHandler",
			"InAppPayLaterHandler",
			"VirtualCardPaymentCredentialHandler",
			"PlPaymentChangeHandler",
			"AutoPaymentsHandler",
			"GetPurchasesHandler",
			"PaymentSessionShareHandler",
		},
	}
	return c, nil
}
