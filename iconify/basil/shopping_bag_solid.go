package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 7.13v.37H5.749a.901.901 0 0 0-.892.77L4.64 9.763a30.316 30.316 0 0 0 0 8.79a2.885 2.885 0 0 0 2.557 2.451l.629.066c2.776.288 5.574.288 8.35 0l.63-.066a2.885 2.885 0 0 0 2.556-2.451a30.318 30.318 0 0 0 0-8.79l-.218-1.493a.901.901 0 0 0-.892-.77H16.75v-.37a4.75 4.75 0 1 0-9.5 0Zm5.56-3.147A3.25 3.25 0 0 0 8.75 7.13v.37h6.5v-.37a3.25 3.25 0 0 0-2.44-3.147ZM8.75 9a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0V9Zm8 0a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}