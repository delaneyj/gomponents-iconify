package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 3V1H5v2H0v12h16V3h-5zM4 14H3V4h1v10zm6-11H6V2h4v1zm3 11h-1V4h1v10z"/>`),
		g.Group(children),
	)
}