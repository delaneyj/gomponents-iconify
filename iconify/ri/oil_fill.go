package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5h11a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V11l4-6Zm5-4h5a1 1 0 0 1 1 1v2h-7V2a1 1 0 0 1 1-1ZM6 12v7h2v-7H6Z"/>`),
		g.Group(children),
	)
}