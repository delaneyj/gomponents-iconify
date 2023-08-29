package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 4a2 2 0 0 0-2 2v6H0l4 4l4-4H5V6h7l2-2H5zm10 4h-3l4-4l4 4h-3v6a2 2 0 0 1-2 2H6l2-2h7V8z"/>`),
		g.Group(children),
	)
}