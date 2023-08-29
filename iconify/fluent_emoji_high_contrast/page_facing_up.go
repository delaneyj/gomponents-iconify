package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageFacingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M9.5 12a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-13Zm0 3a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-13ZM9 18.5a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 0 1h-13a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1h-8Z"/><path d="M4 4a3 3 0 0 1 3-3h13.05a1 1 0 0 1 .707.293l6.95 6.95A1 1 0 0 1 28 8.95V28a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V4Zm3-1a1 1 0 0 0-1 1v24a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V9.364L25.636 9h-4.643A.997.997 0 0 1 20 8V3.364L19.636 3H7Z"/></g>`),
		g.Group(children),
	)
}