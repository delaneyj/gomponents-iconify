package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M19.06 55.61c.485.178 1.03.297 1.576.297c.849 0 1.697-.297 2.424-.772l30-15.98l.303-.297c.788-.772 1.212-1.723 1.212-2.792s-.424-2.08-1.212-2.792l-.303-.297l-30-16.1c-1.091-.832-2.667-1.01-4-.475c-1.515.594-2.485 2.079-2.485 3.683v31.84c0 1.604.97 3.089 2.485 3.683z" transform="matrix(0 -.9544 .9545 0 1.727 66.7)"/><path stroke-linecap="round" d="M17 57.38h38"/></g>`),
		g.Group(children),
	)
}