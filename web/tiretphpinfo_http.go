package web

import (
        "github.com/LeakIX/l9format"
        "github.com/PuerkitoBio/goquery"
        "strings"
)

type TiretPhpInfoHttpPlugin struct {
        l9format.ServicePluginBase
}

func (TiretPhpInfoHttpPlugin) GetVersion() (int, int, int) {
        return 0, 0, 1
}

func (TiretPhpInfoHttpPlugin) GetRequests() []l9format.WebPluginRequest {
        return []l9format.WebPluginRequest{{
                Method:  "GET",
                Path:    "/php-info.php",
                Headers: map[string]string{},
                Body:    []byte(""),
        }}
}

func (TiretPhpInfoHttpPlugin) GetName() string {
        return "TiretPhpInfoHttpPlugin"
}

func (TiretPhpInfoHttpPlugin) GetStage() string {
        return "open"
}
func (plugin TiretPhpInfoHttpPlugin) Verify(request l9format.WebPluginRequest, response l9format.WebPluginResponse, event *l9format.L9Event, options map[string]string) bool {
        if !request.EqualAny(plugin.GetRequests()) || response.Response.StatusCode != 200 || response.Document == nil {
                return false
        }
        event.Summary = "Found PHP info page:\n"
        variableTable := response.Document.Find("h2:containsOwn('PHP Variables')").Next()
        if variableTable.Is("table") {
                variableTable.Find("tr").Each(func(i int, selection *goquery.Selection) {
                        if i == 0 {
                                return
                        }
                        event.Summary += strings.TrimSpace(selection.Find("td.e").Text()) + " = " + strings.TrimSpace(selection.Find("td.v").Text()) + "\n"
                })
                return true
        }

        return false
}