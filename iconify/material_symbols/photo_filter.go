package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h9v2H5v14h14v-9h2v9q0 .825-.588 1.413T19 21H5Zm7-5l-1.25-2.75L8 12l2.75-1.25L12 8l1.25 2.75L16 12l-2.75 1.25L12 16Zm5-6l-.95-2.05L14 7l2.05-.95L17 4l.95 2.05L20 7l-2.05.95L17 10Z"/>`),
		g.Group(children),
	)
}