package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 14.5c.468-2.207.985-4.05 1.604-5.846c.411 1.796.928 3.638 1.337 5.521C16 15.328 15.329 16 14.5 16s-1.5-.672-1.5-1.5zM8 1L6.56 2.44l-2-2a1.539 1.539 0 0 0-2.121 0a1.496 1.496 0 0 0 .001 2.119l2 2L0 8.999l7 7l8-8zm0 1.41L13.59 8H2.41l2.75-2.75a.49.49 0 0 0 .669-.672l.721-.718l1.54 1.53a.502.502 0 0 0 .71-.71L7.27 3.15zm-4.85-.56a.5.5 0 0 1 .355-.854c.138 0 .263.055.355.144l2 2l-.71.71z"/>`),
		g.Group(children),
	)
}