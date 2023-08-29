package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tresorit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.75 34.25v-20.5L24 3.5L6.25 13.75v20.5L24 44.5l17.75-10.25zm-35.5-8.24L34.61 9.63"/>`),
		g.Group(children),
	)
}