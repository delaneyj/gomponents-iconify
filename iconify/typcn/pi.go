package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.707 8.535a.999.999 0 0 0-1.414 0a3.247 3.247 0 0 1-4.586 0a5.25 5.25 0 0 0-7.414 0a.999.999 0 1 0 1.414 1.414c.374-.374.82-.624 1.293-.776V17a1 1 0 1 0 2 0V9.174a3.19 3.19 0 0 1 1.293.775A5.222 5.222 0 0 0 14 11.386V17a1 1 0 1 0 2 0v-5.614a5.215 5.215 0 0 0 2.707-1.437a.999.999 0 0 0 0-1.414z"/>`),
		g.Group(children),
	)
}