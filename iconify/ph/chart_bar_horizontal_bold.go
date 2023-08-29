package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 92h-36V56a12 12 0 0 0-12-12H52v-4a12 12 0 0 0-24 0v176a12 12 0 0 0 24 0v-4h84a12 12 0 0 0 12-12v-36h68a12 12 0 0 0 12-12v-48a12 12 0 0 0-12-12Zm-60-24v24H52V68Zm-32 120H52v-24h72Zm80-48H52v-24h152Z"/>`),
		g.Group(children),
	)
}