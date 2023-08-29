package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1280 256v640H0V256h1280zm-128 512V384H128v384h1024zm256-512h640v640h-640V256zm512 512V384h-384v384h384zM768 1664v-640h1280v640H768zm128-512v384h1024v-384H896zM0 1664v-640h640v640H0zm128-512v384h384v-384H128z"/>`),
		g.Group(children),
	)
}