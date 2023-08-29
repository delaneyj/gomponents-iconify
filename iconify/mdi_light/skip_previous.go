package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipPrevious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.6 12.5L18 6v13l-1-.62l-9.4-5.88m1.9 0l7.5 4.7V7.8l-7.5 4.7M5 6h1v13H5V6Z"/>`),
		g.Group(children),
	)
}