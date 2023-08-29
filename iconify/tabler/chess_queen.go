package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessQueen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 16l2-11l-4 4l-2-5l-2 5l-4-4l2 11m0 0l-1.447.724a1 1 0 0 0-.553.894V20h12v-2.382a1 1 0 0 0-.553-.894L16 16H8z"/><path d="M11 4a1 1 0 1 0 2 0a1 1 0 1 0-2 0M5 5a1 1 0 1 0 2 0a1 1 0 1 0-2 0m12 0a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/></g>`),
		g.Group(children),
	)
}