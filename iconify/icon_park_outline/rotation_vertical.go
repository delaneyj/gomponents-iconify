package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotationVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m38 28l-4-4l-4 4"/><path d="M33.168 16C31.625 8.936 28.1 4 24 4c-5.523 0-10 8.954-10 20s4.477 20 10 20s10-8.954 10-20"/></g>`),
		g.Group(children),
	)
}