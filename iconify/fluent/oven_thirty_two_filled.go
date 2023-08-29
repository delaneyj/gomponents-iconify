package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvenThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7.5A4.5 4.5 0 0 1 7.5 3h17A4.5 4.5 0 0 1 29 7.5V12H3V7.5Zm5.5.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Zm6 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0ZM22 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM3 14v10.5A4.5 4.5 0 0 0 7.5 29h17a4.5 4.5 0 0 0 4.5-4.5V14H3Zm5.5 3h15a1.5 1.5 0 0 1 1.5 1.5v5a1.5 1.5 0 0 1-1.5 1.5h-15A1.5 1.5 0 0 1 7 23.5v-5A1.5 1.5 0 0 1 8.5 17Z"/>`),
		g.Group(children),
	)
}