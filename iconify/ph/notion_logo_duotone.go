package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotionLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 208h-40L64 48h40Z" opacity=".2"/><path d="M216 40h-48a8 8 0 0 0 0 16h16v120.85L111 44.14a8 8 0 0 0-7-4.14H40a8 8 0 0 0 0 16h16v144H40a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16H72V79.15l73 132.71a8 8 0 0 0 7 4.14h40a8 8 0 0 0 8-8V56h16a8 8 0 0 0 0-16Zm-59.27 160L77.53 56h21.74l79.2 144Z"/></g>`),
		g.Group(children),
	)
}