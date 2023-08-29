package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFilter0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFilter1" fill="currentColor" fill-rule="nonzero"><path id="feFilter2" d="M18 5.97V4H6v1.97l6 4.286l6-4.285ZM13 12v8l-2 2V12L4.838 7.598A2 2 0 0 1 4 5.971V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v1.97a2 2 0 0 1-.838 1.628L13 12Z"/></g></g>`),
		g.Group(children),
	)
}