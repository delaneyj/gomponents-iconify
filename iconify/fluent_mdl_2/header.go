package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Header(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 0h1792v2048H128V0zm1664 1920V128H256v1792h1536zM1664 256v512H384V256h1280zm-128 384V384H512v256h1024z"/>`),
		g.Group(children),
	)
}