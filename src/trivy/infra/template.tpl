<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="description" content=" Infrastructure report for AWS account - {{ now.Format "2006-01-02" }} {{ now.Format "15:04:05" }} generated with trivyalreporting">
    <meta name="author" content="trivyalreporting, a trivy report generator tool">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/css/styles.css">
    <title>Infrastructure report for your AWS account - {{ now.Format "2006-01-02" }} {{ now.Format "15:04:05" }}</title>
  </head>
{{- if . }}
  <body>
    <h1>Infrastructure report for your AWS account - {{ now.Format "2006-01-02" }} {{ now.Format "15:04:05" }}</h1>
    <div>
    <table>
    {{- range . }}
      <tr class="group-header"><th colspan="5">{{ escapeXML .Target }}</th></tr>
      {{- if (gt (len .Misconfigurations) 0) }}
      <tr class="sub-header">
        <th>ID</th>
        <th>Issue</th>
        <th>Severity</th>
        <th>Description</th>
        <th>Resolution</th>
      </tr>
        {{- range .Misconfigurations }}
      <tr class="severity-{{ escapeXML .Severity }}">
        <td><a href={{ escapeXML .PrimaryURL }} target="_blank">{{ escapeXML .ID }}</a></td>
        <td>{{ escapeXML .Title }}</td>
        <td class="severity">{{ escapeXML .Severity }}</td>
        <td>{{ escapeXML .Description }}</td>
        <td>{{ escapeXML .Resolution }}</td>
      </tr>
        {{- end }}
      {{- end }}
    {{- end }}
    </table>
    </div>
  </body>
{{- else }}
  </head>
  <body>
    <h1>No Vulnerabilities found for scanned resources in the environment.</h1>
  </body>
{{- end }}
</html>