package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextboxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 80v96a16 16 0 0 1-16 16h-88a8 8 0 0 1-8-8V72a8 8 0 0 1 8-8h88a16 16 0 0 1 16 16ZM120 48v160a8 8 0 0 1-16 0v-16H24a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h80V48a8 8 0 0 1 16 0Zm-32 64a8 8 0 0 0-8-8H48a8 8 0 0 0 0 16h8v24a8 8 0 0 0 16 0v-24h8a8 8 0 0 0 8-8Z"/>`),
		g.Group(children),
	)
}