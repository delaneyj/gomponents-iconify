package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserSquareAlternateSquareGeometricHumanPersonSingleUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="5.5" r="2.5"/><path d="M2.1 13.5a5 5 0 0 1 9.8 0"/><rect width="13" height="13" x=".5" y=".5" rx="1"/></g>`),
		g.Group(children),
	)
}