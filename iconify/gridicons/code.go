package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23 12l-5.45 6.5L16 17.21L20.39 12L16 6.79l1.55-1.29zM8 6.79L6.45 5.5L1 12l5.45 6.5L8 17.21L3.61 12zm.45 14.61l1.93.52L15.55 2.6l-1.93-.52z"/>`),
		g.Group(children),
	)
}