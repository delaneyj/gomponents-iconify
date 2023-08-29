package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookieFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 120a40 40 0 0 1-40-40a8 8 0 0 0-8-8a40 40 0 0 1-40-40a8 8 0 0 0-8-8a104 104 0 1 0 104 104a8 8 0 0 0-8-8ZM75.51 99.51a12 12 0 1 1 0 17a12 12 0 0 1 0-17Zm25 73a12 12 0 1 1 0-17a12 12 0 0 1-.02 16.98Zm23-40a12 12 0 1 1 17 0a12 12 0 0 1-17-.02Zm41 48a12 12 0 1 1 0-17a12 12 0 0 1-.02 16.98Z"/>`),
		g.Group(children),
	)
}