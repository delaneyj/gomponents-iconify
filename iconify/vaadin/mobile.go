package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 1v14h8V1H4zm5 13H7v-1h2v1zm2-2H5V3h6v9z"/>`),
		g.Group(children),
	)
}