package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyRoseFlowerRosePlantTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 4.1a8.62 8.62 0 0 1 5.5-1.6V6a5.5 5.5 0 0 1-11 0V2.5s6.56-.79 9.56 7.21"/><path d="M9.91 2.76a3 3 0 0 0-5.82 0M7 11.5v2"/></g>`),
		g.Group(children),
	)
}