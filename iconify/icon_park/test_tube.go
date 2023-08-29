package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M18 4H30"/><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M24 44C27.3137 44 30 41.3137 30 38V10H18V38C18 41.3137 20.6863 44 24 44Z" clip-rule="evenodd"/><path stroke="#fff" d="M24 27V28"/><path stroke="#fff" d="M24 18V21"/><path stroke="#000" d="M19 35H30"/></g>`),
		g.Group(children),
	)
}