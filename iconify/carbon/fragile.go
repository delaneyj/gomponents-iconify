package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fragile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 12V6h-2v6a5 5 0 0 1-10 0V4h5.586l-2.293 2.293a1 1 0 0 0 0 1.414L15.586 9l-2.293 2.293l1.414 1.414l3-3a1 1 0 0 0 0-1.414L16.414 7l3.293-3.293A1 1 0 0 0 19 2h-9a1 1 0 0 0-1 1v9a7.005 7.005 0 0 0 6 6.92V28h-5v2h12v-2h-5v-9.08A7.005 7.005 0 0 0 23 12Z"/>`),
		g.Group(children),
	)
}