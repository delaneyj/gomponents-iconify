package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertOctagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m272 0l112 112v160L272 384H112L0 272V112L112 0h160zm-80 305q12 0 20-8t8-19.5t-8-19.5t-20-8t-20 8t-8 19.5t8 19.5t20 8zm21-92V85h-42v128h42z"/>`),
		g.Group(children),
	)
}