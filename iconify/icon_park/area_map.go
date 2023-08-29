package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M18 24L12 33L4 28.5V42H44V15L37 23L31 18L24 26L18 24Z"/><path d="M4 28.5V17L11 23L16.5 15L22.5 18L31 9L36.5 13.5L44 6V15.5"/></g>`),
		g.Group(children),
	)
}