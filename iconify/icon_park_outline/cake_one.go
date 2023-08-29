package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M27 14L9 21.9h30L34 15"/><circle cx="31" cy="13" r="4" fill="currentColor"/><path stroke-linecap="round" d="m33 9l2-5"/><path d="M9.5 26.957c-.602.3-1.162.62-1.678.956C5.418 29.481 4 31.412 4 33.5c0 5.247 8.954 9.5 20 9.5s20-4.253 20-9.5c0-2.14-1.488-4.113-4-5.701"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 22h30v12H9z"/><path d="M9 22h31"/></g>`),
		g.Group(children),
	)
}