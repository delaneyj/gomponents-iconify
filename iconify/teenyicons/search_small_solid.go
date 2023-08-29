package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 4a3 3 0 1 0 1.738 5.445l1.408 1.409l.708-.708l-1.409-1.408A3 3 0 0 0 7 4Z"/>`),
		g.Group(children),
	)
}