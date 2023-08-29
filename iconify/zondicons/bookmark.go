package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 2c0-1.1.9-2 2-2h12a2 2 0 0 1 2 2v18l-8-4l-8 4V2z"/>`),
		g.Group(children),
	)
}