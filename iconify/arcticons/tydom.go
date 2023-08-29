package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tydom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M24 42.5v-37M24 24H5.5m0 0L24 42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.5 5.5h-27a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h27a8 8 0 0 0 8-8v-21a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}