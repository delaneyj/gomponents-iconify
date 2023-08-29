package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MopedFrontThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 44h-44.23a36 36 0 0 0-71.54 0H48a4 4 0 0 0 0 8h44.23a36.06 36.06 0 0 0 17.34 26.91A60.11 60.11 0 0 0 68 136v64a12 12 0 0 0 12 12h20v4a28 28 0 0 0 56 0v-4h20a12 12 0 0 0 12-12v-64a60.11 60.11 0 0 0-41.57-57.09A36.06 36.06 0 0 0 163.77 52H208a4 4 0 0 0 0-8Zm-60 172a20 20 0 0 1-40 0v-48a20 20 0 0 1 40 0Zm32-80v64a4 4 0 0 1-4 4h-20v-36a28 28 0 0 0-56 0v36H80a4 4 0 0 1-4-4v-64a52 52 0 0 1 104 0Zm-52-60a28 28 0 1 1 28-28a28 28 0 0 1-28 28Z"/>`),
		g.Group(children),
	)
}