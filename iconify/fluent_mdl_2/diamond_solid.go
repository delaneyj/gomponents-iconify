package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m625 768l267-640h264l267 640H625zm971 0l-267-640h123l542 640h-398zm-174 128l-398 946l-398-946h796zm575 0l-847 1058l446-1058h401zm-1545 0l446 1058L51 896h401zm0-128H54l540-640h125L452 768z"/>`),
		g.Group(children),
	)
}