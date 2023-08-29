package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageTemplate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M23 4H4V26H23V4Z"/><path d="M44 34H4V43H44V34Z"/><path d="M44 4H31V12H44V4Z"/><path d="M44 18H31V26H44V18Z"/></g>`),
		g.Group(children),
	)
}