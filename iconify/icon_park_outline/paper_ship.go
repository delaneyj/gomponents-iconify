package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperShip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="m4 24l8.571 18L24 29L4 24Zm40 0l-8.571 18L24 29l20-5ZM13 42h22L24 29L13 42Z"/><path d="M12 26L24 4l12 22"/></g>`),
		g.Group(children),
	)
}