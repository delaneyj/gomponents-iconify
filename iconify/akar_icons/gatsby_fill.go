package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GatsbyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#akarIconsGatsbyFill0)"><path fill="currentColor" d="M12 0C5.323 0 0 5.317 0 12s5.317 12 12 12s12-5.323 12-12S18.683 0 12 0ZM2.608 12.101l9.29 9.286c-5.114.005-9.29-4.171-9.29-9.286Zm11.477 9.083L2.821 9.909C3.76 5.733 7.515 2.603 12 2.603a9.493 9.493 0 0 1 7.616 3.861l-1.355 1.147A7.666 7.666 0 0 0 11.9 4.267A7.57 7.57 0 0 0 4.693 9.38l9.814 9.819c2.4-.837 4.277-2.923 4.8-5.43h-4.07V12h6.155c0 4.485-3.13 8.245-7.307 9.184Z"/></g><defs><clipPath id="akarIconsGatsbyFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}