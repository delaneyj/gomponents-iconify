package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.35 10.48H4.5l-.24-1.25h9.13a1.24 1.24 0 0 0 1.22-1l.84-4a1.25 1.25 0 0 0-1.22-1.51H3l-.22-1.24H.5v1.25h1.25l1.5 7.84a2 2 0 0 0-1.54 1.93a2.09 2.09 0 0 0 2.16 2a2.08 2.08 0 0 0 2.13-2a2 2 0 0 0-.16-.77h5.49a2 2 0 0 0-.16.77a2.09 2.09 0 0 0 2.16 2a2 2 0 1 0 0-4zM14.23 4l-.84 4H4l-.74-4zM3.87 13.27A.85.85 0 0 1 3 12.5a.85.85 0 0 1 .91-.77a.84.84 0 0 1 .9.77a.84.84 0 0 1-.94.77zm9.48 0a.85.85 0 0 1-.91-.77a.92.92 0 0 1 1.81 0a.85.85 0 0 1-.9.77z"/>`),
		g.Group(children),
	)
}