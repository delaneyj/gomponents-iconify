package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CRMProcesses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h2048v1152H0V384zm128 1024h293l449-448l-449-448H128v896zm1792 0V512H603l447 448l-447 448h1317z"/>`),
		g.Group(children),
	)
}