package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37 21V37M11 37V44H37V37M11 37H4V10L18 4H30L44 10V37H37M11 37V21"/><path d="M30 4L24 18M24 18L18 4M24 18V37V44"/><path d="M30 4L24 18"/><path d="M24 18L18 4"/><path d="M18 4L14 12L18 17.5L16 23L24 37"/><path d="M30 4L35 12L30 17.5L32 23L24 37"/></g>`),
		g.Group(children),
	)
}