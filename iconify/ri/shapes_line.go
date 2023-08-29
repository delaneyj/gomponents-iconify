package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1L6 11h12L12 1Zm0 3.887L14.468 9H9.532L12 4.887ZM6.75 20a2.75 2.75 0 1 1 0-5.5a2.75 2.75 0 0 1 0 5.5Zm0 2a4.75 4.75 0 1 0 0-9.5a4.75 4.75 0 0 0 0 9.5ZM15 15.5v4h4v-4h-4Zm-2 6v-8h8v8h-8Z"/>`),
		g.Group(children),
	)
}