package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HboMax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m38.5 26.7l-5.8 7.7m5.8 0l-5.8-7.7M9.5 34.4v-4.6c0-1.6 1.3-2.9 2.9-2.9h0c1.6 0 2.9 1.3 2.9 2.9v4.6m0-4.6c0-1.6 1.3-2.9 2.9-2.9h0c1.6 0 2.9 1.3 2.9 2.9v4.6m8.6-2.9c0 1.6-1.3 2.9-2.9 2.9h0c-1.6 0-2.9-1.3-2.9-2.9c0 0 0 0 0 0v-1.9c0-1.6 1.3-2.9 2.9-2.9c0 0 0 0 0 0h0c1.6 0 2.9 1.3 2.9 2.9m0 4.8v-7.7M11.3 13.6v10.5m7-10.5v10.5m-7-5.3h7m12.6 1.8c0 1.9 1.6 3.5 3.5 3.4c1.9 0 3.4-1.5 3.4-3.4v-3.5c0-1.9-1.6-3.5-3.5-3.4c-1.9 0-3.4 1.5-3.4 3.4v3.5zm-5.3-1.8c1.4 0 2.6 1.2 2.6 2.6S27 24 25.6 24c0 0 0 0 0 0h-4.3V13.6h4.3c1.4 0 2.6 1.2 2.6 2.6s-1.1 2.6-2.6 2.6c0 0 0 0 0 0zm0 0h-4.3m13.1 1.3v-2.6"/><path d="M5.5 19.6v20.9c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2h-33c-1.1 0-2 .9-2 2v12.1"/></g>`),
		g.Group(children),
	)
}