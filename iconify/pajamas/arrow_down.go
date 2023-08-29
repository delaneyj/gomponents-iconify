package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.72 10.159a.75.75 0 1 1 1.06 1.06l-3.25 3.25L8 15l-.53-.53l-3.25-3.25a.75.75 0 0 1 1.06-1.061l1.97 1.97V1.75a.75.75 0 1 1 1.5 0v10.379l1.97-1.97Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}