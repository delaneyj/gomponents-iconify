package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowULeftUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 80v88a64 64 0 0 1-128 0V88H40a8 8 0 0 1-5.66-13.66l48-48a8 8 0 0 1 11.32 0l48 48A8 8 0 0 1 136 88H96v80a48 48 0 0 0 96 0V80a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}