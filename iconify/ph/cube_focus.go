package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 157.43V98.57a13 13 0 0 0-6.42-11.24l-51.34-29.65a12.48 12.48 0 0 0-12.48 0L70.43 87.33A13 13 0 0 0 64 98.57v58.86a13 13 0 0 0 6.42 11.24l51.34 29.65a12.49 12.49 0 0 0 12.48 0l51.33-29.65a13 13 0 0 0 6.43-11.24Zm-64-84.88L168.59 96L128 119.43L87.41 96Zm-48 37.64l40 23.1v45.54l-40-23.1Zm56 68.64v-45.54l40-23.1v45.54ZM232 48v40a8 8 0 0 1-16 0V56h-32a8 8 0 0 1 0-16h40a8 8 0 0 1 8 8ZM80 208a8 8 0 0 1-8 8H32a8 8 0 0 1-8-8v-40a8 8 0 0 1 16 0v32h32a8 8 0 0 1 8 8Zm152-40v40a8 8 0 0 1-8 8h-40a8 8 0 0 1 0-16h32v-32a8 8 0 0 1 16 0ZM24 88V48a8 8 0 0 1 8-8h40a8 8 0 0 1 0 16H40v32a8 8 0 0 1-16 0Z"/>`),
		g.Group(children),
	)
}