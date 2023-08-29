package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kettle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6.778 37h36v7h-36zm32-1l-2-26h-32l6.566 7.55a3 3 0 0 1 .727 2.206L10.778 36m16-18h-5m5 6h-5m5 6h-5m15-20h8v16h-6m-19-17V4h9v5"/>`),
		g.Group(children),
	)
}