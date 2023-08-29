package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextCreation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 22.14V9.86A4 4 0 1 0 22.14 5H9.86A4 4 0 1 0 5 9.86v12.28A4 4 0 1 0 9.86 27h12.28A4 4 0 1 0 27 22.14ZM26 4a2 2 0 1 1-2 2a2 2 0 0 1 2-2ZM4 6a2 2 0 1 1 2 2a2 2 0 0 1-2-2Zm2 22a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm16.14-3H9.86A4 4 0 0 0 7 22.14V9.86A4 4 0 0 0 9.86 7h12.28A4 4 0 0 0 25 9.86v12.28A4 4 0 0 0 22.14 25ZM26 28a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/><path fill="currentColor" d="M21 11H11v2h4v9h2v-9h4v-2z"/>`),
		g.Group(children),
	)
}