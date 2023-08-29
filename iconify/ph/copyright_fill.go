package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyrightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 56a72 72 0 1 0 72 72a72.08 72.08 0 0 0-72-72Zm0 104a31.8 31.8 0 0 0 25.61-12.8a8 8 0 1 1 12.79 9.61a48 48 0 1 1 0-57.63a8 8 0 1 1-12.79 9.61A32 32 0 1 0 128 160Zm0-136a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Z"/>`),
		g.Group(children),
	)
}