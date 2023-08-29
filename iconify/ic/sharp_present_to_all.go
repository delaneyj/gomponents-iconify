package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPresentToAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3H1v18h22V3zm-2 16.02H3V4.98h18v14.04zM10 12H8l4-4l4 4h-2v4h-4v-4z"/>`),
		g.Group(children),
	)
}