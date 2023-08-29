package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineKeyboardTab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.59 7.41L15.17 11H1v2h14.17l-3.59 3.59L13 18l6-6l-6-6l-1.41 1.41zM20 6v12h2V6h-2z"/>`),
		g.Group(children),
	)
}