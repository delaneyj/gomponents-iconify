package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="m1000 1l500 500l-500 500l-154-154l346-346l-346-346zM308 501l346 346l-154 154L0 501L500 1l154 154z"/>`),
		g.Group(children),
	)
}