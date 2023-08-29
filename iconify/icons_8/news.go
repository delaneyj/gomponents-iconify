package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func News(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 5v18c0 2.21 1.79 4 4 4h18c2.21 0 4-1.79 4-4V12h-6V5H3zm2 2h16v16c0 .73.22 1.41.563 2H7c-1.19 0-2-.81-2-2V7zm2 2v5h12V9H7zm2 2h8v1H9v-1zm14 3h4v9c0 1.19-.81 2-2 2s-2-.81-2-2v-9zM7 15v2h5v-2H7zm7 0v2h5v-2h-5zm-7 3v2h5v-2H7zm7 0v2h5v-2h-5zm-7 3v2h5v-2H7zm7 0v2h5v-2h-5z"/>`),
		g.Group(children),
	)
}