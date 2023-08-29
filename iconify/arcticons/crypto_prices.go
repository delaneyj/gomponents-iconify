package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CryptoPrices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.39 42.5l5.421-11.353l1.377 3.084l3.475-9.627l1.745 2.304l2.514-5.563l1.128 1.43L42.5 9.133m-28.228 4.949h-1.717"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M14.272 14.082v4.846h3.199a2.423 2.423 0 1 0 0-4.846Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M14.272 18.928v5.165h3.41a2.583 2.583 0 1 0 0-5.165Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.272 24.09h-1.717m1.713.032v1.589m2.779-1.589v1.589m-2.779-13.379v1.59m2.779-1.59v1.59"/>`),
		g.Group(children),
	)
}