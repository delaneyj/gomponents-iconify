package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWrapOverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M8 10V38"/><path d="M24 4V16"/><path d="M16 24H42"/><path stroke-linejoin="round" d="M37.0561 19.0113L42.0929 24.0255L37.0561 29.123"/><path d="M24 32V44"/></g>`),
		g.Group(children),
	)
}