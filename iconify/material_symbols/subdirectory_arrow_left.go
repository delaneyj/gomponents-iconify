package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubdirectoryArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 21l-6-6l6.025-6.025l1.4 1.425l-3.6 3.6H17V4h2v12H7.825l3.6 3.575L10 21Z"/>`),
		g.Group(children),
	)
}