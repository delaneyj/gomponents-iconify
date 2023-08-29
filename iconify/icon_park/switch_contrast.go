package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchContrast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30 10H40C41.8856 10 42.8284 10 43.4142 10.5858C44 11.1716 44 12.1144 44 14V34C44 35.8856 44 36.8284 43.4142 37.4142C42.8284 38 41.8856 38 40 38H30"/><path d="M18 10H8C6.11438 10 5.17157 10 4.58579 10.5858C4 11.1716 4 12.1144 4 14V34C4 35.8856 4 36.8284 4.58579 37.4142C5.17157 38 6.11438 38 8 38H18"/><path d="M24 6V42"/></g>`),
		g.Group(children),
	)
}