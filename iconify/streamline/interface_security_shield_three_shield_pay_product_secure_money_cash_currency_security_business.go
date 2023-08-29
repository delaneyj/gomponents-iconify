package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSecurityShieldThreeShieldPayProductSecureMoneyCashCurrencySecurityBusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13 6.94a5.64 5.64 0 0 1-3.13 5L7 13.5L4.13 12A5.64 5.64 0 0 1 1 6.94V1.5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/><path d="M5.5 8A1.5 1.5 0 0 0 7 9.5A1.4 1.4 0 0 0 8.5 8c0-1.5-3-1.5-3-3A1.4 1.4 0 0 1 7 3.5A1.5 1.5 0 0 1 8.5 5M7 3.5v-1m0 8v-1"/></g>`),
		g.Group(children),
	)
}