package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkSimpleHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 122h96a6 6 0 0 1 0 12H80a6 6 0 0 1 0-12Zm24 48H64a42 42 0 0 1 0-84h40a6 6 0 0 0 0-12H64a54 54 0 0 0 0 108h40a6 6 0 0 0 0-12Zm88-96h-40a6 6 0 0 0 0 12h40a42 42 0 0 1 0 84h-40a6 6 0 0 0 0 12h40a54 54 0 0 0 0-108Z"/>`),
		g.Group(children),
	)
}