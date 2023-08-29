package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretLeftSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.707 1.026L1.735 7.5l6.972 6.474l-.68.733L.264 7.5L8.026.293l.68.733Zm5.5 0L7.235 7.5l6.972 6.474l-.68.733L5.764 7.5L13.526.293l.68.733Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}