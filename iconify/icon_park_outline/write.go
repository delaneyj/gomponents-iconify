package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Write(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M5.325 43.5h8.485l31.113-31.113l-8.486-8.485L5.325 35.015V43.5Z"/><path stroke-linecap="round" d="m27.952 12.387l8.485 8.485"/></g>`),
		g.Group(children),
	)
}