package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 7c2.8 0 5 2.2 5 5c0 1.3-.5 2.4-1.3 3.3c1.9.7 3.3 2.5 3.3 4.7c0 2.8-2.2 5-5 5H7V7h9m-7 8h7c1.7 0 3-1.3 3-3s-1.3-3-3-3H9v6m0 8h9c1.7 0 3-1.3 3-3s-1.3-3-3-3H9v6m7-18H5v22h13c3.9 0 7-3.1 7-7c0-2.1-1-4.1-2.5-5.4c.3-.8.5-1.7.5-2.6c0-3.9-3.1-7-7-7zm-5 6h5c.6 0 1 .4 1 1s-.4 1-1 1h-5v-2zm0 8h7c.6 0 1 .4 1 1s-.4 1-1 1h-7v-2z"/>`),
		g.Group(children),
	)
}