package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vianavigo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.1 37.5c.858-1.407 7.93-11.123 10.299-11.123c3.604 0 9.112 8.446 10.985 11.123M13.272 26.257c1.922-1.682 6.626-3.673 11.02-4.291s9.646-3.021 10.608-4.428M24.566 10.5c1.923 0 2.781 1.51 2.781 3.261s-1.51 4.394-3.605 4.394a2.849 2.849 0 0 1-2.849-3.124c0-1.407 1.099-4.531 3.673-4.531Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}