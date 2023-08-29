package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YCombinator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 2v28H2V2zM16.75 17.969l4.844-9.094H19.55l-2.863 5.688a38.406 38.406 0 0 0-.8 1.675l-.762-1.675L12.3 8.875h-2.188l4.794 8.988v5.906h1.844z"/>`),
		g.Group(children),
	)
}