package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChristmasTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M20 14L16 12L24 4L32 12L28 14L36 22L30 24L39 34H9L17 24L12 22L20 14Z"/><path d="M31 44H17"/><path d="M21 34L20 44"/><path d="M27 34L28 44"/></g>`),
		g.Group(children),
	)
}