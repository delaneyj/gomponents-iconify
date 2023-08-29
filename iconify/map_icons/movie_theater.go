package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieTheater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M42.215 13h-5.674l5.709-5.474L36.602 2h-7.994l5.727 5.473L28.464 13h-6.789l5.838-5.474L21.746 2h-8.131l5.727 5.473L13.463 13H7.709l5.841-5.474L7.772 2H2v46h46V2h-5.642L48 7.504L42.215 13zM42 42H8v-6h34v6zm0-11H8v-6h34v6z"/>`),
		g.Group(children),
	)
}