package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Froward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m0 382.5l256-128l-256-128v256zm512-128l-256-128v256l256-128z"/>`),
		g.Group(children),
	)
}