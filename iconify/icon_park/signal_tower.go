package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M17.0812 6.00005C10.9612 10.2853 9.47386 18.7205 13.7591 24.8405C18.0444 30.9605 26.4796 32.4479 32.5996 28.1626L17.0812 6.00005Z"/><path d="M22 31V42"/><path d="M13 24.5L6 42H42"/><path d="M40 6L25 17"/><path d="M17 6H40L33 27.5"/></g>`),
		g.Group(children),
	)
}