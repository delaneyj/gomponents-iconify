package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gemini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 6C4 6 15.5878 14 24 14C32.4122 14 44 6 44 6"/><path d="M4 42C4 42 15.5878 34 24 34C32.4122 34 44 42 44 42"/><path d="M15 12V36"/><path d="M33 12V36"/></g>`),
		g.Group(children),
	)
}