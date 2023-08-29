package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropperSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.768 1.06a2.5 2.5 0 0 0-3.536 0L7.5 3.794l-.646-.647l-.708.708l5 5l.708-.708l-.647-.646l2.732-2.732a2.5 2.5 0 0 0 0-3.536l-.171-.171ZM6.146 6.146a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708L2.707 15H.5a.5.5 0 0 1-.5-.5v-2.207l6.146-6.147Z"/>`),
		g.Group(children),
	)
}