package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFooterDismissTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 5.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-2.646-1.146a.5.5 0 0 0-.708-.708L14.5 4.793l-1.146-1.147a.5.5 0 0 0-.708.708L13.793 5.5l-1.147 1.146a.5.5 0 0 0 .708.708L14.5 6.207l1.146 1.147a.5.5 0 0 0 .708-.708L15.207 5.5l1.147-1.146ZM15 16v-5.022a5.489 5.489 0 0 0 1-.185V16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4.257A5.504 5.504 0 0 0 9.6 3H6a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1Zm-8-2a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H7Z"/>`),
		g.Group(children),
	)
}