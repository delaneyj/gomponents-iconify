package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.58 0H4.322C3.592 0 3 .598 3 1.334v13.333C3 15.404 3.592 16 4.322 16h9.258c.729 0 1.42-.596 1.42-1.333V1.334C15 .598 14.31 0 13.58 0zM7.021 14H4.987v-1h2.034v1zm0-5.979H4.987V7h2.034v1.021zM10 14H7.986v-1H10v1zm-2.979-3H4.987v-1h2.034v1zM10 11H7.986v-1H10v1zm0-3H7.986V7H10v1zm3 6h-2V9.979h2V14zm0-5.98h-2.014V7H13v1.02zM14.014 6H4V2h10.014v4z"/>`),
		g.Group(children),
	)
}