package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Status(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4c4.411 0 8 3.589 8 8s-3.589 8-8 8s-8-3.589-8-8s3.589-8 8-8m0-2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zM7.55 13c-.019.166-.05.329-.05.5a4.5 4.5 0 0 0 9 0c0-.171-.032-.334-.05-.5h-8.9zM10 10V8a1 1 0 0 0-2 0v2a1 1 0 0 0 2 0zm6 0V8a1 1 0 0 0-2 0v2a1 1 0 0 0 2 0z"/>`),
		g.Group(children),
	)
}