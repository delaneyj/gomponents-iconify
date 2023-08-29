package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortLinesAscending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m3 483l317-318l317 318l-90 90l-163-163v1510H256V410L93 573L3 483zm1277-99v128H768V384h512zm384 384v128H768V768h896zm-896 384h1280v128H768v-128z"/>`),
		g.Group(children),
	)
}