package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeploymentUnitData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 23h-4V9h4a4.005 4.005 0 0 1 4 4v6a4.004 4.004 0 0 1-4 4Zm-2-2h2a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2h-2Z"/>`),
		g.Group(children),
	)
}