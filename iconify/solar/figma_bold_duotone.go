package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.667 2H8.333a3.333 3.333 0 1 0 0 6.667h3.334V2Z" opacity=".6"/><path d="M11.667 8.666H8.333a3.333 3.333 0 0 0 0 6.667h3.334V8.666Z" opacity=".4"/><path d="M18.333 12a3.333 3.333 0 1 1-6.667 0a3.333 3.333 0 0 1 6.667 0Z"/><path d="M8.333 15.334h3.334v3.333a3.333 3.333 0 1 1-3.334-3.334Z" opacity=".2"/><path d="M11.666 2h3.333a3.333 3.333 0 1 1 0 6.667h-3.333V2Z" opacity=".8"/></g>`),
		g.Group(children),
	)
}