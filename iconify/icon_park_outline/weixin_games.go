package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinGames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 24L24 6l18 18l-18 18L6 24Zm18 5v13m-5-18H6m18-5V6m5 18h13"/><path d="M29 24a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/></g>`),
		g.Group(children),
	)
}