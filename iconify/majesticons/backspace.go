package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M22 8a3 3 0 0 0-3-3H8.828a3 3 0 0 0-2.12.879l-4 4a3 3 0 0 0 0 4.242l4 4a3 3 0 0 0 2.12.879H19a3 3 0 0 0 3-3V8zm-6.707 1.293a1 1 0 1 1 1.414 1.414L15.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L14 13.414l-1.293 1.293a1 1 0 0 1-1.414-1.414L12.586 12l-1.293-1.293a1 1 0 0 1 1.414-1.414L14 10.586l1.293-1.293z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}