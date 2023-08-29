package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21V5m2-2h10a2 2 0 0 1 2 2v10m0 4.01V21l-3-2l-2 2l-2-2l-2 2l-2-2l-3 2m6-14h4m-6 4h2m2 4h2m0-4v.01M3 3l18 18"/>`),
		g.Group(children),
	)
}