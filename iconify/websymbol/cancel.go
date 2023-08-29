package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="m654 501l346 346l-154 154l-346-346l-346 346L0 847l346-346L0 155L154 1l346 346L846 1l154 154z"/>`),
		g.Group(children),
	)
}