package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.7 11l8.67-6.75l8.93 6.98L11.63 18L2.7 11m16 .21L11.39 5.5L4.32 11l7.31 5.73l7.07-5.52M11.63 21L2.7 14l.8-.6l8.11 6.35l7.89-6.16l.8.64L11.63 21Z"/>`),
		g.Group(children),
	)
}