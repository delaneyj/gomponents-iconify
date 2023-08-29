package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiPay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.624 12.665l6.681 7.3L33.376 7.5M16.71 25.642V15.135m10.692 10.507V13.978m5.833 0v11.664m-6.495 9.376a2.34 2.34 0 0 1-2.333 2.333h0a2.34 2.34 0 0 1-2.333-2.333V33.5a2.34 2.34 0 0 1 2.333-2.333h0A2.34 2.34 0 0 1 26.74 33.5m0 3.851v-6.182m6.689 3.849v3.15a2.34 2.34 0 0 1-2.333 2.332h0a1.957 1.957 0 0 1-1.633-.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.429 31.169v3.849a2.34 2.34 0 0 1-2.333 2.333h0a2.34 2.34 0 0 1-2.333-2.333v-3.85m-14.192 6.239v-8.64h2.808a2.916 2.916 0 0 1 0 5.833h-2.808"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 43.5h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}