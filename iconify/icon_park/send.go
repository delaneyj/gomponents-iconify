package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Send(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M43 5L29.7 43L22.1 25.9L5 18.3L43 5Z"/><path stroke-linecap="round" d="M43.0001 5L22.1001 25.9"/></g>`),
		g.Group(children),
	)
}