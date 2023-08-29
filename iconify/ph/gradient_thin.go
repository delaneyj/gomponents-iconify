package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradientThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M28 104a4 4 0 0 1 4-4h80a4 4 0 0 1 0 8H32a4 4 0 0 1-4-4Zm196-4h-80a4 4 0 0 0 0 8h80a4 4 0 0 0 0-8ZM72 140H32a4 4 0 0 0 0 8h40a4 4 0 0 0 0-8Zm152 0h-40a4 4 0 0 0 0 8h40a4 4 0 0 0 0-8Zm-124 4a4 4 0 0 0 4 4h48a4 4 0 0 0 0-8h-48a4 4 0 0 0-4 4Zm-44 36H32a4 4 0 0 0 0 8h24a4 4 0 0 0 0-8Zm56 0H88a4 4 0 0 0 0 8h24a4 4 0 0 0 0-8Zm56 0h-24a4 4 0 0 0 0 8h24a4 4 0 0 0 0-8Zm56 0h-24a4 4 0 0 0 0 8h24a4 4 0 0 0 0-8ZM32 68h192a4 4 0 0 0 0-8H32a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}