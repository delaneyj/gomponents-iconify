package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 14L34 4L30.25 7.75L26.5 11.5L19 19L11.5 26.5L7.75 30.25L4 34L14 44L44 14Z"/><path d="M30.25 7.75L7.75 30.25"/><path d="M9 29L13 33"/><path d="M14 24L20 30"/><path d="M19 19L23 23"/><path d="M24 14L30 20"/><path d="M29 9L33 13"/></g>`),
		g.Group(children),
	)
}