package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutlineRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 10a10 10 0 1 1-20 0a10 10 0 0 1 20 0zm-2 0a8 8 0 1 0-16 0a8 8 0 0 0 16 0zm-8 2H5V8h5V5l5 5l-5 5v-3z"/>`),
		g.Group(children),
	)
}