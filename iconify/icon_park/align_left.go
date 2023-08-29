package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="28" height="14" x="14.5" y="17" fill="#2F88FF"/><path stroke-linecap="round" d="M6.5 6V42"/></g>`),
		g.Group(children),
	)
}