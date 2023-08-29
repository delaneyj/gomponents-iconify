package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 14V6a1 1 0 0 0-2 0v8a1 1 0 1 0 2 0ZM7 4a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V6a2 2 0 0 0-2-2Zm7 10V6a1 1 0 1 0-2 0v8a1 1 0 1 0 2 0ZM13 4a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V6a2 2 0 0 0-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}