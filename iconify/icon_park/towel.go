package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Towel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path fill="#2F88FF" d="M36 18H4V26H36V18Z"/><path d="M36 12V32C36 34.2091 34.2091 36 32 36H12M12 36H8C5.79086 36 4 34.2091 4 32V8C4 5.79086 5.79086 4 8 4H40C42.2091 4 44 5.79086 44 8V40C44 42.21 42.21 44 40 44H16C13.79 44 12 42.21 12 40V36Z"/></g>`),
		g.Group(children),
	)
}