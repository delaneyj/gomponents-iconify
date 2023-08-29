package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="14" height="28" x="17" y="6" fill="#2F88FF"/><path stroke-linecap="round" d="M42 42H6"/></g>`),
		g.Group(children),
	)
}