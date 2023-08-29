package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickReply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M13 18l-7-7l7-7v4c8 0 11 11 11 11s-3-5-11-5v4zm-6 0l-7-7l7-7v3l-4 4l4 4v3z" fill="currentColor"/>`),
		g.Group(children),
	)
}