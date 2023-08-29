package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReaderFollowingConversation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.8 14.5l3.2-3.2V5c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v9c0 1.1.9 2 2 2h2v5l8.7-8.7l2.1 2.2z"/><path fill="currentColor" d="m22.6 11.1l-6.1 6.1l-2.1-2.2l-1.4 1.4l3.5 3.6l7.5-7.6z"/>`),
		g.Group(children),
	)
}