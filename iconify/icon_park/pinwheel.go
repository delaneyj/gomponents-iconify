package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinwheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M21 4V21H11L21 4Z"/><path d="M27 44V27H37L27 44Z"/><path d="M27 11L44 21H27V11Z"/><path d="M21 37L4 27H21V37Z"/></g>`),
		g.Group(children),
	)
}