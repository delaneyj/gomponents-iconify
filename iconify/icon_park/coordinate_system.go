package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoordinateSystem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 12L38 20V36L24 44L10 36V20L24 12Z"/><path stroke="#000" stroke-linecap="round" d="M24 6V12"/><path stroke="#fff" stroke-linecap="round" d="M10 20L24 28L38 20"/><path stroke="#000" stroke-linecap="round" d="M38 36L44 39"/><path stroke="#000" stroke-linecap="round" d="M4 39L10 36"/><path stroke="#fff" stroke-linecap="round" d="M24 28V44"/><path stroke="#000" d="M31 16L38 20V28M17 16L10 20V28M17 40L24 44L31 40"/></g>`),
		g.Group(children),
	)
}