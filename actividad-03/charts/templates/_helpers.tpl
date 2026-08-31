{{- define "actividad-03.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "actividad-03.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{- define "actividad-03.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "actividad-03.labels" -}}
helm.sh/chart: {{ include "actividad-03.chart" . }}
{{ include "actividad-03.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "actividad-03.selectorLabels" -}}
app.kubernetes.io/name: {{ include "actividad-03.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "actividad-03.appLabels" -}}
{{ include "actividad-03.selectorLabels" . }}
app.kubernetes.io/component: app
{{- end }}

{{- define "actividad-03.mysqlLabels" -}}
{{ include "actividad-03.selectorLabels" . }}
app.kubernetes.io/component: mysql
{{- end }}

{{- define "actividad-03.mysqlFullname" -}}
{{- printf "%s-mysql" (include "actividad-03.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "actividad-03.dbHost" -}}
{{- if .Values.db.host }}
{{- .Values.db.host }}
{{- else if .Values.mysql.enabled }}
{{- include "actividad-03.mysqlFullname" . }}
{{- else }}
{{- fail "db.host is required when mysql.enabled is false" }}
{{- end }}
{{- end }}
