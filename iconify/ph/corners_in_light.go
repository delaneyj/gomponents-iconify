package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornersInLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M154 96V48a6 6 0 0 1 12 0v42h42a6 6 0 0 1 0 12h-48a6 6 0 0 1-6-6Zm-58 58H48a6 6 0 0 0 0 12h42v42a6 6 0 0 0 12 0v-48a6 6 0 0 0-6-6Zm112 0h-48a6 6 0 0 0-6 6v48a6 6 0 0 0 12 0v-42h42a6 6 0 0 0 0-12ZM96 42a6 6 0 0 0-6 6v42H48a6 6 0 0 0 0 12h48a6 6 0 0 0 6-6V48a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}