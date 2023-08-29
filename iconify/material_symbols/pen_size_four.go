package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenSizeFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.725 18.275Q5 17.55 5 16.5t.725-1.775l9-9Q15.45 5 16.5 5t1.775.725Q19 6.45 19 7.5t-.725 1.775l-9 9Q8.55 19 7.5 19t-1.775-.725Z"/>`),
		g.Group(children),
	)
}