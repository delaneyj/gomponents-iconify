package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BacklightHigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 15v-2h4v2H1Zm5.35-5.25L3.525 6.925l1.4-1.425L7.75 8.35l-1.4 1.4ZM7 18v-3h10v3H7Zm4-11V2h2v5h-2Zm6.65 2.75l-1.4-1.4l2.825-2.825l1.425 1.4l-2.85 2.825ZM19 15v-2h4v2h-4Z"/>`),
		g.Group(children),
	)
}