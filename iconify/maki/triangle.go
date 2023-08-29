package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.538 1c-.294 0-.488.177-.615.385l-5.846 9.538C1 11 1 11.153 1 11.308c0 .538.385.692.692.692h11.616c.384 0 .692-.154.692-.692c0-.154 0-.231-.077-.385l-5.77-9.538C8.029 1.177 7.789 1 7.539 1z"/>`),
		g.Group(children),
	)
}