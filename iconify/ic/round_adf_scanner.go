package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundAdfScanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 12h-1V6c0-1.1-.9-2-2-2H8c-1.1 0-2 .9-2 2v6H5c-1.66 0-3 1.34-3 3v3c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2v-3c0-1.66-1.34-3-3-3zm-3 0H8V6h8v6zm2 5c-.55 0-1-.45-1-1s.45-1 1-1s1 .45 1 1s-.45 1-1 1z"/>`),
		g.Group(children),
	)
}