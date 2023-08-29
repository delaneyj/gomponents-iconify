package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcecreamOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 4H31L34 32H14L17 4Z"/><path d="M21 32V41C21 42.6569 22.3431 44 24 44V44C25.6569 44 27 42.6569 27 41V32"/><path d="M16 14V14C18.2091 16.2091 21.7909 16.2091 24 14V14V14C26.2091 11.7909 29.7909 11.7909 32 14V14"/><path d="M15 22V22C17.1706 24.7132 21.1768 25.0409 23.7594 22.7165L24 22.5L24.2406 22.2835C26.8232 19.9591 30.8294 20.2868 33 23V23"/></g>`),
		g.Group(children),
	)
}