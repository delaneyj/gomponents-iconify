package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 48C141.31 48 48 141.31 48 256s93.31 208 208 208s208-93.31 208-208S370.69 48 256 48Zm32 304h-64a16 16 0 0 1 0-32h64a16 16 0 0 1 0 32Zm48-64H176a16 16 0 0 1 0-32h160a16 16 0 0 1 0 32Zm32-64H144a16 16 0 0 1 0-32h224a16 16 0 0 1 0 32Z"/>`),
		g.Group(children),
	)
}