package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinningTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M24 44V41"/><path fill="#2F88FF" d="M44 20L24 41L4 20H44Z"/><path d="M44 12H4V20H44V12Z"/><path d="M24 12V4"/></g>`),
		g.Group(children),
	)
}