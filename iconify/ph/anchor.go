package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 136a8 8 0 0 0-8 8a40 40 0 0 1-40 40a47.79 47.79 0 0 0-32 12.27V128h32a8 8 0 0 0 0-16h-32V87a32 32 0 1 0-16 0v25H88a8 8 0 0 0 0 16h32v68.27A47.79 47.79 0 0 0 88 184a40 40 0 0 1-40-40a8 8 0 0 0-16 0a56.06 56.06 0 0 0 56 56a32 32 0 0 1 32 32a8 8 0 0 0 16 0a32 32 0 0 1 32-32a56.06 56.06 0 0 0 56-56a8 8 0 0 0-8-8ZM112 56a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}