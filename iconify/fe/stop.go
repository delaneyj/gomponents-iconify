package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feStop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStop1" fill="currentColor" fill-rule="nonzero"><path id="feStop2" d="M7 5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2Z"/></g></g>`),
		g.Group(children),
	)
}