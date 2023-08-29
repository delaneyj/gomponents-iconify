package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FTwelveDevTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1280h-86l-128-128h86V384H128v1024h512v128H0V256h2048zm-267 1280h-576q-218 219-437 437V523l1013 1013zm-309-128L896 832v832l256-256h320z"/>`),
		g.Group(children),
	)
}