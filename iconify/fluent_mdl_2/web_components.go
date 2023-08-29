package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebComponents(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1024h768v896H128V256h896v768zM256 384v640h640V384H256zm640 1408v-640H256v640h640zm768 0v-640h-640v640h640zM1280 0h768v768h-768V0zm640 640V128h-512v512h512z"/>`),
		g.Group(children),
	)
}