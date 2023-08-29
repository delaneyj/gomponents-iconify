package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarWashingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21H5v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-9l2.417-4.029a2 2 0 0 1 1.715-.97h11.736a2 2 0 0 1 1.715.97L22 13.001v9a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1ZM4.332 13h15.336l-1.8-3H6.132l-1.8 3ZM6.5 18a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm11 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM5.44 3.44L6.5 2.379l1.06 1.06a1.5 1.5 0 1 1-2.121 0Zm5.5 0L12 2.379l1.06 1.06a1.5 1.5 0 1 1-2.121 0Zm5.5 0l1.06-1.061l1.06 1.06a1.5 1.5 0 1 1-2.121 0Z"/>`),
		g.Group(children),
	)
}