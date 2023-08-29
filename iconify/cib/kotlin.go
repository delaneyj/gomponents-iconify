package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kotlin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m1.734 32l15.068-15.333L32 32zM0 0h16L0 16.667zm17.865 0L0 18.667V32L32 0z"/>`),
		g.Group(children),
	)
}