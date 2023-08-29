package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsMergeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m160 200l-32 32l-32-32Z" opacity=".2"/><path d="M184 32a8 8 0 0 0-8 8v60.69l-48 48l-48-48V40a8 8 0 0 0-16 0v64a8 8 0 0 0 2.34 5.66L120 163.31V192H96a8 8 0 0 0-5.66 13.66l32 32a8 8 0 0 0 11.32 0l32-32A8 8 0 0 0 160 192h-24v-28.69l53.66-53.65A8 8 0 0 0 192 104V40a8 8 0 0 0-8-8Zm-56 188.69L115.31 208h25.38Z"/></g>`),
		g.Group(children),
	)
}