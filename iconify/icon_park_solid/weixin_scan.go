package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWeixinScan0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 30H12l2 5h12l2.088 6.265A4 4 0 0 0 31.883 44H38a4 4 0 0 0 4-4V30l1-7l-3.646.73a2 2 0 0 0-1.58 1.632L37 30ZM11 18h25l-2-5H22l-2.088-6.265A4 4 0 0 0 16.117 4H10a4 4 0 0 0-4 4v10l-1 7l3.646-.73a2 2 0 0 0 1.58-1.632L11 18Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWeixinScan0)"/>`),
		g.Group(children),
	)
}