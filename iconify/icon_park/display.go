package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Display(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="20" x="6" y="6" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M14 13L22 13"/><path stroke="#fff" stroke-linecap="round" d="M14 19L34 19"/><path stroke="#000" stroke-linecap="round" d="M8 44L12.8889 38H34.6667L40 44"/><path stroke="#000" stroke-linecap="round" d="M24 26L24 44"/></g>`),
		g.Group(children),
	)
}