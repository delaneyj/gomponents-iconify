package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesSweater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37 19V37M11 37V44H37V37M11 37H4V12L9 7L16 4H32L39 7L44 12V37H37M11 37V19"/><path d="M24 21V44"/><path d="M9 7L24 21"/><path d="M16 4L24 21"/><path d="M32 4L24 21"/><path d="M39 7L24 21"/></g>`),
		g.Group(children),
	)
}