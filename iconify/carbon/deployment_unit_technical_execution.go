package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeploymentUnitTechnicalExecution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 11V9h-8v14h8v-2h-6v-4h5v-2h-5v-4h6zM7 11h3v12h2V11h3V9H7v2z"/>`),
		g.Group(children),
	)
}