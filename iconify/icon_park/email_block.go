package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M44 35V9H24H4V23V37H26"/><circle cx="35" cy="35" r="9" fill="#2F88FF" stroke="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M37 33L33 37"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M4 9L24 22L44 9"/></g>`),
		g.Group(children),
	)
}