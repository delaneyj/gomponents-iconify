package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeploymentUnitPresentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 23h-2V9h6a2.002 2.002 0 0 1 2 2v5a2.002 2.002 0 0 1-2 2h-4Zm0-7h4v-5h-4Z"/>`),
		g.Group(children),
	)
}