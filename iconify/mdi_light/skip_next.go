package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.4 12.5L5 6v13l1-.62l9.4-5.88m-1.9 0L6 17.2V7.8l7.5 4.7M18 6h-1v13h1V6Z"/>`),
		g.Group(children),
	)
}