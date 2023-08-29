package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 128a72 72 0 1 1-72-72a72 72 0 0 1 72 72Z" opacity=".2"/><path d="m211 103.43l-70.13 28l49.47 63.61a8 8 0 1 1-12.63 9.82L128 141l-49.68 63.91a8 8 0 0 1-12.63-9.82l49.47-63.61L45 103.43a8 8 0 0 1 6-14.86l69 27.61V40a8 8 0 0 1 16 0v76.18l69-27.61a8 8 0 1 1 6 14.86Z"/></g>`),
		g.Group(children),
	)
}