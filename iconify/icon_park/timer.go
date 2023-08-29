package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><circle cx="24" cy="28" r="16" fill="#2F88FF" stroke="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 4L20 4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 4V12"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M35 16L38 13"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M24 28V22"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M24 28H18"/></g>`),
		g.Group(children),
	)
}