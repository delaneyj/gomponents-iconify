package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneCommentBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 18l2-2h14V4H4v14zm9-12h5v8l-2.5-1.5L13 14V6z" opacity=".3"/><path fill="currentColor" d="M18 14V6h-5v8l2.5-1.5z"/><path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v18l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm0 14H6l-2 2V4h16v12z"/>`),
		g.Group(children),
	)
}