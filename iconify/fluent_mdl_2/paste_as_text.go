package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasteAsText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1792h768v128H128V256h512q0-53 20-99t55-82t81-55T896 0q53 0 99 20t82 55t55 81t20 100h512v640h-128V384h-128v256H384V384H256v1408zM512 384v128h768V384h-256V256q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50v128H512zm1497 1664h-135l-85-256h-378l-85 256h-135l342-1024h134l342 1024zm-263-384l-146-437l-146 437h292z"/>`),
		g.Group(children),
	)
}