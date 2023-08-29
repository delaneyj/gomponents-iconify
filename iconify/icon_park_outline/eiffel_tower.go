package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EiffelTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M21 8c0 11-3.5 25-10 36M27 8c0 11 3.5 25 10 36M4 44h40"/><path d="M14 30h20m-17-9h14M20 8h8m-4-4v4"/><path stroke-linejoin="round" d="M18 44s.813-2.813 2-4c1-1 2-2 4-2s3 1 4 2c1.344 1.344 2 4 2 4"/></g>`),
		g.Group(children),
	)
}