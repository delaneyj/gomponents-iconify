package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 3H5a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Zm0 2v4H5V5Zm-10 6h10v7H17Zm-2 7H5v-7h10ZM5 20h10v7H5Zm12 7v-7h10v7Z"/>`),
		g.Group(children),
	)
}