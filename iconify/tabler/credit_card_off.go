package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18M9 5h9a3 3 0 0 1 3 3v8a3 3 0 0 1-.128.87m-2.002 2.002A3 3 0 0 1 18 19H6a3 3 0 0 1-3-3V8a3 3 0 0 1 2.124-2.87M3 11h8m4 0h6M7 15h.01M11 15h2"/>`),
		g.Group(children),
	)
}