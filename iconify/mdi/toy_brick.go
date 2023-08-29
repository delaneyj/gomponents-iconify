package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToyBrick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 6V5a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v1h-2V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v1H3v14h18V6Z"/>`),
		g.Group(children),
	)
}