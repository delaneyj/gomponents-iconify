package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Credentials(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 22a4 4 0 1 0-4-4a4 4 0 0 0 4 4zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2zM14 6h4v2h-4z"/><path fill="currentColor" d="M24 2H8a2.002 2.002 0 0 0-2 2v24a2.002 2.002 0 0 0 2 2h16a2.003 2.003 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2Zm-4 26h-8v-2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1Zm2 0v-2a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3v2H8V4h16v24Z"/>`),
		g.Group(children),
	)
}