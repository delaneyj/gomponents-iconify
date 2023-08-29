package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 6a4 4 0 1 1 8 0v8h3l-4 4l-4-4h3V6a2 2 0 0 0-2-2a2 2 0 0 0-2 2v8a4 4 0 1 1-8 0V6H0l4-4l4 4H5v8a2 2 0 0 0 2 2a2 2 0 0 0 2-2V6z"/>`),
		g.Group(children),
	)
}