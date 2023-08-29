package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoteYea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 5v16h16V5H8zm2 2h12v12H10V7zm9.3 2.9L15 14.2l-2.3-2.3l-1.4 1.5l3 3l.7.7l.7-.7l5-5l-1.4-1.5zM2 19v8h2v-6h2v-2H2zm24 0v2h2v6h2v-8h-4zM6 23v2h20v-2H6z"/>`),
		g.Group(children),
	)
}