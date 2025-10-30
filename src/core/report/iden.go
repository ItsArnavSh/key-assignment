package sourceidentifier

import (
	"fmt"
	"os"
	"stack/src/entity"
	"time"
)

// Report generator is responsible for flagging the original issue
// And generate a basic report of the issue
// Then doing a domain search on finding the identity of who could have probably done it
// Finally notify the admin

type Report struct {
	Timestamp time.Time
	Entries   []string
}

// GenerateReport builds a detailed timestamped report string
func (r *Report) GenerateReport(basic entity.BasicReport, url string) {
	r.Timestamp = time.Now()

	report := fmt.Sprintf(`
================ Secret Detection Report ================
Found At     : %s
Source       : %s
----------------------------------------------------------
Provider     : %s
Token Type   : %s
Owner        : %s
Description  : %s
Value        : %s

Context:
%s
==========================================================
`,
		r.Timestamp.Format("2006-01-02 15:04:05"),
		basic.Source,
		basic.Secret.Provider,
		basic.Secret.TokenType,
		basic.Secret.Owner,
		basic.Secret.Description,
		basic.Secret.Value,
		basic.Context,
	)

	r.Entries = append(r.Entries, report)

	info, contr, _ := r.GetInfo(url)
	SendSlackAlert(os.Getenv("SLACK"), basic, info, contr)
}
