package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheveronOutlineDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 10a10 10 0 1 1-20 0a10 10 0 0 1 20 0zM10 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16zm-.7 10.54L5.75 9l1.41-1.41L10 10.4l2.83-2.82L14.24 9L10 13.24l-.7-.7z"/>`),
		g.Group(children),
	)
}