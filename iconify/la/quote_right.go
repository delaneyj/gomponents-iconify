package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 8v10h8c0 2.219-1.781 4-4 4v2c3.3 0 6-2.7 6-6V8zm14 0v10h8c0 2.219-1.781 4-4 4v2c3.3 0 6-2.7 6-6V8zM6 10h6v6H6zm14 0h6v6h-6z"/>`),
		g.Group(children),
	)
}