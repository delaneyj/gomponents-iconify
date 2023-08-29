package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudyPrevious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M8.07 21L2 15l6.07-6l1.43 1.41L4.86 15l4.64 4.59L8.07 21z" fill="currentColor"/><path d="M20 4a5 5 0 1 1-5 5a5 5 0 0 1 5-5m0-2a7 7 0 1 0 7 7a7 7 0 0 0-7-7z" fill="currentColor"/><path d="M30 30h-2v-5a5 5 0 0 0-5-5h-6a5 5 0 0 0-5 5v5h-2v-5a7 7 0 0 1 7-7h6a7 7 0 0 1 7 7z" fill="currentColor"/>`),
		g.Group(children),
	)
}