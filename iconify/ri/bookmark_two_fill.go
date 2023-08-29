package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h14a1 1 0 0 1 1 1v19.143a.5.5 0 0 1-.766.424L12 18.03l-7.234 4.536A.5.5 0 0 1 4 22.143V3a1 1 0 0 1 1-1Zm3 7v2h8V9H8Z"/>`),
		g.Group(children),
	)
}