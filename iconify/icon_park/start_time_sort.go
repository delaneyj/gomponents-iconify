package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StartTimeSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5V30H42V5"/><path stroke-linejoin="round" d="M30.0367 10.0001L24.9999 15.0143L30.0367 20.1118"/><path stroke-linejoin="round" d="M30 37L24 43L18 37"/><path stroke-linejoin="round" d="M24 30V43"/><path d="M20 9.00195V21.0001"/></g>`),
		g.Group(children),
	)
}