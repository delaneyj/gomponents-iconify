package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 21V10H10V21"/><rect width="40" height="16" x="4" y="21" fill="#2F88FF" stroke="#000" stroke-width="4" rx="2"/><rect width="4" height="4" x="14" y="27" fill="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 29H36"/></g>`),
		g.Group(children),
	)
}