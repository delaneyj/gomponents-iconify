package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningStandardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34.6 29.21L20.71 3.65a3.22 3.22 0 0 0-5.66 0L1.17 29.21A3.22 3.22 0 0 0 4 34h27.77a3.22 3.22 0 0 0 2.83-4.75ZM16.6 10a1.4 1.4 0 0 1 2.8 0v12a1.4 1.4 0 0 1-2.8 0ZM18 29.85a1.8 1.8 0 1 1 1.8-1.8a1.8 1.8 0 0 1-1.8 1.8Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}