    {{- range .}}
        {{- if (gt (len .Misconfigurations) 0) }}
            <table>
                    <tr class="group-header"><th colspan="5">{{ escapeXML .Target }}</th></tr>
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
            </table>
        {{- end }}
    {{- end }}
    <br><br><br><br>