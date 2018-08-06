{{ range $key, $value := .users }}
    {{ $value.Firstname }}<br/>
{{ end }}