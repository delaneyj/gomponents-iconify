package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6 5.621a1.621 1.621 0 0 1 2.768-1.146l.378.379l.708-.708l-.38-.378A2.621 2.621 0 0 0 5 5.62v.464A1.5 1.5 0 0 0 4 7.5v3A1.5 1.5 0 0 0 5.5 12h4a1.5 1.5 0 0 0 1.5-1.5v-3A1.5 1.5 0 0 0 9.5 6H6v-.379Z"/>`),
		g.Group(children),
	)
}