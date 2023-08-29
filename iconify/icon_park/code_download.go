package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M16 12L4 24.4322L16 36"/><path stroke-linejoin="round" d="M32 12L44 24.4322L32 36"/><path d="M24 17V31"/><path stroke-linejoin="round" d="M18 25L24 31L30 25"/></g>`),
		g.Group(children),
	)
}