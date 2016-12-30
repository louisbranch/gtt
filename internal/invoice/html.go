package invoice

import (
	"fmt"
	"os"
	"text/template"
)

const html = `
<html>
  <head>

		<link rel="stylesheet" href="https://storage.googleapis.com/code.getmdl.io/1.0.4/material.indigo-cyan.min.css" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>
			Invoice -
			{{.From.Format "Jan 02, 2006" }} / {{.To.Format "Jan 02, 2006" }}
		</title>
    <style type="text/css" media="print">
      * {
          overflow: visible !important;
      }
    </style>
  </head>
	<body style="background-color: #eee">
	<div class="mdl-layout">
		<header class="mdl-layout__header">
			<div class="mdl-layout__header-row" style="padding-left: 20px">
				<div class="mdl-layout-title">
					Invoice
				</div>
				<div style="margin-left: 20px">
					<strong>From:</strong> {{.From.Format "Jan 02, 2006" }}
					-
					<strong>To:</strong> {{.To.Format "Jan 02, 2006" }}
				</div>
				<div class="mdl-layout-spacer"></div>
				<div style="text-align: right">
					<strong>Hours Worked:</strong> {{.DurationFormated}}<br>
					<strong>Amount:</strong> ${{ printf "%.2f" .TotalCost}}
				</div>
			</div>
		</header>
	<div class="mdl-grid">
	{{ range .Days }}
		<div class="mdl-color--white mdl-shadow--2dp mdl-cell mdl-cell--12-col mdl-card">
			<div class="mdl-card__title mdl-color--grey">
				{{ .Start.Format "Jan 02, Monday" }}
				<div class="mdl-layout-spacer"></div>
				{{if len .Pauses }}
					<div>
						<strong>Pauses:</strong>
						{{ range $i, $Pause := .Pauses }}
							{{if $i }}/{{ end }}
							{{ $Pause.Start.Format "15:04" }}
							-
							{{ $Pause.End.Format "15:04" }}
						{{ end }}
					</div>
				{{ end }}
			</div>
		<div class="mdl-card__media">
		<table class="mdl-data-table mdl-shadow--2dp" style="width: 100%">
				<tbody>
					<tr>
						<td class="mdl-data-table__cell--non-numeric" style="width: 20%">
							{{ .Start.Format "15:04" }}
						</td>
						<td class="mdl-data-table__cell--non-numeric">
							{{ .Start.Format "15:04" }}
							Start
						</td>
					</tr>
					{{ range .Tasks }}
						<tr>
							<td class="mdl-data-table__cell--non-numeric" style="width: 20%">
								{{ .End.Format "15:04" }}
							</td>
							<td class="mdl-data-table__cell--non-numeric">
								{{ .Description }}
							</td>
						</tr>
					{{ end }}
					<tr>
						<td class="mdl-data-table__cell--non-numeric" style="width: 20%">
							Duration
						</td>
						<td class="mdl-data-table__cell--non-numeric">
							<strong>
								{{ .Status }}
							</strong>
						</td>
					</tr>
				</tbody>
			</table>
			</div>
		</div>
			{{ end }}
		</div>
  </body>
</html>
`

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.New("invoice").Parse(html)
	if err != nil {
		panic(err)
	}
}

// ToHTML formats the days and cost using a HTML template
func (i *Invoice) ToHTML() error {
	err := tmpl.Execute(os.Stdout, i)
	if err != nil {
		return fmt.Errorf("Invoice template error %s", err.Error())
	}
	return nil
}
