package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Performance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 19l-2 3l-3-1l-.5-3.5L3 17l-1-3l3-2l-3-2l1-3l3.5-.5L7 3l3-1l2 3l2-3l3 1l.5 3.5L21 7l1 3l-3 2l3 2l-1 3l-3.5.5L17 21l-3 1l-2-3Zm0-3a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}