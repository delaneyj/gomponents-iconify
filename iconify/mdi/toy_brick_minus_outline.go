package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToyBrickMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.09 20H3V6h2V5c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1h2V5c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1h2v7.35c-.63-.22-1.3-.35-2-.35V8H5v10h8.09c-.05.33-.09.66-.09 1s.04.67.09 1M23 18h-8v2h8v-2Z"/>`),
		g.Group(children),
	)
}