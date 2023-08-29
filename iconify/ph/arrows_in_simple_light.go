package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212.24 52.24L158.48 106H192a6 6 0 0 1 0 12h-48a6 6 0 0 1-6-6V64a6 6 0 0 1 12 0v33.52l53.76-53.76a6 6 0 0 1 8.48 8.48ZM112 138H64a6 6 0 0 0 0 12h33.52l-53.76 53.76a6 6 0 1 0 8.48 8.48L106 158.48V192a6 6 0 0 0 12 0v-48a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}