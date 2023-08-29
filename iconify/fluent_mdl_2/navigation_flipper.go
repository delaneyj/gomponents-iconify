package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationFlipper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m771 1315l290-291l-290-291l90-90l382 381l-382 381l-90-90zM1664 0v2048H256V0h1408zm-128 128H384v1792h1152V128z"/>`),
		g.Group(children),
	)
}