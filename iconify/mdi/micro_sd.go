package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroSd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2a2 2 0 0 0-2 2v7l-2 2v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8m1 2h2v4H9V4m3 0h2v4h-2V4m3 0h2v4h-2V4Z"/>`),
		g.Group(children),
	)
}