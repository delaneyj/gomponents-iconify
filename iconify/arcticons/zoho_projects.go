package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZohoProjects(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.79 12.72L21.55 41.78L3.79 28.79l4.87-6.66l11.1 8.12l16.37-22.4l3.312 2.45l.342.223l-19.167 25.533l-11.687-8.43"/>`),
		g.Group(children),
	)
}