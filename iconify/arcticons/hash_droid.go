package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashDroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.506 28.309V40.04a1.947 1.947 0 0 0 1.947 1.947h33.1A1.947 1.947 0 0 0 42.5 40.04V28.309m-3.495-14.593l2.59-2.551a3.022 3.022 0 0 0-4.282-4.266l-.002.002L34.76 9.5a18.41 18.41 0 0 0-21.515 0l-2.57-2.57a3.174 3.174 0 0 0-2.17-.915a2.979 2.979 0 0 0-2.123.876a3.018 3.018 0 0 0 0 4.264l2.6 2.55a18.497 18.497 0 0 0-3.476 10.807v3.797H42.5v-3.797a18.497 18.497 0 0 0-3.495-10.796Zm-19.87 6.98a2.92 2.92 0 1 1-2.92-2.921a2.92 2.92 0 0 1 2.92 2.92Zm12.656 2.92a2.92 2.92 0 1 1 2.92-2.92a2.92 2.92 0 0 1-2.92 2.92Zm-10.438 7.532v8m5.3-8v8m-5.3-4.015h5.3"/>`),
		g.Group(children),
	)
}