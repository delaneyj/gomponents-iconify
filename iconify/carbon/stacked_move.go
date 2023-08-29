package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M15 28H4a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v15a2 2 0 0 1-2 2zM4 11v15h11V11z" fill="currentColor"/><path d="M26 6l-1.41 1.41L27 10h-5V6a2 2 0 0 0-2-2H10v2h10v14h2v-8h5l-2.41 2.59L26 16l5-5z" fill="currentColor"/>`),
		g.Group(children),
	)
}