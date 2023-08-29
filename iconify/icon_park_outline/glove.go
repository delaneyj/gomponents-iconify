package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M8 42h27.833v-8S41 20.582 42 18c1-2.582-.5-5.335-4-5c-3.5.335-6.889 8.33-6.889 8.33S30.5 13 30 10.5S29 4 19.306 4C9.61 4 8 11.12 8 15v27Z"/>`),
		g.Group(children),
	)
}