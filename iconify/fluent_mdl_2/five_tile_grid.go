package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveTileGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h2048v1408H0V256zm1024 128v512h384V384h-384zM896 1536V384H128v1152h768zm512 0v-512h-384v512h384zm512 0v-512h-384v512h384zm-384-640h384V384h-384v512z"/>`),
		g.Group(children),
	)
}