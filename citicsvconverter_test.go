package citicsvconverter

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

// ----- TREATING THESE AS INPUTS TO TESTS

var wanted = [][]string{
	{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"},
	{"01/01/2018", "my company", "Job Expense", "", "", "1,000"},
	{"02/01/2018", "a shop", "Job Expense", "", "10.00", ""},
}

var expectedrecords = []CitiRecord{
	{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"},
	{"01/01/2018", "my company", "Job Expense", "", "", "1,000"},
	{"02/01/2018", "a shop", "Job Expense", "", "10.00", ""}}

const inputted = `"Account Number","Account Name","Transaction Date","Post Date","Reference Number","Transaction Detail","Billing Amount","Source Currency","Source Amount","Customer Ref","Employee Number"
"XXXXXXXXXX","Foo","01/01/2018","02/01/2018","12345","my company"," -1,000","GBP"," -1,000",,"98765"
"XXXXXXXXXX","Foo","02/01/2018","03/01/2018","23456","a shop","10.00","GBP","10.00",,"98765"`

// ----- END OF INPUTS TO TEST

// ----- BEGIN TESTS
//	TestYnabFullParser - does an end-to-end test and receives a new CSV
//  TestYnabParser - posts the input and checks that the slice of structs looks good
//

// TestYnabParser - stop moaning at me linter!
func TestFullYnabParser(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args
		want    [][]string
		wantErr bool
	}{
		{
			name:    "a test",
			args:    args{strings.NewReader(inputted)},
			want:    wanted,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get output with pointer to slice of structs
			output, err := YnabParser(tt.args.r)
			if err != nil {
				t.Errorf("Could not parse data: %v", err)
			}
			// pass pointer to CSV converter
			got, err := YnabToCSV(output)
			if err != nil {
				t.Errorf("Could not create CSV: %v", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("YnabParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YnabParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYnabParser(t *testing.T) {
	// Tests structs and slices of
	// We can detect layered issues like this
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args
		want    []CitiRecord
		wantErr bool
	}{
		{
			name:    "a test",
			args:    args{strings.NewReader(inputted)},
			want:    expectedrecords,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get output with pointer to slice of structs
			got, err := YnabParser(tt.args.r)
			if err != nil {
				t.Errorf("Could not parse data: %v", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("YnabParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("YnabParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
