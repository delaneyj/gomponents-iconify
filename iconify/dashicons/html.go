package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Html(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 16v-2H2v2H1v-5h1v2h2v-2h1v5H4zm3 0v-4H5.6v-1h3.7v1H8v4H7zm3 0v-5h1l1.4 3.4h.1L14 11h1v5h-1v-3.1h-.1l-1.1 2.5h-.6l-1.1-2.5H11V16h-1zm9 0h-3v-5h1v4h2v1zM9.4 4.2L7.1 6.5l2.3 2.3l-.6 1.2l-3.5-3.5L8.8 3l.6 1.2zm1.2 4.6l2.3-2.3l-2.3-2.3l.6-1.2l3.5 3.5l-3.5 3.5l-.6-1.2z"/>`),
		g.Group(children),
	)
}