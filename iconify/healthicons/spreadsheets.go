package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spreadsheets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M16 20v3h16v-3H16Zm0 8v-3h4v3h-4Zm0 2v3h4v-3h-4Zm6 0v3h4v-3h-4Zm6 0v3h4v-3h-4Zm0-2v-3h4v3h-4Zm-2 0h-4v-3h4v3Z"/><path fill-rule="evenodd" d="M30 4H11a1 1 0 0 0-1 1v38a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V12h-7a1 1 0 0 1-1-1V4ZM14 19a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H15a1 1 0 0 1-1-1V19Z" clip-rule="evenodd"/><path d="M37.414 10H32V4.586L37.414 10Z"/></g>`),
		g.Group(children),
	)
}