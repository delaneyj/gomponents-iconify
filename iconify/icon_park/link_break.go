package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 15H42C43.1046 15 44 15.8954 44 17V31C44 32.1046 43.1046 33 42 33H32"/><path d="M17 15H6C4.89543 15 4 15.8954 4 17V31C4 32.1046 4.89543 33 6 33H17"/><path d="M24 6V42"/><path d="M12 24H16"/><path d="M32 24H36"/></g>`),
		g.Group(children),
	)
}