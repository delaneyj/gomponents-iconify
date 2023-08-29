package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarProtonThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.072 26.31a3.761 3.761 0 0 1 3.75 3.75h0a3.761 3.761 0 0 1-3.75 3.75h-1.5c-2.625 0-3.562-.374-4.687-1.312"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.885 20.123c1.125-.937 2.25-1.312 4.687-1.312h1.5a3.761 3.761 0 0 1 3.75 3.75h0a3.761 3.761 0 0 1-3.75 3.75h-3.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 25.793v-13.93a4 4 0 0 0-4-4h-31a4 4 0 0 0-4 4v24.275a4 4 0 0 0 4 4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 36.138V11.862a4 4 0 0 0-4-4h-31a4 4 0 0 0-4 4v24.276a4 4 0 0 0 4 4h31a4 4 0 0 0 4-4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.207 40.138V15.819a3.138 3.138 0 0 0-3.138-3.138H4.5"/>`),
		g.Group(children),
	)
}