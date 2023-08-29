package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmbianceAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.245 17.631V30.52m5.585 4.891V12.66m5.585 2.66v17.547M24 37.743V10.156M29.585 5.5v37m5.585-7.202V12.9m5.585 7.152v8.09"/>`),
		g.Group(children),
	)
}