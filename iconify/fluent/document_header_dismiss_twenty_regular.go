package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentHeaderDismissTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h4.257a5.503 5.503 0 0 1-.657-1H6a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v5.022a5.5 5.5 0 0 1 1 .185V4a2 2 0 0 0-2-2H6Zm1 2a1 1 0 0 0 0 2h6a1 1 0 1 0 0-2H7Zm12 10.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-2.646-1.146a.5.5 0 0 0-.708-.708L14.5 13.793l-1.146-1.147a.5.5 0 0 0-.708.708l1.147 1.146l-1.147 1.146a.5.5 0 0 0 .708.708l1.146-1.147l1.146 1.147a.5.5 0 0 0 .708-.708L15.207 14.5l1.147-1.146Z"/>`),
		g.Group(children),
	)
}