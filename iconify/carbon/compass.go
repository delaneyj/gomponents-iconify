package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4A12 12 0 1 1 4 16A12 12 0 0 1 16 4m0-2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Z"/><path fill="currentColor" d="M23 10.41L21.59 9l-4.3 4.3a3 3 0 0 0-4 4L9 21.59L10.41 23l4.3-4.3a3 3 0 0 0 4-4ZM17 16a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/><circle cx="16" cy="7.5" r="1.5" fill="currentColor"/>`),
		g.Group(children),
	)
}