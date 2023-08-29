package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeglassesThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 44a4 4 0 0 0 0 8a20 20 0 0 1 20 20v68a40 40 0 0 0-71.2 16h-41.6A40 40 0 0 0 36 140V72a20 20 0 0 1 20-20a4 4 0 0 0 0-8a28 28 0 0 0-28 28v92a40 40 0 0 0 80 0h40a40 40 0 0 0 80 0V72a28 28 0 0 0-28-28ZM68 196a32 32 0 1 1 32-32a32 32 0 0 1-32 32Zm120 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}