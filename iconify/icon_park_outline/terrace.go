package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terrace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M5 24v16a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V24m0 7H5"/><path d="M32 23a8 8 0 1 0-16 0"/><path stroke-linejoin="round" d="M24 6v2m11.414 2L34 11.414M12 10l1.414 1.414"/></g>`),
		g.Group(children),
	)
}