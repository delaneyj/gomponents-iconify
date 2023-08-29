package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38 6V44"/><path d="M10 40H38"/><path d="M10 34H38"/><path d="M6 12H42"/><path d="M35 6L41 6"/><path d="M7 6L13 6"/><path d="M10 6V44"/><path d="M17 12V28"/><path d="M24 12V28"/><path d="M31 12V28"/></g>`),
		g.Group(children),
	)
}