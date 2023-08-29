package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imbalance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 19V4M12 22l12-3l12-3m-8 14l8-14m8 14l-8-14M20 36l-8-14M4 36l8-14"/><path d="M12 44a8 8 0 0 0 8-8H4a8 8 0 0 0 8 8Zm24-6a8 8 0 0 0 8-8H28a8 8 0 0 0 8 8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}