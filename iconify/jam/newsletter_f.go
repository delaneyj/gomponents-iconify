package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewsletterF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 .565h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm3 8a1 1 0 1 0 0 2h6a1 1 0 0 0 0-2H5zm10-5v2h2v-2h-2z"/>`),
		g.Group(children),
	)
}