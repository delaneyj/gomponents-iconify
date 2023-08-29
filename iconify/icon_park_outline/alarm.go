package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M14 25c0-5.523 4.477-10 10-10s10 4.477 10 10v16H14V25Z"/><path stroke-linecap="round" d="M24 5v3m11.892 1.328l-1.929 2.298m8.256 8.661l-2.955.521m-33.483-.521l2.955.521m3.373-11.48l1.928 2.298M6 41h37"/></g>`),
		g.Group(children),
	)
}