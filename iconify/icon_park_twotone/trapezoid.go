package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trapezoid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTrapezoid0"><path fill="#555" stroke="#fff" stroke-width="4" d="M31.794 8H16.206a3 3 0 0 0-2.864 2.105l-8.125 26A3 3 0 0 0 8.081 40h31.838a3 3 0 0 0 2.864-3.895l-8.125-26A3 3 0 0 0 31.794 8Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTrapezoid0)"/>`),
		g.Group(children),
	)
}