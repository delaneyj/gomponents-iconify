package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DictionaryCom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.146 41.353V14.77l-13.682-2.622l8.176 4.885v24.32H5.5V6.647h19.646A17.354 17.354 0 0 1 42.5 24h0a17.354 17.354 0 0 1-17.354 17.353Z"/>`),
		g.Group(children),
	)
}