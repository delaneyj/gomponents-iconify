package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatIndentIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h18v2H3Zm8-4v-2h10v2H11Zm0-4v-2h10v2H11Zm0-4V7h10v2H11ZM3 5V3h18v2H3Zm0 11V8l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}