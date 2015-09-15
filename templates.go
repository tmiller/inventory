package main

const (
	layoutStart = `
	<html>
		<head>
			<title>{{ .Title }}</title>
		</head>
		<body>
	`

	serverList = `
	{{ template "layoutStart" .}}
	<ul>
	{{ range .Hosts }}
		<li>{{ .Name }}</li>
	{{ else }}
		<li>No Servers Found</li>
	{{ end }}
	</ul>
	{{ template "layoutEnd" }}
`

	layoutEnd = `
		</body>
	</html>
`
)
