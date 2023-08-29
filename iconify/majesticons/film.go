package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6zM5 6a1 1 0 0 1 1-1h1v2H5V6zm14 0v1h-2V5h1a1 1 0 0 1 1 1zm-2 5V9h2v2h-2zm0 4v-2h2v2h-2zm1 4h-1v-2h2v1a1 1 0 0 1-1 1zM5 18v-1h2v2H6a1 1 0 0 1-1-1zm2-3H5v-2h2v2zm0-4H5V9h2v2z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}