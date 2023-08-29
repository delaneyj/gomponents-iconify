package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TradingviewAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.735 20.11a9.475 9.475 0 0 1 18.445-1.46m.001-.001a6.497 6.497 0 0 1 6.057-.381a6.378 6.378 0 0 1 3.688 5.356"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.926 23.624a5.452 5.452 0 0 1-1.128 10.503m0 0H10.864m.871-14.017a7.024 7.024 0 0 0-.87 14.016m19.282-6.796l7.454-8.204m-16.787 2.717l6.5 5.7M6.08 31.606l11.716-9.83"/><circle cx="19.334" cy="20.5" r="2" fill="none" stroke="currentColor" stroke-miterlimit="10"/><circle cx="28.836" cy="28.836" r="2" fill="none" stroke="currentColor" stroke-miterlimit="10"/>`),
		g.Group(children),
	)
}