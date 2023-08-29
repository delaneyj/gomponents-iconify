package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9V7h2v2H8zm-2 2V9h2v2H6zm0 2H4v-2h2v2zm2 2v-2H6v2h2zm0 0h2v2H8v-2zm8-6V7h-2v2h2zm2 2V9h-2v2h2zm0 2v-2h2v2h-2zm-2 2v-2h2v2h-2zm0 0v2h-2v-2h2z"/>`),
		g.Group(children),
	)
}