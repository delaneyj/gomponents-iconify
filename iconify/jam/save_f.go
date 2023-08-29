package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 18.565v-7a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v7H2a2 2 0 0 1-2-2v-14a2 2 0 0 1 2-2h11.22a2 2 0 0 1 1.345.52l2.78 2.527A2 2 0 0 1 18 5.092v11.473a2 2 0 0 1-2 2h-2zm-1-15a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1zm-8 0a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1H5zm1 9h6v4H6v-4zm0-5v-3h3v3H6z"/>`),
		g.Group(children),
	)
}