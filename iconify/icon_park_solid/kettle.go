package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kettle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSKettle0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M6.778 37h36v7h-36z"/><path d="m38.778 36l-2-26h-32l6.566 7.55a3 3 0 0 1 .727 2.206L10.778 36m16-18h-5m5 6h-5m5 6h-5m15-20h8v16h-6m-19-17V4h9v5"/></g></mask><path fill="currentColor" d="M0 0h49v48H0z" mask="url(#ipSKettle0)"/>`),
		g.Group(children),
	)
}