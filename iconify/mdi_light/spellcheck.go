package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spellcheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.5 17.5l.71-.71l3.5 3.5l7.79-7.79l.71.71l-8.5 8.5L8.5 17.5M4.71 13L3.5 16h-1L7.35 4h1.4l4.85 12h-1l-1.21-3H4.71m.41-1H11L8.05 4.74L5.12 12Z"/>`),
		g.Group(children),
	)
}