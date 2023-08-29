package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreetRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm14.372 2.343L6.657 13.372l-1.029-1.715l2.111-1.267L5.64 6.611l1.749-.97l2.066 3.72l7.888-4.733l1.029 1.715Zm0 6l-2.111 1.267l2.099 3.779l-1.748.97l-2.067-3.72l-7.888 4.733l-1.029-1.715l11.715-7.029l1.029 1.715Z"/>`),
		g.Group(children),
	)
}