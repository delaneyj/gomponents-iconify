package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M31.489 10.982L14.796.146c-.657-.427-1.463.136-1.463 1.02V7.85L1.462.15C.805-.277-.001.286-.001 1.17v21.663c0 .884.805 1.447 1.463 1.02l11.871-7.705v6.685c0 .884.805 1.447 1.463 1.02l16.693-10.835a1.27 1.27 0 0 0 .003-2.037l-.003-.002z"/>`),
		g.Group(children),
	)
}