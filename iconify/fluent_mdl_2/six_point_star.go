package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixPointStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1648 1024l307 512h-615l-316 514l-316-514H93l306-511L76 512h632L1024-2l316 514h615l-307 512zm-149 0l230-384h-460l-245-398l-245 398H308l242 383l-231 385h460l245 398l245-398h460l-230-384z"/>`),
		g.Group(children),
	)
}