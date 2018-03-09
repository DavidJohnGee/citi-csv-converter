# Citi CSV Converter
Converters CSV statements from Citi Bank Credit Card into a format that personal finance apps can understand.

Usage example:

```go
package main

import (
	"encoding/csv"
	"log"
	"os"

	cc "github.com/davidjohngee/citicsvconverter"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("cannot read csv: %v", err)
	}
	defer data.Close()

	// Let's parse them datas
	output, err := cc.YnabParser(data)
	if err != nil {
		log.Fatalf("couldn't parse csv file: %v", err)
	}

	// Build new CitiCSV
	w := csv.NewWriter(os.Stdout)

	// Get CSV
	csv, err := cc.YnabToCSV(output)
	if err != nil {
		log.Fatalf("Create CSV file: %v", err)
	}

	// We now build a writer func
	w.WriteAll(csv)

	if err = w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

```


Build the code, then run:

```bash
go build
./app testdata.csv

Date,Payee,Category,Memo,Outflow,Inflow
01/01/2018,my company,Job Expense,,,"1,000"
02/01/2018,a shop,Job Expense,,10.00,
```

testdata.csv contents
```text
testdata.csv
"Account Number","Account Name","Transaction Date","Post Date","Reference Number","Transaction Detail","Billing Amount","Source Currency","Source Amount","Customer Ref","Employee Number"
"XXXXXXXXXX","Foo","01/01/2018","02/01/2018","12345","my company"," -1,000","GBP"," -1,000",,"98765"
"XXXXXXXXXX","Foo","02/01/2018","03/01/2018","23456","a shop","10.00","GBP","10.00",,"98765"
```

This library has two basic tests. Go to the project directory and run the tests as below:
```bash
go test -v
=== RUN   TestFullYnabParser
=== RUN   TestFullYnabParser/a_test
--- PASS: TestFullYnabParser (0.00s)
    --- PASS: TestFullYnabParser/a_test (0.00s)
=== RUN   TestYnabParser
=== RUN   TestYnabParser/a_test
--- PASS: TestYnabParser (0.00s)
    --- PASS: TestYnabParser/a_test (0.00s)
PASS
ok  	github.com/davidjohngee/citi-csv-converter	0.006s
```

