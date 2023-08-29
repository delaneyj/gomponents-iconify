package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessQueenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2a2 2 0 0 1 1.572 3.236l.793 1.983l1.702-1.702A2.003 2.003 0 0 1 18 3a2 2 0 0 1 .674 3.884l-1.69 9.295a1 1 0 0 1-.865.814L16 17H8a1 1 0 0 1-.956-.705l-.028-.116l-1.69-9.295a2 2 0 1 1 2.607-1.367l1.701 1.702l.794-1.983A2 2 0 0 1 12 2zm6 16H6a1 1 0 0 0-1 1a2 2 0 0 0 2 2h10a2 2 0 0 0 1.987-1.768l.011-.174A1 1 0 0 0 18 18z"/></g>`),
		g.Group(children),
	)
}