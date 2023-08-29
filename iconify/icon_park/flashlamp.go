package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashlamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#2F88FF" stroke="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M23 14L18 24H30L25 34"/></g>`),
		g.Group(children),
	)
}