package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocLibrary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1627 640l421 421v987H896V640h731zm37 384h165l-165-165v165zm256 896v-768h-384V768h-512v1152h896zM608 128q45 0 77 9t58 24t46 31t40 31t44 23t55 10h992q27 0 50 10t40 27t28 41t10 50v451l-128-128V384H928q-31 0-54 9t-44 24t-41 31t-45 31t-58 23t-78 10H128v1152h640v128H0V256q0-27 10-50t27-40t41-28t50-10h480zm0 256q24 0 42-4t33-13t29-20t32-27q-17-15-31-26t-30-20t-33-13t-42-5H128v128h480z"/>`),
		g.Group(children),
	)
}