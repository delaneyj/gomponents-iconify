package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 14a1 1 0 0 0 .894-1.447L19.118 9l1.776-3.553A1 1 0 0 0 20 4H9V3a1 1 0 1 0-2 0v18a1 1 0 1 0 2 0v-7h11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}