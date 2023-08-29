package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.657 14.828l-1.414-1.414L17.657 12A4 4 0 1 0 12 6.343l-1.414 1.414l-1.414-1.414l1.414-1.414a6 6 0 0 1 8.485 8.485l-1.414 1.414Zm-2.828 2.829l-1.415 1.414a6 6 0 0 1-8.485-8.485L6.343 9.17l1.415 1.415L6.343 12A4 4 0 0 0 12 17.657l1.415-1.415l1.414 1.415Zm0-9.9l1.414 1.414l-7.071 7.072l-1.414-1.415l7.07-7.07Z"/>`),
		g.Group(children),
	)
}