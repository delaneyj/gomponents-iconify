package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderListMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M608 128q45 0 77 9t58 24t46 31t40 31t44 23t55 10h992q27 0 50 10t40 27t28 41t10 50v896h-128V384H928q-31 0-54 9t-44 24t-41 31t-45 31t-58 23t-78 10H128v1152h896v128H0V256q0-27 10-50t27-40t41-28t50-10h480zm0 256q24 0 42-4t33-13t29-20t32-27q-17-15-31-26t-30-20t-33-13t-42-5H128v128h480zm544 1536h640v128h-640v-128zm768 0h128v128h-128v-128zm0-256h128v128h-128v-128zm-768 0h640v128h-640v-128zm768-256h128v128h-128v-128zm-768 0h640v128h-640v-128z"/>`),
		g.Group(children),
	)
}