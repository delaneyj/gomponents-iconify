package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosFaceRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 34V44H14"/><path d="M34 44H44V34"/><path d="M34 4H44V14"/><path d="M14 4H4V14"/><path d="M16 34C16 34 19 37 24 37C29 37 32 34 32 34"/><path d="M24 14V23C24 25 22 27 20 27H19"/><path d="M34 14V16"/><path d="M14 14V16"/></g>`),
		g.Group(children),
	)
}