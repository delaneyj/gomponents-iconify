package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDataLinkThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 2v7a3 3 0 0 0 3 3h7v15a3 3 0 0 1-3 3h-7.531a5.985 5.985 0 0 0 2.031-4.5a5.977 5.977 0 0 0-1.5-3.969V19a1 1 0 1 0-2 0v1.044a5.978 5.978 0 0 0-2.5-.544H12V13a1 1 0 1 0-2 0v6.5H6.5a6.01 6.01 0 0 0-1.5.189V5a3 3 0 0 1 3-3h9Zm2 .117V9a1 1 0 0 0 1 1h6.883a3 3 0 0 0-.762-1.293L20.293 2.88A3 3 0 0 0 19 2.117ZM22 16a1 1 0 1 0-2 0v9a1 1 0 1 0 2 0v-9ZM8 22a1 1 0 0 0-1-1h-.5a4.5 4.5 0 1 0 0 9H7a1 1 0 1 0 0-2h-.5a2.5 2.5 0 0 1 0-5H7a1 1 0 0 0 1-1Zm4-1a1 1 0 1 0 0 2h.5a2.5 2.5 0 0 1 0 5H12a1 1 0 1 0 0 2h.5a4.5 4.5 0 1 0 0-9H12Zm-5 3.5a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2H7Z"/>`),
		g.Group(children),
	)
}