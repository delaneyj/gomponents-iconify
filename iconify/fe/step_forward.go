package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feStepForward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStepForward1" fill="currentColor" fill-rule="nonzero"><path id="feStepForward2" d="M13 9v10h7a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2h7V9H8.76C8.34 9 8 8.68 8 8.285a.68.68 0 0 1 .128-.396l3.24-4.57a.79.79 0 0 1 1.054-.199a.74.74 0 0 1 .21.198l3.24 4.57a.689.689 0 0 1-.21.992a.795.795 0 0 1-.422.12H13Zm7 4a1 1 0 0 1 0 2h-5v-2h5ZM8 13h1v2H4a1 1 0 0 1 0-2h4Z"/></g></g>`),
		g.Group(children),
	)
}