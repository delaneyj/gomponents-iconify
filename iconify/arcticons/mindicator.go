package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mindicator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 18.54h39v11.048h-39Zm28.302 11.071a10.438 10.438 0 0 1-17.603 0m-.091-11.077a10.438 10.438 0 0 1 17.784 0m6.386 11.064a16.273 16.273 0 0 1-30.559 0m-.045-11.07a16.273 16.273 0 0 1 30.652.01"/><circle cx="24" cy="24" r="1.867" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.532 24h9.78m-3.792-2.19L37.312 24l-3.792 2.19M20.468 24h-9.78m3.792 2.19L10.688 24l3.792-2.19"/>`),
		g.Group(children),
	)
}