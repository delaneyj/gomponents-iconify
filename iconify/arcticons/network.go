package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Network(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.259 12.482a24.15 24.15 0 0 1-.883-4.394C18.218 6.604 16.924 5.5 15.43 5.5H8.507c-1.78 0-3.152 1.538-2.995 3.311C7.091 26.677 21.322 40.91 39.19 42.488c1.773.157 3.31-1.21 3.31-2.99v-6.173c0-2.253-1.103-3.543-2.587-3.7a24.142 24.142 0 0 1-4.394-.884a4.913 4.913 0 0 0-4.88 1.247l-2.963 2.963A31.337 31.337 0 0 1 15.05 20.325l2.963-2.963a4.913 4.913 0 0 0 1.247-4.88h0Z"/><circle cx="35.483" cy="12.515" r="7.015" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.483" cy="8.732" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.483 11.187v5.311"/>`),
		g.Group(children),
	)
}