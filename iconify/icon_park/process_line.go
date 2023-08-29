package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProcessLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 40H14"/><path d="M4 32H24"/><path d="M22 40H27"/><path d="M37 40H44"/><path d="M35 32H44"/><path d="M32 24H44"/><path d="M16 24H22"/><path d="M4 24H6"/><path d="M4 16H8"/><path d="M4 8H11L19 16H44"/><path d="M22 8H44"/></g>`),
		g.Group(children),
	)
}