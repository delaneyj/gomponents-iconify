package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 8h-4V4H8v4H4v2h4v4h2v-4h4V8zM4 19h10v2H4zm0 5h10v2H4zM18 8h10v2H18zm6.41 14L28 18.41L26.59 17L23 20.59L19.41 17L18 18.41L21.59 22L18 25.59L19.41 27L23 23.41L26.59 27L28 25.59L24.41 22z"/>`),
		g.Group(children),
	)
}