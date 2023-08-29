package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserFocusThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 48v28a4 4 0 0 1-8 0V48a4 4 0 0 0-4-4h-28a4 4 0 0 1 0-8h28a12 12 0 0 1 12 12Zm-4 128a4 4 0 0 0-4 4v28a4 4 0 0 1-4 4h-28a4 4 0 0 0 0 8h28a12 12 0 0 0 12-12v-28a4 4 0 0 0-4-4ZM76 212H48a4 4 0 0 1-4-4v-28a4 4 0 0 0-8 0v28a12 12 0 0 0 12 12h28a4 4 0 0 0 0-8ZM40 80a4 4 0 0 0 4-4V48a4 4 0 0 1 4-4h28a4 4 0 0 0 0-8H48a12 12 0 0 0-12 12v28a4 4 0 0 0 4 4Zm136 92a4 4 0 0 1-3.2-1.6a56 56 0 0 0-89.6 0a4 4 0 1 1-6.4-4.8a63.65 63.65 0 0 1 32.5-22.85a36 36 0 1 1 37.4 0a63.65 63.65 0 0 1 32.5 22.85a4 4 0 0 1-3.2 6.4Zm-48-32a28 28 0 1 0-28-28a28 28 0 0 0 28 28Z"/>`),
		g.Group(children),
	)
}