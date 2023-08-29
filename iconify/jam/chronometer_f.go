package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChronometerF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11V8a1 1 0 1 0-2 0v4a1 1 0 0 0 1 1h3a1 1 0 0 0 0-2H8zm-1 7A6 6 0 1 1 7 6a6 6 0 0 1 0 12zM9.928 4.553A7.98 7.98 0 0 0 7 4a7.98 7.98 0 0 0-2.928.553A2.5 2.5 0 0 1 5.5 0h3a2.5 2.5 0 0 1 1.428 4.553zM5.5 2a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3zm5.554 3.101a1.5 1.5 0 1 1 2.077 1.76a8.039 8.039 0 0 0-2.077-1.76zM.869 6.861a1.5 1.5 0 1 1 2.077-1.76A8.039 8.039 0 0 0 .87 6.861z"/>`),
		g.Group(children),
	)
}