package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrophyF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.294 15c-.6-1.825-1.032-3.825-1.294-6c-.362-3-.362-6 0-9h8c.311 3.329.311 6.329 0 9a32.729 32.729 0 0 1-1.287 6H8a2 2 0 0 1 2 2v3H0v-3a2 2 0 0 1 2-2h.294zM3 16v2h4v-2H3z"/>`),
		g.Group(children),
	)
}