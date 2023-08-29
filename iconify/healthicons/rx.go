package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 24C4 12.954 12.954 4 24 4s20 8.954 20 20s-8.954 20-20 20S4 35.046 4 24Zm12-10a1 1 0 0 1 1-1h6a5 5 0 0 1 .398 9.984L28 27.586l4.293-4.293l1.414 1.414L29.414 29l4.293 4.293l-1.414 1.414L28 30.414l-4.293 4.293l-1.414-1.414L26.586 29l-6-6H18v7h-2V14Zm7 7h-5v-6h5a3 3 0 1 1 0 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}