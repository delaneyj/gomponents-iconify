package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Format(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="32" height="18" x="8" y="24" stroke-linejoin="round"/><path fill="#2F88FF" stroke-linejoin="round" d="M4 13H18V6H30V13H44V24H4V13Z"/><path d="M16 32L16 42"/></g>`),
		g.Group(children),
	)
}