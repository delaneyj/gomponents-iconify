package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportAllMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1030 896v128l772 1l-289 290l90 90l445-445l-445-445l-90 90l292 292l-775-1zM902 512H6v896h896V512zM134 896V640h256v256H134zm384-256h256v256H518V640zm256 384v256H518v-256h256zm-384 256H134v-256h256v256z"/>`),
		g.Group(children),
	)
}