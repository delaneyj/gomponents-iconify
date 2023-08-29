package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DateComesBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 8H9a3 3 0 0 0-3 3v28a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V18"/><path d="m32.243 12.485l4.242-4.242L32.243 4M24 17v17m8-10v10M16 24v10"/></g>`),
		g.Group(children),
	)
}