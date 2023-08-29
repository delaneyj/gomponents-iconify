package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4.094l-.719.687l-8.5 8.5L8.22 14.72L15 7.938V28h2V7.937l6.781 6.782l1.438-1.438l-8.5-8.5z"/>`),
		g.Group(children),
	)
}