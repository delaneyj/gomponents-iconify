package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddFavorite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1370 1536l-103-329l405-311h-502l-146-467l-138 467H384l397 311l-154 493l397-306l256 197v162l-256-197l-640 492l248-794L0 768h784L1024 0l240 768h784l-632 486l88 282h-134zm422-128v256h256v128h-256v256h-128v-256h-256v-128h256v-256h128z"/>`),
		g.Group(children),
	)
}