package bullshit

import "html/template"

var bullshitHTML = `<!DOCTYPE html>
<html>
	<head>
        <meta charset="utf-8">
		<style type="text/css">
		table {
			border-collapse: collapse;
		}
		td {
			font: sans-serif 14pt;
			border: 1px solid black;
			padding: 1em;
			vertical-align: middle;
			text-align: center;
			font-weight: bold;
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
				<td id="{{ .Index }}">{{ .Term }}</td>
				{{ end }}
			</tr>
			{{ end }}
		</table>
		<script>
            var i = 0;
            while (true) {
                var element = document.getElementById(i++)
                if (element == null) {
                    break;
                }
                element.addEventListener("click", function(e) {
                    var element = document.getElementById(e.srcElement.id);
                    if (element.style.backgroundColor != "green") {
                        element.style.backgroundColor = "green";
                        element.style.color = "white";
                    } else {
                        element.style.backgroundColor = "white";
                        element.style.color = "black";
                    }
                });
            }
        </script>
	</body>
</html>
`

var BullshitTemplate = template.Must(template.New("bullshit").
	Parse(bullshitHTML))
