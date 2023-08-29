package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockersFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v176a8 8 0 0 0 16 0v-16h72v16a8 8 0 0 0 16 0v-16h72v16a8 8 0 0 0 16 0V48a16 16 0 0 0-16-16ZM96 112H56a8 8 0 0 1 0-16h40a8 8 0 0 1 0 16Zm0-32H56a8 8 0 0 1 0-16h40a8 8 0 0 1 0 16Zm40 104a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0Zm64-72h-40a8 8 0 0 1 0-16h40a8 8 0 0 1 0 16Zm0-32h-40a8 8 0 0 1 0-16h40a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}