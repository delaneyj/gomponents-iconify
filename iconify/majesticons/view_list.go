package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M4 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H4zm1 2a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2H5zm5 0a1 1 0 0 0 0 2h9a1 1 0 1 0 0-2h-9zm-5 4a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H5zm5 0a1 1 0 1 0 0 2h9a1 1 0 1 0 0-2h-9zm-5 4a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H5zm5 0a1 1 0 1 0 0 2h9a1 1 0 1 0 0-2h-9z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}