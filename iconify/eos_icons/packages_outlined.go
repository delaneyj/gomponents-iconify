package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackagesOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6h16v2H4zm2-4h12v2H6zm-4 8v2h2v8a2.003 2.003 0 0 0 2 2h12a2.006 2.006 0 0 0 2.004-2L20 12h2v-2Zm16 10H6v-8h12Z"/><path fill="currentColor" d="M8 17h4v2H8z"/>`),
		g.Group(children),
	)
}