package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlassFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 112a56 56 0 1 1-56-56a56 56 0 0 1 56 56Zm61.66 117.66a8 8 0 0 1-11.32 0l-50.06-50.07a88 88 0 1 1 11.32-11.31l50.06 50.06a8 8 0 0 1 0 11.32ZM112 184a72 72 0 1 0-72-72a72.08 72.08 0 0 0 72 72Z"/>`),
		g.Group(children),
	)
}