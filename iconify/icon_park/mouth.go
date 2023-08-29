package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24C4 24 10 15 14 15C18 15 22 17 24 17C26 17 30 15 34 15C38 15 44 24 44 24C44 24 34 34 24 34C14 34 4 24 4 24Z"/><path d="M4 24H44"/></g>`),
		g.Group(children),
	)
}