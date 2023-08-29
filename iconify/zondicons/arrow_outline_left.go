package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutlineLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 10a10 10 0 1 1 20 0a10 10 0 0 1-20 0zm2 0a8 8 0 1 0 16 0a8 8 0 0 0-16 0zm8-2h5v4h-5v3l-5-5l5-5v3z"/>`),
		g.Group(children),
	)
}