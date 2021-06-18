{{/*
Return the appropriate apiVersion for deployment.
*/}}
{{- define "daiops.deployment.apiVersion" -}}
{{- print "apps/v1" -}}
{{- end -}}


{{/*
Return the appropriate apiVersion for service.
*/}}
{{- define "daiops.service.apiVersion" -}}
{{- print "v1" -}}
{{- end -}}

{{/*
Return the appropriate apiVersion for namespace.
*/}}
{{- define "daiops.namespace.apiVersion" -}}
{{- print "v1" -}}
{{- end -}}


{{- define "daiops.metadata.namespace" -}}
{{- .Release.Namespace -}}
{{- end -}}


