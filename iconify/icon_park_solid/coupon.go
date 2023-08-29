package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coupon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCoupon0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 19.313V9h40v10.313a5.5 5.5 0 0 0-3.636 5.187A5.5 5.5 0 0 0 44 29.687V40H4V29.687A5.5 5.5 0 0 0 7.636 24.5A5.5 5.5 0 0 0 4 19.313Z"/><path stroke="#000" stroke-linecap="round" d="m19 16l5 5l5-5m-11 6h12m-12 6.167h12M24 22v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCoupon0)"/>`),
		g.Group(children),
	)
}