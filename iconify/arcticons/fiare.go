package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fiare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 7.542v5.888l9.405 3.812L5.5 21.056v5.888l9.405 3.812L5.5 34.571v5.887l17.104-6.575v-6.247c-3.145-1.214-6.298-2.425-9.45-3.636l9.45-3.631v-6.248C16.91 11.924 11.2 9.733 5.5 7.541Zm37 0v5.888l-9.405 3.812l9.405 3.814v5.888l-9.405 3.812l9.405 3.815v5.887l-17.104-6.575v-6.247c3.145-1.214 6.298-2.425 9.45-3.636l-9.45-3.631v-6.248C31.09 11.924 36.8 9.733 42.5 7.541Z"/>`),
		g.Group(children),
	)
}