{{ $reportType := .ReportType }}
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="description" content="{{ $reportType }} report index for your existing checks - generated with trivyalreporting">
        <meta name="author" content="trivyalreporting, a simple report generator tool">
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="/css/styles.css">
        <title>{{ $reportType }} report index for {{ $reportType }} checks - generated with trivyalreporting</title>
    </head>
    <body>
        <h1>{{ $reportType }} reports for your platform</h1>
        {{ range $key, $value := .ReportMap }}
            {{ $reportName := $key }}
            <table>
                <tr class="group-header"><th colspan="2">{{ $reportName }}</th></tr>
                <tr class="sub-header">
                    <th>Scanned resource</th>
                    <th>Direct link to each scanned resource on report {{ $reportName }} (opens in new tab)</th>
                </tr>
            {{range $value := .}}
                {{ $reportResource := $value }}
                <tr class="severity-LOW">
                    <td>Scan result reference for {{ $value }} </td>
                    <td><a href="/{{ $reportType }}/reports/{{ $reportName }}/{{ $value }}" target="_blank">Open stored results for {{ $reportType }} report related to resource {{ $value }}.</a></td>
                </tr>
            {{end}}
            </table>
        {{ end }}
        <hr>
        <div class="footer">
            <p>{{ $reportType }} report index for your existing checks | Report generated using the <a href="https://github.com/ivanol55/trivyalreporting" target="_blank">Trivyalreporting</a> reporting tool.</p>
        </div>
    </body>
</html>
