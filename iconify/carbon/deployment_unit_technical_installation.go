package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeploymentUnitTechnicalInstallation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 11h3v10h-3v2h8v-2h-3V11h3V9h-8v2zM7 11h3v12h2V11h3V9H7v2z"/>`),
		g.Group(children),
	)
}