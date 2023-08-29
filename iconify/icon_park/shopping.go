package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shopping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M39 32H13L8 12H44L39 32Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M3 6H6.5L8 12M8 12L13 32H39L44 12H8Z"/><circle cx="13" cy="39" r="3" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><circle cx="39" cy="39" r="3" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/></g>`),
		g.Group(children),
	)
}