package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinGames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 24L24 6L42 24L24 42L6 24Z"/><path d="M24 29V42"/><path d="M19 24H6"/><path d="M24 19V6"/><path d="M29 24H42"/><path d="M29 24C29 26.7614 26.7614 29 24 29C21.2386 29 19 26.7614 19 24C19 21.2386 21.2386 19 24 19C26.7614 19 29 21.2386 29 24Z"/></g>`),
		g.Group(children),
	)
}