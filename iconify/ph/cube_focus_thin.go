package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeFocusThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188 157.43V98.57a9 9 0 0 0-4.42-7.78l-51.34-29.65a8.5 8.5 0 0 0-8.48 0L72.43 90.79A9 9 0 0 0 68 98.57v58.86a9 9 0 0 0 4.42 7.78l51.34 29.65a8.52 8.52 0 0 0 8.48 0l51.33-29.65a9 9 0 0 0 4.43-7.78Zm-60.24-89.36a.45.45 0 0 1 .48 0L176.59 96L128 124.05L79.41 96Zm-51.33 90.21a1 1 0 0 1-.43-.85v-54.17L124 131v54.78Zm103.14 0L132 185.76V131l48-27.72v54.17a1 1 0 0 1-.43.83ZM228 48v40a4 4 0 0 1-8 0V52h-36a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4ZM76 208a4 4 0 0 1-4 4H32a4 4 0 0 1-4-4v-40a4 4 0 0 1 8 0v36h36a4 4 0 0 1 4 4Zm152-40v40a4 4 0 0 1-4 4h-40a4 4 0 0 1 0-8h36v-36a4 4 0 0 1 8 0ZM28 88V48a4 4 0 0 1 4-4h40a4 4 0 0 1 0 8H36v36a4 4 0 0 1-8 0Z"/>`),
		g.Group(children),
	)
}