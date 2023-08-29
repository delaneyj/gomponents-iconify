package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeadlineSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5V30.0036H42V5"/><path stroke-linejoin="round" d="M30 37L24 43L18 37"/><path stroke-linejoin="round" d="M24 30V43"/><path d="M18.3442 20.6577L29.6579 9.34401"/><path d="M18.3438 9.34326L29.6575 20.657"/></g>`),
		g.Group(children),
	)
}