package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bybit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.096 17.056v9.258m2.271-6.943H40.5m-3.067 9.258v-9.258m-16.143 0L18.224 24l-3.067-4.629m3.067 9.258V24m-6.905 0a2.314 2.314 0 1 1 0 4.629H7.5v-9.258h3.819a2.314 2.314 0 1 1 0 4.629Zm0 0H7.5m19.88 0a2.314 2.314 0 1 1 0 4.629h-3.819v-9.258h3.819a2.314 2.314 0 1 1 0 4.629Zm0 0h-3.819"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}