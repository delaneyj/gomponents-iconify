package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthProducts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M33 12H15L10 17.8428V38.0919L15 44H33L38 38.0919V17.8428L33 12Z"/><path stroke="#fff" d="M19 20H23.5455H29"/><path stroke="#000" d="M33 12V7C33 5.34315 31.6569 4 30 4H18C16.3431 4 15 5.34315 15 7V12"/><circle cx="24" cy="32" r="5" stroke="#fff"/></g>`),
		g.Group(children),
	)
}