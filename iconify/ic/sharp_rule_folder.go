package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpRuleFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 6H12l-2-2H2v16h20V6zM7.83 16L5 13.17l1.41-1.41l1.41 1.41l3.54-3.54l1.41 1.41L7.83 16zm9.58-3L19 14.59L17.59 16L16 14.41L14.41 16L13 14.59L14.59 13L13 11.41L14.41 10L16 11.59L17.59 10L19 11.41L17.41 13z"/>`),
		g.Group(children),
	)
}