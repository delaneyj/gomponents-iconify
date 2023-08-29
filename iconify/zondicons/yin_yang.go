package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 20a10 10 0 1 1 0-20a10 10 0 0 1 0 20zm0-18a8 8 0 1 0 0 16a4 4 0 1 1 0-8a4 4 0 1 0 0-8zm0 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm0-8a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}