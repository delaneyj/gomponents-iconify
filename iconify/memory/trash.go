package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 7v9H8V7h2m2 0h2v9h-2V7M8 2h6v1h5v2h-1v14h-1v1H5v-1H4V5H3V3h5V2M6 5v13h10V5H6Z"/>`),
		g.Group(children),
	)
}