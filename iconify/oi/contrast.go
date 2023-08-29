package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contrast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c1.66 0 3 1.34 3 3S5.66 7 4 7V1z"/>`),
		g.Group(children),
	)
}