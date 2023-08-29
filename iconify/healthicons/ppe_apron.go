package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeApron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m20.342 6.621l-.292-.658l-.716.068A7 7 0 0 0 13 13v12a3 3 0 0 0 3 3h1v11.72l3.521 1.175a11 11 0 0 0 6.957 0L31 39.72V28h1a3 3 0 0 0 3-3V13a7 7 0 0 0-6.334-6.969l-.716-.067l-.292.657a4.001 4.001 0 0 1-7.316 0ZM19 23v2h10v-2H19Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}