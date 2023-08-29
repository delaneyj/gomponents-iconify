package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AzurePipelines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2.401 24H0v8h8v-2.401H2.401zM30.667 0H18.761a2.666 2.666 0 0 0-2.229 1.188l-4.531 6.797H1.334C.6 7.985.001 8.584.001 9.318v7.333l4.667 4.818l1.599-1.724l2 2l-1.599 1.63l1.974 1.974l1.625-1.63l2 2.016l-1.599 1.599l4.667 4.667h7.333c.734 0 1.333-.599 1.333-1.333V20.001l6.813-4.531a2.673 2.673 0 0 0 1.188-2.24V1.334c0-.734-.599-1.333-1.333-1.333zM24 12c-5.333 0-5.333-8 0-8s5.333 8 0 8z"/>`),
		g.Group(children),
	)
}