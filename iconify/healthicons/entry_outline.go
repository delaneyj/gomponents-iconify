package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M23 6.775v34.45l11-3.666V10.442L23 6.775Zm-.684-2.336A1 1 0 0 0 21 5.387v37.226a1 1 0 0 0 1.316.948l13-4.333a1 1 0 0 0 .684-.949V9.721a1 1 0 0 0-.684-.949l-13-4.333Z" clip-rule="evenodd"/><path d="M27 23c0 1.105-.448 2-1 2s-1-.895-1-2s.448-2 1-2s1 .895 1 2Z"/><path fill-rule="evenodd" d="M21 8h-9v31h9v-2h-7V10h7V8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}