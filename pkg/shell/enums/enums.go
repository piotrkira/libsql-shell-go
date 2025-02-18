package enums

type PrintMode string

const (
	TABLE_MODE PrintMode = "table"
	CSV_MODE   PrintMode = "csv"
)

type HistoryMode int

const (
	SingleHistory      HistoryMode = 0
	PerDatabaseHistory HistoryMode = 1
	LocalHistory       HistoryMode = 2
)
