package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bedside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="12" x="4" y="18" fill="#2F88FF" stroke="#000"/><rect width="40" height="12" x="4" y="30" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M22 24H26"/><path stroke="#fff" d="M22 36H26"/><path stroke="#000" d="M8 42V46"/><path stroke="#000" d="M40 42V46"/><path stroke="#000" d="M24 18V10"/><path stroke="#000" d="M32 10H16"/></g>`),
		g.Group(children),
	)
}