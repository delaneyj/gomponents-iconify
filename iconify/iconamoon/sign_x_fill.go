package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignXFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 9a1 1 0 1 0 0-2h-.68a3.64 3.64 0 0 0-2.912 1.456l-1.237 1.65l-.547-1.094A3.64 3.64 0 0 0 8.368 7H8a1 1 0 0 0 0 2h.368c.622 0 1.19.351 1.467.907l.994 1.987l-1.837 2.45A1.64 1.64 0 0 1 7.68 15H7a1 1 0 1 0 0 2h.68a3.64 3.64 0 0 0 2.912-1.456l1.237-1.65l.547 1.094A3.64 3.64 0 0 0 15.632 17H16a1 1 0 1 0 0-2h-.368a1.64 1.64 0 0 1-1.467-.907l-.994-1.987l1.837-2.45A1.64 1.64 0 0 1 16.32 9H17Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}