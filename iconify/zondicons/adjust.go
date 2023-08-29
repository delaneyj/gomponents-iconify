package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2v16a8 8 0 1 0 0-16zm0 18a10 10 0 1 1 0-20a10 10 0 0 1 0 20z"/>`),
		g.Group(children),
	)
}