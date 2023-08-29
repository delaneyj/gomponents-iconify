package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22v-2h3v-3h2v3.5c0 .4-.2.7-.5 1c-.3.3-.7.5-1 .5H17M7 22H3.5c-.4 0-.7-.2-1-.5c-.3-.3-.5-.7-.5-1V17h2v3h3v2M17 2h3.5c.4 0 .7.2 1 .5c.3.3.5.6.5 1V7h-2V4h-3V2M7 2v2H4v3H2V3.5c0-.4.2-.7.5-1s.6-.5 1-.5H7m12 9H5v2h14v-2Z"/>`),
		g.Group(children),
	)
}