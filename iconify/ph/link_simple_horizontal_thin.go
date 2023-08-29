package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkSimpleHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 124h96a4 4 0 0 1 0 8H80a4 4 0 0 1 0-8Zm24 48H64a44 44 0 0 1 0-88h40a4 4 0 0 0 0-8H64a52 52 0 0 0 0 104h40a4 4 0 0 0 0-8Zm88-96h-40a4 4 0 0 0 0 8h40a44 44 0 0 1 0 88h-40a4 4 0 0 0 0 8h40a52 52 0 0 0 0-104Z"/>`),
		g.Group(children),
	)
}