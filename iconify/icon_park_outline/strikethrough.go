package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5 24h38m-19 0c16 6 10 20 0 20s-12-8-12-8m24-24s-3-8-12-8s-12.564 7.6-8.39 14"/><path d="M12 36s4 8 12 8s12.564-7.6 8.39-14"/></g>`),
		g.Group(children),
	)
}