package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserFullBodyGeometricHumanPersonSingleUpUserFullBody(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M10.5 8a3.5 3.5 0 0 0-7 0v1.5H5l.5 4h3l.5-4h1.5Z"/></g>`),
		g.Group(children),
	)
}