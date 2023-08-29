package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StampF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 7.465a4 4 0 1 1 4 0V14H8V7.465zM2 15h16a2 2 0 0 1 2 2v2H0v-2a2 2 0 0 1 2-2z"/>`),
		g.Group(children),
	)
}