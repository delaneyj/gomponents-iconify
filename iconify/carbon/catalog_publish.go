package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CatalogPublish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m20 20l1.4 1.4l-2.6 2.6H28v2h-9.2l2.6 2.6L20 30l-5-5zm-6-5h8v2h-8zm0-7h8v2h-8z"/><path fill="currentColor" d="M13 28H8v-4h2v-2H8v-5h2v-2H8v-5h2V8H8V4h18v16h2V4c0-1.1-.9-2-2-2H8c-1.1 0-2 .9-2 2v4H4v2h2v5H4v2h2v5H4v2h2v4c0 1.1.9 2 2 2h5v-2z"/>`),
		g.Group(children),
	)
}