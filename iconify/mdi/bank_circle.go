package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m5 15H7v-2h10v2m-9-3v-3h2v3H8m3 0v-3h2v3h-2m3 0v-3h2v3h-2m3-4H7V8.5L12 6l5 2.5V10Z"/>`),
		g.Group(children),
	)
}