package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpDesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 10h-6a3.003 3.003 0 0 0-3 3v6a2.002 2.002 0 0 0 2 2v7a2.002 2.002 0 0 0 2 2h4a2.002 2.002 0 0 0 2-2v-7a2.002 2.002 0 0 0 2-2v-6a3.003 3.003 0 0 0-3-3zm1 9h-2v9h-4v-9h-2v-6a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1zM20 5a4 4 0 1 1 4 4a4.004 4.004 0 0 1-4-4zm2 0a2 2 0 1 0 2-2a2.002 2.002 0 0 0-2 2zm-8 11v-3a3.003 3.003 0 0 0-3-3H5a3.003 3.003 0 0 0-3 3v3H0v2h16v-2zM4 13a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v3H4zm0-8a4 4 0 1 1 4 4a4.004 4.004 0 0 1-4-4zm2 0a2 2 0 1 0 2-2a2.002 2.002 0 0 0-2 2z"/>`),
		g.Group(children),
	)
}