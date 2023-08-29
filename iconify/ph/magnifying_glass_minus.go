package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlassMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 112a8 8 0 0 1-8 8H80a8 8 0 0 1 0-16h64a8 8 0 0 1 8 8Zm77.66 117.66a8 8 0 0 1-11.32 0l-50.06-50.07a88.11 88.11 0 1 1 11.31-11.31l50.07 50.06a8 8 0 0 1 0 11.32ZM112 184a72 72 0 1 0-72-72a72.08 72.08 0 0 0 72 72Z"/>`),
		g.Group(children),
	)
}