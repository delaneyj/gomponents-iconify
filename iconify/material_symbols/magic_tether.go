package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicTether(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.35 15.35l-2.1-2.125q1.475-1.475 3.463-2.35T12 10q2.3 0 4.288.863t3.462 2.362l-2.1 2.125q-1.1-1.1-2.55-1.725T12 13q-1.65 0-3.1.625T6.35 15.35ZM2.125 11.1L0 8.975q2.3-2.35 5.375-3.663T12 4q3.55 0 6.625 1.313T24 8.974L21.875 11.1q-1.925-1.925-4.45-3.013T12 7Q9.1 7 6.575 8.088T2.125 11.1ZM12 22l-1.25-2.75L8 18l2.75-1.25L12 14l1.25 2.75L16 18l-2.75 1.25L12 22Z"/>`),
		g.Group(children),
	)
}