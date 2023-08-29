package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 6V42"/><path fill="#2F88FF" d="M4 34L16 12V34H4Z"/><path fill="#2F88FF" d="M44 34H32V12L44 34Z"/></g>`),
		g.Group(children),
	)
}