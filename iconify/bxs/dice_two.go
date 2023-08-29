package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2zM9.5 13.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 9.5 13.5zm5 0a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 14.5 13.5z"/>`),
		g.Group(children),
	)
}