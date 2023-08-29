package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StereoPerspective(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 44V14L14 4H44V34L34 44H4Z"/><path d="M34 14V44"/><path d="M14 4L14 34"/><path d="M4 14L34 14"/><path d="M44 4L34 14"/><path d="M4 44L14 34"/><path d="M14 34L44 34"/></g>`),
		g.Group(children),
	)
}