package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="26" x="8" y="6" stroke="#000" stroke-width="4" rx="2"/><circle cx="14" cy="27" r="2" fill="#000"/><circle cx="34" cy="27" r="2" fill="#000"/><rect width="20" height="10" x="14" y="12" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 32L40 41"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 32L8 41"/></g>`),
		g.Group(children),
	)
}