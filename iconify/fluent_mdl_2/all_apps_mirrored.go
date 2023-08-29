package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllAppsMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1408v-384h-384v384h384zm-128-256v128h-128v-128h128zm128-256V512h-384v384h384zm-128-256v128h-128V640h128zm128-256V0h-384v384h384zm-128-256v128h-128V128h128zm-512 640V640H256v128h1152zm-896 384v128h896v-128H512zm896-1024H0v128h1408V128zm640 1792v-384h-384v384h384zm-128-256v128h-128v-128h128zm-512 128v-128H256v128h1152z"/>`),
		g.Group(children),
	)
}