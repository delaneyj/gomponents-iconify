package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallerId(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.26 12.482a24.151 24.151 0 0 1-.884-4.394C18.219 6.604 16.924 5.5 15.43 5.5H8.507c-1.78 0-3.152 1.538-2.995 3.311C7.091 26.677 21.322 40.91 39.19 42.488c1.773.157 3.31-1.21 3.31-2.99v-6.173c0-2.253-1.103-3.543-2.587-3.7a24.16 24.16 0 0 1-4.394-.884a4.913 4.913 0 0 0-4.88 1.247l-2.963 2.963A31.336 31.336 0 0 1 15.05 20.325l2.963-2.963a4.913 4.913 0 0 0 1.247-4.88ZM31.383 5.5v11.023m3.814 0V5.5h2.48a4.823 4.823 0 0 1 4.823 4.823V11.7a4.823 4.823 0 0 1-4.823 4.822h-2.48Z"/>`),
		g.Group(children),
	)
}