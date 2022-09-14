<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
{{- if . }}
    <style>
      * {
        font-family: Arial, Helvetica, sans-serif;
      }
      h1 {
        text-align: center;
      }
      .group-header th {
        font-size: 200%;
        padding: 0.5em;
      }
      .sub-header th {
        font-size: 150%;
        padding: 0.4em;
      }
      table, th, td {
        border: 1px solid black;
        border-collapse: collapse;
      }
      table, tr, td {
        padding: 0.5em;
      }
      table {
        margin: 0 auto;
      }
      .severity {
        text-align: center;
        font-weight: bold;
        color: #fafafa;
      }
      .severity-LOW .severity { background-color: #5fbb31; }
      .severity-MEDIUM .severity { background-color: #e9c600; }
      .severity-HIGH .severity { background-color: #ff8800; }
      .severity-CRITICAL .severity { background-color: #e40000; }
      .severity-UNKNOWN .severity { background-color: #747474; }
      .severity-LOW { background-color: #5fbb3160; }
      .severity-MEDIUM { background-color: #e9c60060; }
      .severity-HIGH { background-color: #ff880060; }
      .severity-CRITICAL { background-color: #e4000060; }
      .severity-UNKNOWN { background-color: #74747460; }
      table tr td:first-of-type {
        font-weight: bold;
        white-space: nowrap;
        padding: 1em;
      }
      .links a,
      .links[data-more-links=on] a {
        display: block;
      }
      .links[data-more-links=off] a:nth-of-type(1n+5) {
        display: none;
      }
      a.toggle-more-links { cursor: pointer; }
    </style>
    <title>Infrastructure report for your AWS account - {{ now.Format "2006-01-02" }} {{ now.Format "15:04:05" }}</title>
  </head>
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
{{- else }}
  </head>
  <body>
    <h1>No Vulnerabilities found for scanned resources in the environment.</h1>
{{- end }}
  </body>
</html>