package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deezer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.68 5.47H22V8h-4.32zm0 3.51H22v2.53h-4.32zm0 3.51H22v2.53h-4.32zM2 16h4.32v2.53H2zm5.22 0h4.32v2.53H7.22zm5.23 0h4.32v2.53h-4.32zm5.23 0H22v2.53h-4.32zm-5.23-3.51h4.32v2.53h-4.32zm-5.23 0h4.32v2.53H7.22zm0-3.51h4.32v2.53H7.22z"/>`),
		g.Group(children),
	)
}