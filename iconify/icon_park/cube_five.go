package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30 27L24 44"/><path d="M18 27L24 44"/><path d="M18 27L30 27"/><path d="M41 34L30 27"/><path d="M41 14L30 27"/><path d="M41 14L24 17"/><path d="M30 27L24 17"/><path d="M24 4V17"/><path d="M7 14L24 17"/><path d="M18 27L24 17"/><path d="M18 27L7 14"/><path d="M18 27L7 34"/><path d="M41.3207 14L24.0002 4L6.67969 14V34L24.0002 44L41.3207 34V14Z"/></g>`),
		g.Group(children),
	)
}