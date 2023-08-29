package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOneSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9.854.146a.5.5 0 0 0-.708 0L6.5 2.793L12.207 8.5l2.647-2.646a.5.5 0 0 0 0-.708l-5-5ZM0 9.293L5.793 3.5L11.5 9.207L5.707 15H.5a.5.5 0 0 1-.5-.5V9.293ZM8 15h7v-1H8v1Z"/>`),
		g.Group(children),
	)
}