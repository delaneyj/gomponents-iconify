package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path stroke="#000" stroke-linejoin="round" d="M43 41H23"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M38.7183 5H21L5 41H22.662L38.7183 5Z"/><path stroke="#fff" stroke-linejoin="round" d="M9.95898 29.8823H17.9872"/><path stroke="#fff" stroke-linejoin="round" d="M13.2646 22.4707H21.2928"/><path stroke="#fff" stroke-linejoin="round" d="M16.7744 14.6H24.8026"/><path stroke="#000" d="M21 5L5 41"/></g>`),
		g.Group(children),
	)
}