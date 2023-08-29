package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSimulation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26 12l3.464 2l-1 1.732l-3.464-2zm0-4h4v2h-4zM7 13.732l-3.464 2l-1-1.732L6 12zM26 30h-2v-5a5.006 5.006 0 0 0-5-5h-6a5.006 5.006 0 0 0-5 5v5H6v-5a7.008 7.008 0 0 1 7-7h6a7.008 7.008 0 0 1 7 7zM16 4a5 5 0 1 1-5 5a5 5 0 0 1 5-5m0-2a7 7 0 1 0 7 7a7 7 0 0 0-7-7zm9 2.268l3.464-2l1 1.732L26 6zM2 8h4v2H2zm4-2L2.536 4l1-1.732l3.464 2z"/>`),
		g.Group(children),
	)
}