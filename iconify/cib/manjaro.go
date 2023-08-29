package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Manjaro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 0v32h9V9h11.5V0zm11.5 11.5V32h9V11.5zM23 0v32h9V0z"/>`),
		g.Group(children),
	)
}