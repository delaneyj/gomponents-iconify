package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheveronOutlineRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0a10 10 0 1 1 0 20a10 10 0 0 1 0-20zM2 10a8 8 0 1 0 16 0a8 8 0 0 0-16 0zm10.54.7L9 14.25l-1.41-1.41L10.4 10L7.6 7.17L9 5.76L13.24 10l-.7.7z"/>`),
		g.Group(children),
	)
}