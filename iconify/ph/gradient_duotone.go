package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradientDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 64v120H32V64Z" opacity=".2"/><path d="M24 104a8 8 0 0 1 8-8h80a8 8 0 0 1 0 16H32a8 8 0 0 1-8-8Zm200-8h-80a8 8 0 0 0 0 16h80a8 8 0 0 0 0-16ZM72 136H32a8 8 0 0 0 0 16h40a8 8 0 0 0 0-16Zm152 0h-40a8 8 0 0 0 0 16h40a8 8 0 0 0 0-16Zm-128 8a8 8 0 0 0 8 8h48a8 8 0 0 0 0-16h-48a8 8 0 0 0-8 8Zm-40 32H32a8 8 0 0 0 0 16h24a8 8 0 0 0 0-16Zm56 0H88a8 8 0 0 0 0 16h24a8 8 0 0 0 0-16Zm56 0h-24a8 8 0 0 0 0 16h24a8 8 0 0 0 0-16Zm56 0h-24a8 8 0 0 0 0 16h24a8 8 0 0 0 0-16ZM32 72h192a8 8 0 0 0 0-16H32a8 8 0 0 0 0 16Z"/></g>`),
		g.Group(children),
	)
}