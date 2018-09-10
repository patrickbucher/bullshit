package bullshit

import "html/template"

var bullshitHTML = `
<!DOCTYPE html>
<html>
	<head>
		<style type="text/css">
		td {
			font: sans-serif 14pt;
			border: 1px solid black;
			padding: 1em;
			vertical-align: middle;
			text-align: center;
		}
		</style>
		<title>Bullshit Bingo</title>
	</head>
	<body>
		<h1>Bullshit Bingo</h1>
		<table>
			{{ range .Rows }}
			<tr>
				{{ range .Cols }}
				<td>{{ .Term }}</td>
				{{ end }}
			</tr>
			{{ end }}
		</table>
	</body>
</html>
`

var BullshitTemplate = template.Must(template.New("bullshit").
	Parse(bullshitHTML))
