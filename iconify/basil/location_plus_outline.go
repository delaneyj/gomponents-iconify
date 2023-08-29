package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6.25a.75.75 0 0 1 .75.75v2h2a.75.75 0 0 1 0 1.5h-2v2a.75.75 0 1 1-1.5 0v-2h-2a.75.75 0 0 1 0-1.5h2V7a.75.75 0 0 1 .75-.75Z"/><path fill="currentColor" fill-rule="evenodd" d="M11.784 1.25a8.288 8.288 0 0 0-8.26 7.607a8.943 8.943 0 0 0 1.99 6.396l4.793 5.861a2.187 2.187 0 0 0 3.386 0l4.793-5.861a8.944 8.944 0 0 0 1.99-6.396a8.288 8.288 0 0 0-8.26-7.607h-.432ZM5.02 8.98a6.788 6.788 0 0 1 6.765-6.23h.432a6.788 6.788 0 0 1 6.765 6.23a7.443 7.443 0 0 1-1.656 5.323l-4.793 5.862a.687.687 0 0 1-1.064 0l-4.793-5.862A7.443 7.443 0 0 1 5.02 8.98Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}