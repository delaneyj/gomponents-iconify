package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NailPolishOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="24" x="8" y="20" fill="#2F88FF" stroke="#000" rx="2"/><rect width="14" height="16" x="17" y="4" fill="#2F88FF" stroke="#000"/><path fill="#43CCF8" stroke="#fff" d="M22 32H26L27 37H21L22 32Z"/><path stroke="#fff" d="M24 20V32"/><path stroke="#000" d="M31 20H17"/></g>`),
		g.Group(children),
	)
}