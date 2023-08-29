package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesTurtleneck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37 23V37M11 37V44H37V37M11 37H4V23C4 20 6 16.5 9 14C12 11.5 18 10 18 10H30C30 10 36 11.5 39 14C42 16.5 44 20 44 23V37H37M11 37V23"/><path fill="#2F88FF" d="M30 10H18V4H30V10Z"/></g>`),
		g.Group(children),
	)
}