package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#080d98"/><path fill="#fff" d="m23 6l-12.029 8.25l6.076 3.875L9 26l13.302-9.208l-5.994-3.875z"/></g>`),
		g.Group(children),
	)
}