package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 31h2v-2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2h2v-2a3.003 3.003 0 0 0-3-3h-6a3.003 3.003 0 0 0-3 3zm6-6a4 4 0 1 1 4-4a4.004 4.004 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.003 2.003 0 0 0-2-2zM2 31h2v-2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2h2v-2a3.003 3.003 0 0 0-3-3H5a3.003 3.003 0 0 0-3 3zm6-6a4 4 0 1 1 4-4a4.004 4.004 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2zm10-3h2v-2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2h2v-2a3.003 3.003 0 0 0-3-3h-6a3.003 3.003 0 0 0-3 3zm6-6a4 4 0 1 1 4-4a4.004 4.004 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2zM2 16h2v-2a1.001 1.001 0 0 1 1-1h6a1.001 1.001 0 0 1 1 1v2h2v-2a3.003 3.003 0 0 0-3-3H5a3.003 3.003 0 0 0-3 3zm6-6a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}