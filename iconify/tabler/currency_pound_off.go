package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyPoundOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18.5a6 6 0 0 1-5 0a6 6 0 0 0-5 .5a3 3 0 0 0 2-2.5V9m1.192-2.825A4 4 0 0 1 16.45 7M13 13H7M3 3l18 18"/>`),
		g.Group(children),
	)
}