package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcecreamOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 4h14l3 28H14l3-28Zm4 28v9a3 3 0 0 0 3 3v0a3 3 0 0 0 3-3v-9"/><path d="M16 14v0a5.657 5.657 0 0 0 8 0v0a5.657 5.657 0 0 1 8 0v0m-17 8v0a6.042 6.042 0 0 0 8.76.716L24 22.5l.24-.216A6.042 6.042 0 0 1 33 23v0"/></g>`),
		g.Group(children),
	)
}