package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReaderFollowConversation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 14v-3h-2v3h-3v2h3v3h2v-3h3v-2z"/><path fill="currentColor" d="M13 16h-2l-5 5v-5H4c-1.1 0-2-.9-2-2V5c0-1.1.9-2 2-2h14c1.1 0 2 .9 2 2v4h-4v3h-3v4z"/>`),
		g.Group(children),
	)
}