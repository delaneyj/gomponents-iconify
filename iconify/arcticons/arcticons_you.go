package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArcticonsYou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 10.84l-11.4 6.58v13.16L24 37.16l11.4-6.58V17.42L24 10.84Z"/><circle cx="8.07" cy="14.8" r="3.11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="5.61" r="3.11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.93" cy="14.8" r="3.11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.93" cy="33.2" r="3.11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="42.39" r="3.11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="8.07" cy="33.2" r="3.11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.6 30.58l-1.84 1.06M24 37.16v2.12m0-28.44V8.71m11.4 21.87l1.84 1.06M35.4 17.42l1.84-1.06M12.6 17.42l-1.84-1.06m16.963 2.02L24 24l-3.723-5.62M24 29.62V24"/>`),
		g.Group(children),
	)
}