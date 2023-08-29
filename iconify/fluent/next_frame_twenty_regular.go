package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextFrameTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 3a.5.5 0 0 1 .5.5v13a.5.5 0 0 1-1 0v-13a.5.5 0 0 1 .5-.5Zm5.447.214A1.25 1.25 0 0 0 8 4.252v11.5a1.25 1.25 0 0 0 1.954 1.033l8.5-5.793a1.25 1.25 0 0 0-.008-2.07l-8.5-5.708ZM9 4.252c0-.2.223-.319.39-.207l8.499 5.707a.25.25 0 0 1 .001.414l-8.5 5.793a.25.25 0 0 1-.39-.207v-11.5Z"/>`),
		g.Group(children),
	)
}