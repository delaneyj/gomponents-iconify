package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkInterrupt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M27 14H42C43.1046 14 44 14.8954 44 16V30C44 31.1046 43.1046 32 42 32H38"/><path d="M11 14H6C4.89543 14 4 14.8954 4 16V30C4 31.1046 4.89543 32 6 32H21"/><path d="M14 6L34 40"/><path d="M32 23H36"/><path d="M12 23H16"/></g>`),
		g.Group(children),
	)
}