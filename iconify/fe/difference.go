package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Difference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDifference0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDifference1" fill="currentColor"><path id="feDifference2" d="M15 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4Zm-4-6h4v4h4V5h-8v4Zm-6 2v8h8v-4H9v-4H5Z"/></g></g>`),
		g.Group(children),
	)
}