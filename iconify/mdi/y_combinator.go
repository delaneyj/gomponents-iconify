package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YCombinator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M2 2h20v20H2V2m9.25 15.5h1.5v-4.44L16 7h-1.5L12 11.66L9.5 7H8l3.25 6.06v4.44z" fill="currentColor"/>`),
		g.Group(children),
	)
}