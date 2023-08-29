package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermosCup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M10 31C10 31 10.0714 37 24 37C37.9286 37 38 31 38 31V15H24H10V31Z"/><path stroke-linecap="round" d="M24 4V10"/><path stroke-linecap="round" d="M16 5V9"/><path stroke-linecap="round" d="M32 5V9"/><path stroke-linecap="round" d="M14 36V41C14 42.1046 14.8954 43 16 43H32C33.1046 43 34 42.1046 34 41V36"/></g>`),
		g.Group(children),
	)
}