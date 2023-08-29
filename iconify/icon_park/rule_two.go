package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M24 13C11 13 4 22.8497 4 35H44C44 22.8497 37 13 24 13Z"/><path stroke-linecap="round" d="M10 31L10 35"/><path stroke-linecap="round" d="M17 31L17 35"/><path stroke-linecap="round" d="M24 31L24 35"/><path stroke-linecap="round" d="M31 31L31 35"/><path stroke-linecap="round" d="M38 31L38 35"/><path d="M24 20C19.4457 20 14 22.134 14 26H34C34 22.134 28.5543 20 24 20Z"/></g>`),
		g.Group(children),
	)
}