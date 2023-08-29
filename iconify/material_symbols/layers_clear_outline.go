package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersClearOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.1 12.8l-1.4-1.45L17.75 9L12 4.55l-2.35 1.8l-1.4-1.45L12 2l9 7l-4.9 3.8Zm2.875 2.85l-1.45-1.45l1.825-1.4L21 14.05l-2.025 1.6Zm.825 6.45l-4-4l-3.8 2.95l-9-7l1.65-1.25L12 18.5l2.35-1.825l-1.425-1.4L12 16L3 9l2.075-1.625l-3.7-3.65L2.8 2.3l18.4 18.4l-1.4 1.4ZM12.175 8.85Z"/>`),
		g.Group(children),
	)
}