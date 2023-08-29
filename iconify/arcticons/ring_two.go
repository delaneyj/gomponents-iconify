package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36 21.574v7.986a3.518 3.518 0 0 1-3.518 3.518h0a3.507 3.507 0 0 1-2.488-1.03"/><rect width="7.037" height="10.274" x="28.963" y="18.055" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.518" transform="rotate(180 32.482 23.192)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 21.574a3.518 3.518 0 0 1 3.518-3.519h0M12 21.574v6.755"/><circle cx="17.525" cy="15.645" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.525 18.055v10.274m9.253 0v-6.755a3.518 3.518 0 0 0-3.518-3.519h0a3.518 3.518 0 0 0-3.518 3.519v6.755"/>`),
		g.Group(children),
	)
}