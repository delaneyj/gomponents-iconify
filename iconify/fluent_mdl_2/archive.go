package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v640h-128v1152H128V768H0V128h2048zm-256 1664V768H256v1024h1536zm128-1152V256H128v384h1792zm-512 512H640v-128h768v128z"/>`),
		g.Group(children),
	)
}