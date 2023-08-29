package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListRadio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 7a2 2 0 1 0-4 0a2 2 0 0 0 4 0Zm2 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM21 4h-8v2h8V4Zm0 7h-8v2h8v-2Zm0 7h-8v2h8v-2ZM6.5 19a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 2a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}