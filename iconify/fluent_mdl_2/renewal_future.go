package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RenewalFuture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1792H0V128h384V0h128v128h1024V0h128v128h384zM128 256v256h1792V256h-256v128h-128V256H512v128H384V256H128zm1792 1536V640H128v1152h1792zm-870-576l-429 429l-90-90l339-339l-339-339l90-90l429 429zm83-429l429 429l-429 429l-90-90l339-339l-339-339l90-90z"/>`),
		g.Group(children),
	)
}