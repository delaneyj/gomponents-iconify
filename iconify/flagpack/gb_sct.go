package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GbSct(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#flagpackGbSct0)"><path fill="#265AAD" d="M0 0h32v24H0z"/><path fill="#fff" fill-rule="evenodd" d="M16 14.5L1.2 25.6l-2.4-3.2L12.667 12L-1.2 1.6l2.4-3.2L16 9.5L30.8-1.6l2.4 3.2L19.333 12L33.2 22.4l-2.4 3.2L16 14.5Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGbSct0"><path fill="#fff" d="M0 0h32v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}