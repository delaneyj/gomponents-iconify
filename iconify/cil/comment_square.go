package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M496 496h-47.229l-69.522-128H40a24.028 24.028 0 0 1-24-24V40a24.028 24.028 0 0 1 24-24h432a24.028 24.028 0 0 1 24 24ZM48 336h350.284L464 456.993V48H48Z"/>`),
		g.Group(children),
	)
}