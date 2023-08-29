package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.262 1.457l6.569 12.747L18.5 1.677l1.822.824l-6.434 14.215V22h-2v-5.258L6.182 5.668L4.38 8.862L2.638 7.88l3.624-6.422Z"/>`),
		g.Group(children),
	)
}