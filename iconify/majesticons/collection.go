package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 21a2 2 0 0 1-2-2h12a2 2 0 0 1-2 2H8zm-4-5a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2H4zm0-1a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}