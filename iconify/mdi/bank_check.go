package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.8 21.2l-2.8-3l1.2-1.2l1.6 1.6l3.6-3.6l1.2 1.4l-4.8 4.8M13 10h-3v7h2.1c.1-.8.5-1.6.9-2.3V10m3 0v2.3c.6-.2 1.3-.3 2-.3c.3 0 .7 0 1 .1V10h-3m-3.9 9H2v3h11.5c-.7-.8-1.2-1.9-1.4-3M21 6l-9.5-5L2 6v2h19V6M7 17v-7H4v7h3Z"/>`),
		g.Group(children),
	)
}