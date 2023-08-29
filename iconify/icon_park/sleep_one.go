package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleepOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M7 6H24.1429L7 24H25"/><path d="M29 15H41L29 29H41"/><path d="M15 32H24.0476L15 42H25"/></g>`),
		g.Group(children),
	)
}