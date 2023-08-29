package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinusThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m202.83 58.83l-144 144a4 4 0 0 1-5.66-5.66l144-144a4 4 0 1 1 5.66 5.66ZM68 112a4 4 0 0 0 8 0V76h36a4 4 0 0 0 0-8H76V32a4 4 0 0 0-8 0v36H32a4 4 0 0 0 0 8h36Zm156 68h-80a4 4 0 0 0 0 8h80a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}