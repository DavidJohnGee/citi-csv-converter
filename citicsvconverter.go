package citicsvconverter

import (
	"encoding/csv"
	"io"
	"strings"
)

// ----- CONST BLOCK TO MAKE CODE EASIER TO READ. These are exportable by choice. You might create "Mega Dave Expense App"?

const (
	// AccountNumber blah blah
	AccountNumber = iota
	// AccountName blah blah
	AccountName
	// TransactionDate blah blah
	TransactionDate
	// PostDate blah blah
	PostDate
	// ReferenceNumber blah blah
	ReferenceNumber
	// TransactionDetail blah blah
	TransactionDetail
	// BillingAmount blah blah
	BillingAmount
	// SourceCurrency blah blah
	SourceCurrency
	// SourceAmount blah blah
	SourceAmount
	// CustomerRef blah blah
	CustomerRef
	// EmployeeID blah blah
	EmployeeID
)

// DefaultExpenseCategory blah
const DefaultExpenseCategory = "Job Expense"

// ----- END OF CONST BLOCK

// ----- BEGIN OF TYPE BLOCK

// CitiRecord contains the fields. Once marshalled into, makes for an easy Citifree life.
// Note, this is not exported and remains private to the package.
type CitiRecord struct {
	Date     string
	Payee    string
	Category string
	Memo     string
	Outflow  string
	Inflow   string
}

// DumpCSV is a method on a type that dumps a CSV format
func (r CitiRecord) DumpCSV() (*[]string, error) {
	var _row []string
	_row = make([]string, 0)
	_row = append(_row, r.Date)
	_row = append(_row, r.Payee)
	_row = append(_row, r.Category)
	_row = append(_row, r.Memo)
	_row = append(_row, r.Outflow)
	_row = append(_row, r.Inflow)

	return &_row, nil
}

// ----- END OF TYPE BLOCK

// ynanParse contains the logic required to build a slice of CitiRecords
func ynabParse(records [][]string) (*[]CitiRecord, error) {

	recordsToReturn := make([]CitiRecord, 0)

	// Put in the columns
	recordsToReturn = append(recordsToReturn, CitiRecord{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"})

	// read all but the first line (which contain Citi headers)
	for _, r := range records[1:] {

		// _inflow and outflow are the only conditions here we need to worry about
		_inflow := ""
		_outflow := ""

		// Does this field contain a dash? If so, that's money coming back
		if strings.Contains(r[BillingAmount], "-") {
			// If yes, then step forwards to ignore said dash
			_inflow = r[BillingAmount][2:]
		} else {
			// This be expenditure
			_outflow = r[BillingAmount]
		}

		// Create a record. Consts make it much more easier to read
		tempRecord := CitiRecord{
			Date:     r[TransactionDate],
			Payee:    r[TransactionDetail],
			Category: DefaultExpenseCategory,
			Memo:     "",
			Outflow:  _outflow,
			Inflow:   _inflow,
		}

		recordsToReturn = append(recordsToReturn, tempRecord)

	}

	return &recordsToReturn, nil
}

// YnabParser blah blah blah
func YnabParser(r io.Reader) (*[]CitiRecord, error) {

	data := csv.NewReader(r)

	records, err := data.ReadAll()
	if err != nil {
		return nil, err
	}

	sliceofrecords, err := ynabParse(records)

	return sliceofrecords, nil
}

// YnabToCSV builds our CSV
func YnabToCSV(r *[]CitiRecord) ([][]string, error) {

	// Create the return data with at least one entry (else what's the point?)
	_returnData := make([][]string, 0)

	// Loop through each record and build a slice of strings for each entry
	for _, record := range *r {

		csvData, _ := record.DumpCSV()
		_returnData = append(_returnData, *csvData)
	}

	return _returnData, nil
}
