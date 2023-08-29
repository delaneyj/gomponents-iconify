package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scorpio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M16 12a6 6 0 0 0-12 0m24 0a6 6 0 0 0-12 0m24 22a6 6 0 0 1-12 0M4 12v18m12-18v18"/><path d="M28 12v22m12-11v11"/><path stroke-linejoin="round" d="m36 27l4-4l4 4"/></g>`),
		g.Group(children),
	)
}