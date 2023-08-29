package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TilesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 2048V768h640v1280H256zM384 896v1024h384V896H384zM256 640V0h640v640H256zm128-512v384h384V128H384zM1024 0h640v1280h-640V0zm512 1152V128h-384v1024h384zm-512 896v-640h640v640h-640zm128-512v384h384v-384h-384z"/>`),
		g.Group(children),
	)
}