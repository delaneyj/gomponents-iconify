package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LessThanOrEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 2.27L5 10.14L18.5 18l1-1.73l-10.53-6.13L19.5 4l-1-1.73M5 20v2h15v-2H5Z"/>`),
		g.Group(children),
	)
}