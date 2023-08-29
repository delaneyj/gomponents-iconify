package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsGate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.646 28.525a3.48 3.48 0 0 0 2.92 1.314h1.75a2.928 2.928 0 0 0 2.92-2.92h0A2.928 2.928 0 0 0 12.316 24H10.42a2.928 2.928 0 0 1-2.92-2.92h0a2.928 2.928 0 0 1 2.92-2.919h1.751a3.134 3.134 0 0 1 2.919 1.314m3.022 10.364V18.16l5.838 11.68l5.838-11.678V29.84m3.122-1.315a3.48 3.48 0 0 0 2.92 1.314h1.75a2.928 2.928 0 0 0 2.92-2.92h0A2.928 2.928 0 0 0 37.58 24h-1.897a2.928 2.928 0 0 1-2.92-2.92h0a2.928 2.928 0 0 1 2.92-2.919h1.752a3.134 3.134 0 0 1 2.919 1.314"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}