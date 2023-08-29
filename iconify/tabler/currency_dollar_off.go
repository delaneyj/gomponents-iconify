package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyDollarOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.7 8A3 3 0 0 0 14 6h-4M7.443 7.431A3 3 0 0 0 10 12h2m4.564 4.558A3 3 0 0 1 14 18h-4a3 3 0 0 1-2.7-2M12 3v3m0 12v3M3 3l18 18"/>`),
		g.Group(children),
	)
}