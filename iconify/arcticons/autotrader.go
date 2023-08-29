package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autotrader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.065 21.473l4.392-11.112h-24.49a8.421 8.421 0 0 0-7.831 5.326l-2.288 5.786h30.217ZM7.79 26.677L3.458 37.64h24.49a8.422 8.422 0 0 0 7.832-5.326l2.228-5.636H7.791Z"/>`),
		g.Group(children),
	)
}