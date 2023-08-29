package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertColorsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-3.325 0-5.663-2.3T4 13.125q0-1.65.625-3.063T6.35 7.55l4.6-4.525q.225-.2.5-.312T12 2.6q.275 0 .55.113t.5.312l4.6 4.525q1.1 1.1 1.725 2.513T20 13.124q0 3.275-2.337 5.575T12 21Zm0-2V4.8L7.75 9q-.875.825-1.313 1.875T6 13.125q0 2.425 1.75 4.15T12 19Z"/>`),
		g.Group(children),
	)
}