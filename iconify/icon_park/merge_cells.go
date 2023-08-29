package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeCells(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M20 14V5C20 4.44772 19.5523 4 19 4H5C4.44772 4 4 4.44772 4 5V43C4 43.5523 4.44772 44 5 44H19C19.5523 44 20 43.5523 20 43V34"/><path d="M28 34V43C28 43.5523 28.4477 44 29 44H43C43.5523 44 44 43.5523 44 43V5C44 4.44772 43.5523 4 43 4H29C28.4477 4 28 4.44772 28 5V14"/><path d="M28 24H44"/><path d="M5 24H20"/><path stroke-linejoin="round" d="M32.7485 28.8184L31.1575 27.2274L27.9756 24.0454L31.1575 20.8634L32.7485 19.2724"/><path stroke-linejoin="round" d="M15.375 28.8184L16.966 27.2274L20.148 24.0454L16.966 20.8634L15.375 19.2724"/></g>`),
		g.Group(children),
	)
}