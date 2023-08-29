package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbbankSecurego(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.307 41.022c-6.083-1.358-5.309-7.723-5.309-10.064c.889.123 3.986.266 5.309-2.173c1.323 2.44 4.42 2.296 5.308 2.173c0 2.341.774 8.706-5.308 10.064ZM20.373 9.946a2.219 2.219 0 0 1 0 4.437h-3.66V5.508h3.66a2.219 2.219 0 0 1 0 4.438Zm.071-.014h-3.661m-5.622.036a2.219 2.219 0 0 1 0 4.438H7.5V5.53h3.661a2.219 2.219 0 0 1 0 4.438Zm0 0H7.5m8.76 19.781v-6.656c9.695 0 17.584-7.889 17.584-17.585H40.5c0 13.367-10.874 24.24-24.24 24.24Z"/>`),
		g.Group(children),
	)
}