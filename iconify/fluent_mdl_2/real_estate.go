package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RealEstate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1280 1920h384v128h-512v-896H768v896H256V987l-83 82l-90-90l877-877l877 877l-90 90l-83-82v165h-128V859L960 282L384 859v1061h256v-896h640v896zm768-640v768h-128v-256h-512v-512h640zm-128 128h-384v256h384v-256z"/>`),
		g.Group(children),
	)
}