package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feWrench0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWrench1" fill="currentColor"><path id="feWrench2" d="M14 11.584V20a2 2 0 1 1-4 0v-8.416a5.001 5.001 0 0 1 2.391-9.569L10 6l4 2l2.215-3.691A5.001 5.001 0 0 1 14 11.584Z"/></g></g>`),
		g.Group(children),
	)
}