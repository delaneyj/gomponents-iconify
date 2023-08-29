package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserHumanGeometricHumanPersonSingleUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 13.5V11h2a1 1 0 0 0 1-1V6a5.5 5.5 0 1 0-8 4.9v2.6"/>`),
		g.Group(children),
	)
}