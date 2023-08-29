package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 136v64h-64Z" opacity=".2"/><path d="M235.06 128.61a8 8 0 0 0-8.72 1.73L200 156.69l-58.34-58.35a8 8 0 0 0-11.32 0L96 132.69L29.66 66.34a8 8 0 0 0-11.32 11.32l72 72a8 8 0 0 0 11.32 0L136 115.31L188.69 168l-26.35 26.34A8 8 0 0 0 168 208h64a8 8 0 0 0 8-8v-64a8 8 0 0 0-4.94-7.39ZM224 192h-36.69L224 155.31Z"/></g>`),
		g.Group(children),
	)
}