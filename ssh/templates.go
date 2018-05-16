package ssh

const sshConfigTemplate = `Host {{ .Name }}
    HostName {{ .Address }}
    {{- if .User }}
    User {{ .User }}
    {{- end }}
    {{- if .Port }}
    Port {{ .Port }}
    {{- end }}
    {{- if .IdentityFile }}
    IdentityFile {{ .IdentityFile }}
    {{- end }}

`
