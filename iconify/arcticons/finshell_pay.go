package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinshellPay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.875 34.706V4.5h15.617c6.449 0 6.43 7.273 0 7.273h-7.874v6.88h7.94c5.929 0 5.927 7.338 0 7.338h-7.48V43.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.95 43.5V33.198c13.345.273 15.176-12.952 8.136-18.346"/>`),
		g.Group(children),
	)
}