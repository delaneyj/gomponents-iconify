package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBook0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBook1" fill="currentColor" fill-rule="nonzero"><path id="feBook2" d="m13 16.006l7-.047V5.992l-5.17.007l-1.814 1.814L13 16.006Zm-2-8.193L9.179 6.038L4 6.003v9.956l7 .047V7.813Zm-1-3.77L12 6l2-2l5.997-.008A2 2 0 0 1 22 5.989v9.97a2 2 0 0 1-1.986 2L14 18l-1.996 2L10 18l-6.014-.041a2 2 0 0 1-1.986-2V6.003a2 2 0 0 1 2-2l6 .04Z"/></g></g>`),
		g.Group(children),
	)
}