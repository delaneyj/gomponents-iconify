package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3H7a2 2 0 0 0-2 2v16l7-4l7 4V5a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}