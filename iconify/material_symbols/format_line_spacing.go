package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 20l-4-4l1.4-1.4L5 16.15v-8.3L3.4 9.4L2 8l4-4l4 4l-1.4 1.4L7 7.85v8.3l1.6-1.55L10 16l-4 4Zm6-1v-2h10v2H12Zm0-6v-2h10v2H12Zm0-6V5h10v2H12Z"/>`),
		g.Group(children),
	)
}