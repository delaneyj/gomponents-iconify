package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M185 45q12 0 20 8.5t8 20.5v327q0 12-8 20t-20 8H28q-11 0-19.5-8T0 401V74q0-12 8.5-20.5T28 45h36V3h85v42h36zM85 387l86-160h-43V109L43 269h42v118z"/>`),
		g.Group(children),
	)
}