package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 144a56.06 56.06 0 0 1-56 56a32 32 0 0 0-32 32a8 8 0 0 1-16 0a32 32 0 0 0-32-32a56.06 56.06 0 0 1-56-56a8 8 0 0 1 16 0a40 40 0 0 0 40 40a47.79 47.79 0 0 1 32 12.27V128H88a8 8 0 0 1 0-16h32V82.83a28 28 0 1 1 16 0V112h32a8 8 0 0 1 0 16h-32v68.27A47.79 47.79 0 0 1 168 184a40 40 0 0 0 40-40a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}