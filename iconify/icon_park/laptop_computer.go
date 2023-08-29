package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="38" height="24" x="5" y="8" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M4 40L44 40"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M22 14L26 14"/></g>`),
		g.Group(children),
	)
}