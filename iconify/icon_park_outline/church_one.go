package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M24 4v12m-4-8h8"/><path stroke-linejoin="round" d="M15 28H9a2 2 0 0 0-2 2v14m26-16h6a2 2 0 0 1 2 2v14"/><path stroke-linecap="round" d="M4 44h40"/><path stroke-linecap="round" stroke-linejoin="round" d="m15 23l9-8l9 8v21H15V23Z"/><path stroke-linecap="round" d="M24 34v10m-4 0h8"/></g>`),
		g.Group(children),
	)
}