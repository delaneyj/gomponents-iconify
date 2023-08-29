package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pexels(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 5a5.001 5.001 0 0 1 0 10v4H6V5h6ZM8 7v10h2v-4h2l.003-.001a2.993 2.993 0 0 0 3.041-3.044l-.007-.39a2.61 2.61 0 0 0-2.711-2.562l-.306.005L12 7H8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}