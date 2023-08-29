package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StampBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 224a12 12 0 0 1-12 12H40a12 12 0 0 1 0-24h176a12 12 0 0 1 12 12Zm0-80v32a20 20 0 0 1-20 20H48a20 20 0 0 1-20-20v-32a20 20 0 0 1 20-20h51.48L84.81 55.54A36 36 0 0 1 120 12h16a36 36 0 0 1 35.2 43.54L156.52 124H208a20 20 0 0 1 20 20Zm-104-20h8l15.74-73.49A12 12 0 0 0 136 36h-16a12 12 0 0 0-11.73 14.51Zm80 24H52v24h152Z"/>`),
		g.Group(children),
	)
}