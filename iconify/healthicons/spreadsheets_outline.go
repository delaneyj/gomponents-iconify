package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpreadsheetsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14 19a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H15a1 1 0 0 1-1-1V19Zm2 6v3h4v-3h-4Zm0 8v-3h4v3h-4Zm6 0v-3h4v3h-4Zm6 0v-3h4v3h-4Zm0-8v3h4v-3h-4Zm-6 3h4v-3h-4v3Zm-6-5h16v-3H16v3Z"/><path d="M10 5a1 1 0 0 1 1-1h20a1 1 0 0 1 .707.293l6 6A1 1 0 0 1 38 11v32a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1V5Zm2 1v36h24V12h-5a1 1 0 0 1-1-1V6H12Zm20 1.414V10h2.586L32 7.414Z"/></g>`),
		g.Group(children),
	)
}