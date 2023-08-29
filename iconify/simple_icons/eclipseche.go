package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eclipseche(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0L1.604 6.021v7.452L12 7.494l3.941 2.254l6.455-3.727zm10.396 10.527L12 16.506l-7.334-4.217l-3.062 1.76v3.93L12 24l10.396-6.021z"/>`),
		g.Group(children),
	)
}