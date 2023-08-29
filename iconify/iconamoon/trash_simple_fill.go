package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2a1 1 0 0 0-.894.553L6.382 6H4a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8h1a1 1 0 1 0 0-2h-2.382l-1.724-3.447A1 1 0 0 0 15 2H9Zm6.382 4l-1-2H9.618l-1 2h6.764Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}