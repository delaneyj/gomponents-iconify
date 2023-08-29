package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceVision(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 14v7l1 2l1-2v-7h-2zm-2 0h-4a2 2 0 0 0-2 2v4l1 2l1-2v-4h2v7l1 2l1-2z"/><path fill="currentColor" d="M4 18A12 12 0 1 0 16 6h-4V1L6 7l6 6V8h4A10 10 0 1 1 6 18Z"/>`),
		g.Group(children),
	)
}