package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.617 2.076a1 1 0 0 1 1.09.217l6 6a1 1 0 0 1-.067 1.475L15.562 14l5.078 4.232a1 1 0 0 1 .067 1.475l-6 6A1 1 0 0 1 13 25v-8.865l-4.36 3.633a1 1 0 1 1-1.28-1.536L12.438 14L7.36 9.768a1 1 0 1 1 1.28-1.536L13 11.865V3a1 1 0 0 1 .617-.924ZM15 16.135v6.45l3.519-3.518L15 16.135Zm0-4.27l3.519-2.932L15 5.414v6.451Z"/>`),
		g.Group(children),
	)
}