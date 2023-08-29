package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackToWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1152h640v640H768v-421L93 2045l-90-90l674-675H256v-128zm1115-384h421v128h-640V256h128v421L1955 3l90 90l-674 675z"/>`),
		g.Group(children),
	)
}