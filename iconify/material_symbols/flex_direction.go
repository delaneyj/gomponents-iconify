package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlexDirection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-7h9v7H2Zm0-9V4h9v7H2Zm2-2h5V6H4v3Zm14 11l-5-5l1.4-1.4l2.6 2.575V4h2v12.175l2.6-2.575L23 15l-5 5Z"/>`),
		g.Group(children),
	)
}