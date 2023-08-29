package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassesOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="14.5" cy="24.5" r="6.5"/><circle r="6.5" transform="matrix(-1 0 0 1 33.5 24.5)"/><path d="M4 24h4m36 0h-4m-20-3c.5-2 2-4 4-4s3.5 2 4 4"/></g>`),
		g.Group(children),
	)
}