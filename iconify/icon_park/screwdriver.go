package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screwdriver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M19 24H29V29C29.9613 29.9763 31.0387 31.5237 32 32.5V44H16V32.5L19 29V24Z"/><path stroke-linecap="round" d="M21 13V24H27V13L29 10L27 4H21L19 10L21 13Z"/></g>`),
		g.Group(children),
	)
}