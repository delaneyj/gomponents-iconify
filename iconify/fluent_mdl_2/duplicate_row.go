package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DuplicateRow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 768h384v768H384v-384H0V384h1664v384zM384 1024V768h1152V512H128v512h256zm1536 384V896H512v512h1408z"/>`),
		g.Group(children),
	)
}