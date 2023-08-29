package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTextTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 36.3056H42"/><path d="M6 42H42"/><path d="M30 12L24 6L18 12V12"/><path d="M24 6V29"/></g>`),
		g.Group(children),
	)
}