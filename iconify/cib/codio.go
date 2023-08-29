package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.995 32L2.14 23.995v-16L16 0l13.86 7.995L22 12.536l-6-3.468l-6.005 3.463v6.939l6 3.463l6-3.463l7.865 4.525L16.005 32z"/>`),
		g.Group(children),
	)
}