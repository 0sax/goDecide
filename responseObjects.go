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
		AccountSweep                string `json:"accountSweep"`
		GamblingRate                int    `json:"gamblingRate"`
		InflowOutflowRate           string `json:"inflowOutflowRate"`
		LoanAmount                  int    `json:"loanAmount"`
		LoanInflowRate              int    `json:"loanInflowRate"`
		LoanRepaymentInflowRate     int    `json:"loanRepaymentInflowRate"`
		LoanRepayments              int    `json:"loanRepayments"`
		TopIncomingTransferAccount  string `json:"topIncomingTransferAccount"`
		TopTransferRecipientAccount string `json:"topTransferRecipientAccount"`
	} `json:"behaviouralAnalysis"`
	CashFlowAnalysis struct {
		AccountActivity           int    `json:"accountActivity"`
		AverageBalance            int    `json:"averageBalance"`
		AverageCredits            int    `json:"averageCredits"`
		AverageDebits             int    `json:"averageDebits"`
		ClosingBalance            int    `json:"closingBalance"`
		FirstDay                  string `json:"firstDay"`
		LastDay                   string `json:"lastDay"`
		MonthPeriod               string `json:"monthPeriod"`
		NetAverageMonthlyEarnings int    `json:"netAverageMonthlyEarnings"`
		NoOfTransactingMonths     int    `json:"noOfTransactingMonths"`
		TotalCreditTurnover       int    `json:"totalCreditTurnover"`
		TotalDebitTurnover        int    `json:"totalDebitTurnover"`
		YearInStatement           string `json:"yearInStatement"`
	} `json:"cashFlowAnalysis"`
	IncomeAnalysis struct {
		AverageOtherIncome                  int         `json:"averageOtherIncome"`
		AverageSalary                       int         `json:"averageSalary"`
		ConfidenceIntervalonSalaryDetection int         `json:"confidenceIntervalonSalaryDetection"`
		ExpectedSalaryDay                   interface{} `json:"expectedSalaryDay"`
		LastSalaryDate                      interface{} `json:"lastSalaryDate"`
		MedianIncome                        int         `json:"medianIncome"`
		NumberOtherIncomePayments           int         `json:"numberOtherIncomePayments"`
		NumberSalaryPayments                int         `json:"numberSalaryPayments"`
		SalaryEarner                        string      `json:"salaryEarner"`
		SalaryFrequency                     interface{} `json:"salaryFrequency"`
	} `json:"incomeAnalysis"`
	SpendAnalysis struct {
		Airtime                        int    `json:"airtime"`
		AtmWithdrawalsSpend            int    `json:"atmWithdrawalsSpend"`
		AverageRecurringExpense        int    `json:"averageRecurringExpense"`
		BankCharges                    int    `json:"bankCharges"`
		Bills                          int    `json:"bills"`
		CableTv                        int    `json:"cableTv"`
		ClubsAndBars                   int    `json:"clubsAndBars"`
		Gambling                       int    `json:"gambling"`
		HasRecurringExpense            string `json:"hasRecurringExpense"`
		InternationalTransactionsSpend int    `json:"internationalTransactionsSpend"`
		PosSpend                       int    `json:"posSpend"`
		ReligiousGiving                int    `json:"religiousGiving"`
		SpendOnTransfers               int    `json:"spendOnTransfers"`
		TotalExpenses                  int    `json:"totalExpenses"`
		UssdTransactions               int    `json:"ussdTransactions"`
		UtilitiesAndInternet           int    `json:"utilitiesAndInternet"`
		WebSpend                       int    `json:"webSpend"`
	} `json:"spendAnalysis"`
	TransactionPatternAnalysis struct {
		MAWWZeroBalanceInAccount struct {
			Month       string `json:"month"`
			WeekOfMonth int    `json:"week_of_month"`
		} `json:"MAWWZeroBalanceInAccount"`
		NODWBalanceLess5000 int `json:"NODWBalanceLess5000"`
		HighestMAWOCredit   struct {
			Month       string `json:"month"`
			WeekOfMonth int    `json:"week_of_month"`
		} `json:"highestMAWOCredit"`
		HighestMAWODebit struct {
			Month       string `json:"month"`
			WeekOfMonth int    `json:"week_of_month"`
		} `json:"highestMAWODebit"`
		LastDateOfCredit                   string        `json:"lastDateOfCredit"`
		LastDateOfDebit                    string        `json:"lastDateOfDebit"`
		MostFrequentBalanceRange           string        `json:"mostFrequentBalanceRange"`
		MostFrequentTransactionRange       string        `json:"mostFrequentTransactionRange"`
		RecurringExpense                   []interface{} `json:"recurringExpense"`
		TransactionsBetween100000And500000 int           `json:"transactionsBetween100000And500000"`
		TransactionsBetween10000And100000  int           `json:"transactionsBetween10000And100000"`
		TransactionsGreater500000          int           `json:"transactionsGreater500000"`
		TransactionsLess10000              int           `json:"transactionsLess10000"`
	} `json:"transactionPatternAnalysis"`
}
