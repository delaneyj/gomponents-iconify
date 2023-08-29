package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SailboatOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M21 31V5L7 31H21Z"/><path fill="#2F88FF" d="M27 31V13L41 31H27Z"/><path d="M5 37H43L38 43H10L5 37Z"/></g>`),
		g.Group(children),
	)
}