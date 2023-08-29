package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M42 4H12v40h30V4Z"/><path stroke-linecap="round" d="M4 11.79L12 10v34L4 11.79Z" clip-rule="evenodd"/><path d="m27 18l-5 6l5 6l5-6l-5-6Z"/><path stroke-linecap="round" d="M18 10v4m18 20v4"/></g>`),
		g.Group(children),
	)
}