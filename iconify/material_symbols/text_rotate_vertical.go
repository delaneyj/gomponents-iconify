package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotateVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.9 16L15 5h2l4.1 11h-1.9l-1-2.8h-4.4l-1 2.8h-1.9Zm3.45-4.4h3.3l-1.6-4.55h-.1l-1.6 4.55ZM6 20l-3.5-3.5l1.4-1.4L5 16.15V3h2v13.15l1.1-1.05l1.4 1.4L6 20Z"/>`),
		g.Group(children),
	)
}