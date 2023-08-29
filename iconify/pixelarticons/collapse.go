package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collapse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3h-2v2h-2v2h-2V5H9V3H7v2h2v2h2v2h2V7h2V5h2V3zM4 13h16v-2H4v2zm9 4h-2v-2h2v2zm2 2h-2v-2h2v2zm0 0h2v2h-2v-2zm-6 0h2v-2H9v2zm0 0H7v2h2v-2z"/>`),
		g.Group(children),
	)
}