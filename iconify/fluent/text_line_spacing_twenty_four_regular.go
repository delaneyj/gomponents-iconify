package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextLineSpacingTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m19.53 3.22l2 2a.75.75 0 0 1-1.06 1.06l-.72-.72v4.69a.75.75 0 0 1-1.5 0V5.56l-.72.72a.75.75 0 1 1-1.06-1.06l2-2a.748.748 0 0 1 .528-.22h.004a.748.748 0 0 1 .528.22ZM2 5.75A.75.75 0 0 1 2.75 5h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 2 5.75Zm0 6.5a.75.75 0 0 1 .75-.75h11.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75ZM2.75 18a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5Zm15.5.44v-4.69a.75.75 0 0 1 1.5 0v4.69l.72-.72a.75.75 0 1 1 1.06 1.06l-2 2a.748.748 0 0 1-1.06 0l-2-2a.75.75 0 1 1 1.06-1.06l.72.72Z"/>`),
		g.Group(children),
	)
}