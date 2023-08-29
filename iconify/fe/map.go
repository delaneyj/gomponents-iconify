package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMap0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMap1" fill="currentColor" fill-rule="nonzero"><path id="feMap2" d="m9 17l6 1.955V7.045L9 5v12ZM3 5l6-2l6 2l6-2v16l-6 2l-6-2l-6 2V5Z"/></g></g>`),
		g.Group(children),
	)
}