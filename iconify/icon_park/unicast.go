package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unicast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000"><path fill="#2F88FF" stroke-width="4" d="M24 15C26.7614 15 29 12.7614 29 10C29 7.23858 26.7614 5 24 5C21.2386 5 19 7.23858 19 10C19 12.7614 21.2386 15 24 15Z"/><path stroke-linecap="round" stroke-width="4" d="M24 15V23"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 33V23L24 23.0034"/><path stroke-linecap="round" stroke-width="5" d="M41 41V42"/><path stroke-linecap="round" stroke-width="5" d="M7 41V42"/><path stroke-linecap="round" stroke-width="5" d="M18 41V42"/><path stroke-linecap="round" stroke-width="5" d="M30 41V42"/></g>`),
		g.Group(children),
	)
}