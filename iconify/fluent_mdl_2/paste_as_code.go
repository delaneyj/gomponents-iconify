package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasteAsCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1792h512v128H128V256h512q0-53 20-99t55-82t81-55T896 0q53 0 99 20t82 55t55 81t20 100h512v640h-128V384h-128v256H384V384H256v1408zM512 384v128h768V384h-256V256q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50v128H512zm525 941l-210 211l210 211l-90 90l-301-301l301-301l90 90zm1005 211l-301 301l-90-90l210-211l-210-211l90-90l301 301zm-549-512h128l-341 1024h-128l341-1024z"/>`),
		g.Group(children),
	)
}