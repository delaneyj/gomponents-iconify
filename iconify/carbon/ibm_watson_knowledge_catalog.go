package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmWatsonKnowledgeCatalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23 27.2l-2.6-2.6L19 26l4 4l7-7l-1.4-1.4zM12 18h8v2h-8zm0-5h8v2h-8zm0-5h8v2h-8z"/><path fill="currentColor" d="M16 28H6v-4h2v-2H6v-5h2v-2H6v-5h2V8H6V4h18v16h2V4c0-1.1-.9-2-2-2H6c-1.1 0-2 .9-2 2v4H2v2h2v5H2v2h2v5H2v2h2v4c0 1.1.9 2 2 2h10v-2z"/>`),
		g.Group(children),
	)
}