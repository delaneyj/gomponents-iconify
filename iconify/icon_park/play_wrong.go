package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayWrong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 44C12.9543 44 4 35.0457 4 24C4 12.9543 12.9543 4 24 4C35.0457 4 44 12.9543 44 24"/><path stroke-linecap="round" d="M33 33L41 41"/><path stroke-linecap="round" d="M41 33L33 41"/><path fill="#2F88FF" d="M20 24V17.0718L26 20.5359L32 24L26 27.4641L20 30.9282V24Z"/></g>`),
		g.Group(children),
	)
}