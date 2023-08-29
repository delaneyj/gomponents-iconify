package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.815 14.632l-4.031 4.031H7.115v2.668H4.447v2.668H0v-4.447l9.368-9.368a7.41 7.41 0 0 1-.474-2.632a7.554 7.554 0 1 1 4.869 7.062l.052.017zm7.532-9.31v-.003a2.668 2.668 0 1 0-2.669 2.668h.001a2.669 2.669 0 0 0 2.668-2.665z"/>`),
		g.Group(children),
	)
}