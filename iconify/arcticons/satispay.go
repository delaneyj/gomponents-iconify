package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satispay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 37.8L3.2 24l13.9-13.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.3 24L11.8 4.5h13.5L44.8 24L25.3 43.5h0h-13.5L31.3 24z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.8 31.1L16.7 24l7.1-7.1m1.5-12.4h12.3l-6 6m0 27l6 6H25.3"/>`),
		g.Group(children),
	)
}