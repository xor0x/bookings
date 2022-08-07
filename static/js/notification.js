// setup message types using notie via notify()
{{with .Flash}}
notify("{{.}}", "success");
{{end}}
 
{{with .Warning}}
notify("{{.}}", "warning");
{{end}}
 
{{with .Error}}
notify("{{.}}", "error");
{{end}}