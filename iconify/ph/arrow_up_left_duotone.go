package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M168 64L64 168V64Z" opacity=".2"/><path d="M197.66 186.34L127.31 116l46.35-46.34A8 8 0 0 0 168 56H64a8 8 0 0 0-8 8v104a8 8 0 0 0 13.66 5.66L116 127.31l70.34 70.35a8 8 0 0 0 11.32-11.32ZM72 72h76.69l-38.34 38.34L72 148.69Z"/></g>`),
		g.Group(children),
	)
}