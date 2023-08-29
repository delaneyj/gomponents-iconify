package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextOutdentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M72 56v80L32 96Z" opacity=".2"/><path d="M224 128a8 8 0 0 1-8 8H112a8 8 0 0 1 0-16h104a8 8 0 0 1 8 8ZM112 72h104a8 8 0 0 0 0-16H112a8 8 0 0 0 0 16Zm104 112H40a8 8 0 0 0 0 16h176a8 8 0 0 0 0-16ZM66.34 141.66l-40-40a8 8 0 0 1 0-11.32l40-40A8 8 0 0 1 80 56v80a8 8 0 0 1-13.66 5.66ZM64 75.31L43.31 96L64 116.69Z"/></g>`),
		g.Group(children),
	)
}