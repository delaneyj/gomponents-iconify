package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M168 192H64V88Z" opacity=".2"/><path d="M197.66 58.34a8 8 0 0 0-11.32 0L116 128.69L69.66 82.34A8 8 0 0 0 56 88v104a8 8 0 0 0 8 8h104a8 8 0 0 0 5.66-13.66L127.31 140l70.35-70.34a8 8 0 0 0 0-11.32ZM72 184v-76.69l38.34 38.34L148.69 184Z"/></g>`),
		g.Group(children),
	)
}