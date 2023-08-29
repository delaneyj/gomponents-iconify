package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterRate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 44C32.8366 44 40 37.4876 40 29.4545C40 21.5758 34.6667 13.0909 24 4C13.3333 13.0909 8 21.5758 8 29.4545C8 37.4878 15.1634 44 24 44Z"/><path stroke="#fff" stroke-linecap="round" d="M18.8572 19L24 24.3684L29.1429 19"/><path stroke="#fff" stroke-linecap="round" d="M18 26.158H30"/><path stroke="#fff" stroke-linecap="round" d="M18 31.5264H30"/><path stroke="#fff" stroke-linecap="round" d="M24 26.158V36.0001"/></g>`),
		g.Group(children),
	)
}