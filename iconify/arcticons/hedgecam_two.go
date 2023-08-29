package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HedgecamTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.303 32.614a2.758 2.758 0 0 1 2.687-2.79a2.776 2.776 0 0 1 1.963 4.753c-1.137.93-4.65 3.616-4.65 3.616h5.477"/><circle cx="24.037" cy="26.295" r="8.126" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.402 11.623h6.171l3.937-4.55h11.499l3.344 4.55h6.245a2.896 2.896 0 0 1 2.902 2.903v23.5a2.896 2.896 0 0 1-2.902 2.902H8.402A2.896 2.896 0 0 1 5.5 38.025v-23.5a2.896 2.896 0 0 1 2.902-2.902Z"/>`),
		g.Group(children),
	)
}