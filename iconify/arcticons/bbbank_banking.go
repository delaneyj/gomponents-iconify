package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbbankBanking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.078 39.165a3.433 3.433 0 0 1-2.634 1.226h0a3.44 3.44 0 0 1-3.44-3.44v-3.38c0-1.9 1.54-3.44 3.44-3.44h0c1.063 0 2.014.482 2.645 1.24m-7.73 2.643h4.344m-4.344 2.518h4.344M20.373 9.946a2.219 2.219 0 0 1 0 4.437h-3.66V5.508h3.66a2.219 2.219 0 0 1 0 4.438Zm.071-.014h-3.661m-5.622.036a2.219 2.219 0 0 1 0 4.438H7.5V5.53h3.661a2.219 2.219 0 0 1 0 4.438Zm0 0H7.5m8.76 19.781v-6.656c9.695 0 17.584-7.889 17.584-17.585H40.5c0 13.367-10.874 24.24-24.24 24.24Z"/>`),
		g.Group(children),
	)
}