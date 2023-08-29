package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForbidTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm4.891-13.477a6.036 6.036 0 0 0-1.414-1.414l-8.368 8.368a6.041 6.041 0 0 0 1.414 1.414l8.368-8.368Z"/>`),
		g.Group(children),
	)
}