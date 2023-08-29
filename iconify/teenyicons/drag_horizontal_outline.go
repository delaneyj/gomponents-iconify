package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHorizontalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-10 4a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/>`),
		g.Group(children),
	)
}