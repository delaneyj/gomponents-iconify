package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditBoxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.757 2.997l-2 2H5v14h14V9.239l2-2v12.758a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1h12.757Zm3.728-.9L21.9 3.511l-9.193 9.193l-1.412.002l-.002-1.416l9.192-9.193Z"/>`),
		g.Group(children),
	)
}