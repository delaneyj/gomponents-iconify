package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullDressLonguette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M18 18L14 10H34L30 18V24L40.5 44H7L18 24V18Z"/><path stroke="#000" d="M20 4V10"/><path stroke="#000" d="M28 4V10"/><path stroke="#fff" d="M18 21L30 21"/><path stroke="#000" d="M18 19V23"/><path stroke="#000" d="M30 19V23"/></g>`),
		g.Group(children),
	)
}