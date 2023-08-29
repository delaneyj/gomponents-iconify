package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cluster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 13v10h22V13Zm12 6H4v-2h9Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1ZM1 1v10h22V1Zm12 6H4V5h9Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}