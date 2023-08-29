package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.531 25.599L20.803 16l5.728-9.599L29.334 16zm-12.667-.932l-7.197-7.068h11.469l5.728 9.599zm0-17.47l10-2.531l-5.728 9.599H6.531c0 .136 7.333-7.068 7.333-7.068zM28.531 0L15.47 3.333l-2 3.333H9.599L0 15.999l9.599 9.333h3.871l1.864 3.333l13.068 3.333l3.463-12.667l-1.864-3.333l2-3.333z"/>`),
		g.Group(children),
	)
}