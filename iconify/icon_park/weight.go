package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M41 4H7C5.34315 4 4 5.34315 4 7V41C4 42.6569 5.34315 44 7 44H41C42.6569 44 44 42.6569 44 41V7C44 5.34315 42.6569 4 41 4Z"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M12 19.0537C15.3249 15.0537 19.3249 13.0537 24 13.0537C28.6751 13.0537 32.6751 15.0537 36 19.0537"/><path fill="#fff" d="M24 31C25.6569 31 27 29.6569 27 28C27 26.3431 25.6569 25 24 25C22.3431 25 21 26.3431 21 28C21 29.6569 22.3431 31 24 31Z"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M19 21L24.0083 28"/></g>`),
		g.Group(children),
	)
}