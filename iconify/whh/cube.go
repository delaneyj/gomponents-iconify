package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1007V559l448-239v448zM0 256L480 0l480 256l-480 256zm448 751L0 768V320l448 239v448z"/>`),
		g.Group(children),
	)
}