package indicina

type loginResponse struct {
	Status string `json:"status"`
	Data   struct {
		Token string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (lr *loginResponse) isSuccess() bool {
	return lr.Status == "success"
}

type analysisResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Code    int              `json:"code"`
	Data    StatementSummary `json:"data"`
}
type analysisResponsePDF struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Code    int                  `json:"code"`
	Data    PDFStatementResponse `json:"data"`
}

func (ar *analysisResponsePDF) isSuccess() bool {
	return ar.Status == "success"
}

func (ar *analysisResponse) isSuccess() bool {
	return ar.Status == "success"
}

type PDFStatementResponse struct {
	JobId          string           `json:"jobId"`
	Status         string           `json:"status"`
	DecideResponse StatementSummary `json:"decideResponse""`
}

type StatementSummary struct {
	BehaviouralAnalysis struct {
		AccountSweep                string  `json:"accountSweep"`
		GamblingRate                float64 `json:"gamblingRate"`
		InflowOutflowRate           string  `json:"inflowOutflowRate"`
		LoanAmount                  float64 `json:"loanAmount"`
		LoanInflowRate              float64 `json:"loanInflowRate"`
		LoanRepaymentInflowRate     float64 `json:"loanRepaymentInflowRate"`
		LoanRepayments              float64 `json:"loanRepayments"`
		TopIncomingTransferAccount  string  `json:"topIncomingTransferAccount"`
		TopTransferRecipientAccount string  `json:"topTransferRecipientAccount"`
	} `json:"behaviouralAnalysis"`
	CashFlowAnalysis struct {
		AccountActivity           float64 `json:"accountActivity"`
		AverageBalance            float64 `json:"averageBalance"`
		AverageCredits            float64 `json:"averageCredits"`
		AverageDebits             float64 `json:"averageDebits"`
		ClosingBalance            float64 `json:"closingBalance"`
		FirstDay                  string  `json:"firstDay"`
		LastDay                   string  `json:"lastDay"`
		MonthPeriod               string  `json:"monthPeriod"`
		NetAverageMonthlyEarnings float64 `json:"netAverageMonthlyEarnings"`
		NoOfTransactingMonths     float64 `json:"noOfTransactingMonths"`
		TotalCreditTurnover       float64 `json:"totalCreditTurnover"`
		TotalDebitTurnover        float64 `json:"totalDebitTurnover"`
		YearInStatement           string  `json:"yearInStatement"`
	} `json:"cashFlowAnalysis"`
	IncomeAnalysis struct {
		AverageOtherIncome                  float64     `json:"averageOtherIncome"`
		AverageSalary                       float64     `json:"averageSalary"`
		ConfidenceIntervalonSalaryDetection float64     `json:"confidenceIntervalonSalaryDetection"`
		ExpectedSalaryDay                   interface{} `json:"expectedSalaryDay"`
		LastSalaryDate                      interface{} `json:"lastSalaryDate"`
		MedianIncome                        float64     `json:"medianIncome"`
		NumberOtherIncomePayments           float64     `json:"numberOtherIncomePayments"`
		NumberSalaryPayments                float64     `json:"numberSalaryPayments"`
		SalaryEarner                        string      `json:"salaryEarner"`
		SalaryFrequency                     interface{} `json:"salaryFrequency"`
	} `json:"incomeAnalysis"`
	SpendAnalysis struct {
		Airtime                        float64 `json:"airtime"`
		AtmWithdrawalsSpend            float64 `json:"atmWithdrawalsSpend"`
		AverageRecurringExpense        float64 `json:"averageRecurringExpense"`
		BankCharges                    float64 `json:"bankCharges"`
		Bills                          float64 `json:"bills"`
		CableTv                        float64 `json:"cableTv"`
		ClubsAndBars                   float64 `json:"clubsAndBars"`
		Gambling                       float64 `json:"gambling"`
		HasRecurringExpense            string  `json:"hasRecurringExpense"`
		InternationalTransactionsSpend float64 `json:"internationalTransactionsSpend"`
		PosSpend                       float64 `json:"posSpend"`
		ReligiousGiving                float64 `json:"religiousGiving"`
		SpendOnTransfers               float64 `json:"spendOnTransfers"`
		TotalExpenses                  float64 `json:"totalExpenses"`
		UssdTransactions               float64 `json:"ussdTransactions"`
		UtilitiesAndInternet           float64 `json:"utilitiesAndInternet"`
		WebSpend                       float64 `json:"webSpend"`
	} `json:"spendAnalysis"`
	TransactionPatternAnalysis struct {
		MAWWZeroBalanceInAccount struct {
			Month       string  `json:"month"`
			WeekOfMonth float64 `json:"week_of_month"`
		} `json:"MAWWZeroBalanceInAccount"`
		NODWBalanceLess5000 float64 `json:"NODWBalanceLess5000"`
		HighestMAWOCredit   struct {
			Month       string  `json:"month"`
			WeekOfMonth float64 `json:"week_of_month"`
		} `json:"highestMAWOCredit"`
		HighestMAWODebit struct {
			Month       string  `json:"month"`
			WeekOfMonth float64 `json:"week_of_month"`
		} `json:"highestMAWODebit"`
		LastDateOfCredit                   string        `json:"lastDateOfCredit"`
		LastDateOfDebit                    string        `json:"lastDateOfDebit"`
		MostFrequentBalanceRange           interface{}   `json:"mostFrequentBalanceRange"`
		MostFrequentTransactionRange       interface{}   `json:"mostFrequentTransactionRange"`
		RecurringExpense                   []interface{} `json:"recurringExpense"`
		TransactionsBetween100000And500000 float64       `json:"transactionsBetween100000And500000"`
		TransactionsBetween10000And100000  float64       `json:"transactionsBetween10000And100000"`
		TransactionsGreater500000          float64       `json:"transactionsGreater500000"`
		TransactionsLess10000              float64       `json:"transactionsLess10000"`
	} `json:"transactionPatternAnalysis"`
}
