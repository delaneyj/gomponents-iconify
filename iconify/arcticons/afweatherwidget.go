package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Afweatherwidget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.484 10.504l3.572-4.036l3.547 4.083l5.584-2.146v5.731l5.603.597l-1.254 5.366m.811 1.106l3.217 2.766l-1.396 1.17m-6.981 9.953v4.352l-5.613-2.093l-3.51 4.179l-3.626-4.181l-5.515 2.102V33.84l-5.654-.735l1.357-5.714L4.5 23.965l4.149-3.498l-1.318-5.713l5.583-.671v-5.62l5.57 2.04"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.765 25.121c1.265-8.768-11.424-9.587-12.835-1.568c-3.072-2.057-6.934.216-6.208 3.8c-7.245.058-6.19 7.48-1.96 7.671h22.872c6.18.135 7.413-9.578-1.87-9.903l.236-.134"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 29.209a9.966 9.966 0 0 1 16.08-11.26"/>`),
		g.Group(children),
	)
}