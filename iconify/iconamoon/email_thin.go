package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5v-.5a.5.5 0 0 0-.5.5H3Zm18 0h.5a.5.5 0 0 0-.5-.5V5ZM3 5.5h18v-1H3v1ZM20.5 5v12h1V5h-1ZM19 18.5H5v1h14v-1ZM3.5 17V5h-1v12h1ZM5 18.5A1.5 1.5 0 0 1 3.5 17h-1A2.5 2.5 0 0 0 5 19.5v-1ZM20.5 17a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5h-1Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3 5l9 9l9-9"/></g>`),
		g.Group(children),
	)
}