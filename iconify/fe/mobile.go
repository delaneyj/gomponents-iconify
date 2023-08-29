package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMobile0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMobile1" fill="currentColor" fill-rule="nonzero"><path id="feMobile2" d="M8 2h8a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm0 2v15h8V4H8Z"/></g></g>`),
		g.Group(children),
	)
}