package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSamsungThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.137 24a3.51 3.51 0 0 1 3.5 3.5h0a3.51 3.51 0 0 1-3.5 3.5h-1.4c-2.45 0-3.324-.35-4.374-1.225m0-11.55c1.05-.875 2.1-1.225 4.375-1.225h1.4a3.51 3.51 0 0 1 3.5 3.5h0a3.51 3.51 0 0 1-3.5 3.5h-3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 5.5h-4v1a3.5 3.5 0 1 1-7 0v-1h-8v1a3.5 3.5 0 1 1-7 0v-1h-3a4 4 0 0 0-4 4v29a4 4 0 0 0 4 4h29a4 4 0 0 0 4-4v-29a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}