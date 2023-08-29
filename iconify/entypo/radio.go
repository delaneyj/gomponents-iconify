package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 8H5.021l8.974-5.265L13 1L1.736 7.571A1.482 1.482 0 0 0 1 8.852V17a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2zm-1.5 9a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 15.5 17zm1.5-5H3v-2h14v2z"/>`),
		g.Group(children),
	)
}