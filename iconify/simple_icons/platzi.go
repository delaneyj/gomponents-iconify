package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Platzi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.64 1.127L2.485 9.282a3.842 3.842 0 0 0 0 5.436l8.155 8.155a3.842 3.842 0 0 0 5.436 0l2.72-2.718l-2.72-2.718l-2.718 2.718L5.203 12l8.155-8.155l5.437 5.437l-5.437 5.436l2.718 2.72L21.513 12a3.842 3.842 0 0 0 0-5.437l-5.448-5.436a3.828 3.828 0 0 0-5.425 0z"/>`),
		g.Group(children),
	)
}