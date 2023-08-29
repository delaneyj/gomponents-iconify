package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardCar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDashboardCar0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6.572 37.428A21.904 21.904 0 0 1 2 24C2 11.85 11.85 2 24 2s22 9.85 22 22c0 5.056-1.705 9.713-4.572 13.428"/><path d="M12.303 31.697A13.935 13.935 0 0 1 10 24c0-7.732 6.268-14 14-14s14 6.268 14 14c0 2.843-.847 5.488-2.303 7.697"/><path fill="#555" fill-rule="evenodd" d="m24 30l16 16H8l16-16Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDashboardCar0)"/>`),
		g.Group(children),
	)
}