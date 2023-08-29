package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m205.66 61.66l-144 144a8 8 0 0 1-11.32-11.32l144-144a8 8 0 0 1 11.32 11.32ZM64 112a8 8 0 0 0 16 0V80h32a8 8 0 0 0 0-16H80V32a8 8 0 0 0-16 0v32H32a8 8 0 0 0 0 16h32Zm160 64h-80a8 8 0 0 0 0 16h80a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}