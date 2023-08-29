package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 256v1024h-896v-256H384v1024H256V0h896v256h768zm-896 640V128H384v768h640zm768-512h-640v768h640V384z"/>`),
		g.Group(children),
	)
}