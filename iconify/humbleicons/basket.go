package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 10l1.36 8.164a1 1 0 0 0 .987.836h9.306a1 1 0 0 0 .986-.836L19 10M5 10h3m-3 0H4m15 0h-3m3 0h1M8 10l1-5m-1 5h8m0 0l-1-5m-6 9v1m6-1v1m-3-1v1"/>`),
		g.Group(children),
	)
}