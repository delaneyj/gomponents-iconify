package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SidePanelMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1408H0V256h2048zm-128 384h-384v896h384V640zM128 1536h1280V640H128v896zM1920 512V384H128v128h1792z"/>`),
		g.Group(children),
	)
}