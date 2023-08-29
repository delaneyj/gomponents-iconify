package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EthereumSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a.5.5 0 0 1 .384.18l5 6a.5.5 0 0 1 .04.585l-5 8a.5.5 0 0 1-.848 0l-5-8a.5.5 0 0 1 .04-.585l5-6A.5.5 0 0 1 7.5 0ZM3.241 6.742L7.5 5.04l4.259 1.703L7.5 13.557L3.241 6.742Zm7.61-1.44L7.5 3.962l-3.35 1.34L7.5 1.28l3.35 4.02Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}