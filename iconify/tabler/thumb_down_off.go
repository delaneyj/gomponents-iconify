package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDownOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 13V7M4 4a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h3a4 4 0 0 1 4 4v1a2 2 0 1 0 4 0v-3m2-2h1a2 2 0 0 0 2-2l-1-5c-.295-1.26-1.11-2.076-2-2h-7c-.57 0-1.102.159-1.556.434M3 3l18 18"/>`),
		g.Group(children),
	)
}