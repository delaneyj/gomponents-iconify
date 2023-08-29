package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Org(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 1280h256v640h-640v-640h256v-256H384v256h256v640H0v-640h256V896h640V640H640V0h640v640h-256v256h640v384zM768 128v384h384V128H768zM512 1792v-384H128v384h384zm1280 0v-384h-384v384h384z"/>`),
		g.Group(children),
	)
}