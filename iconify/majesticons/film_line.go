package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v1h2V5H6zm3 0v6h6V5H9zm8 0v2h2V6a1 1 0 0 0-1-1h-1zm2 4h-2v2h2V9zm0 4h-2v2h2v-2zm0 4h-2v2h1a1 1 0 0 0 1-1v-1zm-4 2v-6H9v6h6zm-8 0v-2H5v1a1 1 0 0 0 1 1h1zm-2-4h2v-2H5v2zm0-4h2V9H5v2z"/></g>`),
		g.Group(children),
	)
}