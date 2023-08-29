package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ccx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M10 23H5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5v2H5v6h5z" fill="currentColor"/><path d="M19 23h-5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5v2h-5v6h5z" fill="currentColor"/><path d="M29 9h-2l-2 6l-2-6h-2l2.75 7L21 23h2l2-6l2 6h2l-2.75-7L29 9z" fill="currentColor"/>`),
		g.Group(children),
	)
}