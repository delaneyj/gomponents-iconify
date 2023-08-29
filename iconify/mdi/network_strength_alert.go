package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkStrengthAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M18.999 17.003H21V8.999H19m0 12.002H21V19H19M1 21h16.002V7.003H21V1" fill="currentColor"/>`),
		g.Group(children),
	)
}