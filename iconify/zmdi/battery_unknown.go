package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryUnknown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M185 45q12 0 20 8.5t8 20.5v327q0 12-8 20t-20 8H28q-11 0-19.5-8T0 401V74q0-12 8.5-20.5T28 45h36V3h85v42h36zm-58 298v-41H86v41h41zm29-112q15-15 15-36q0-27-19-45.5T106.5 131t-45 18.5T43 195h32q0-14 9-23t22.5-9t23 9t9.5 22.5t-10 22.5l-20 20q-19 21-19 43h34q0-16 17-34z"/>`),
		g.Group(children),
	)
}