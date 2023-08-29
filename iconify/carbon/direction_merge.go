package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17.878 15.465A2.982 2.982 0 0 1 17 13.343V7.828l4.586 4.586L23 11l-7-7l-7 7l1.414 1.414L15 7.828v5.515a2.98 2.98 0 0 1-.878 2.121l-2.658 2.657A4.968 4.968 0 0 0 10 21.657V28h2v-6.343a2.98 2.98 0 0 1 .878-2.121l2.658-2.657a5.002 5.002 0 0 0 .464-.566a5.002 5.002 0 0 0 .464.566l2.658 2.656A2.982 2.982 0 0 1 20 21.657V28h2v-6.343a4.968 4.968 0 0 0-1.464-3.536Z"/>`),
		g.Group(children),
	)
}