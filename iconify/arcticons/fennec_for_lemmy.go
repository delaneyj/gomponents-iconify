package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FennecForLemmy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.11 24A8.11 8.11 0 0 1 24 32.11h0A8.11 8.11 0 0 1 15.89 24h0a8.11 8.11 0 0 1 16.22 0h0Zm8.567 0c0 9.21-7.467 16.677-16.677 16.677S7.323 33.21 7.323 24S14.79 7.323 24 7.323S40.677 14.79 40.677 24Zm4.823 0c0 11.874-9.626 21.5-21.5 21.5h0C12.126 45.5 2.5 35.874 2.5 24h0C2.5 12.126 12.126 2.5 24 2.5h0c11.874 0 21.5 9.626 21.5 21.5h0Z"/>`),
		g.Group(children),
	)
}