package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16ZM96 120a32 32 0 1 1 32 32a32 32 0 0 1-32-32Zm-27.33 88a64.36 64.36 0 0 1 19.13-25.8a64 64 0 0 1 80.4 0a64.36 64.36 0 0 1 19.13 25.8ZM208 208h-3.67a79.9 79.9 0 0 0-46.68-50.29a48 48 0 1 0-59.3 0A79.9 79.9 0 0 0 51.67 208H48V48h160v160Z"/>`),
		g.Group(children),
	)
}