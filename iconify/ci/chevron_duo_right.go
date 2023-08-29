package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDuoRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.414 18.01l-1.415-1.413l4.6-4.6l-4.6-4.6l1.415-1.407l6.01 6.01l-6.009 6.01h-.001Zm-5.425 0l-1.414-1.413l4.6-4.6l-4.6-4.593L6.989 5.99L13 12l-6.01 6.01h-.001Z"/>`),
		g.Group(children),
	)
}