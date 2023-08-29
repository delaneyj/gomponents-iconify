package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.537 7.348c-.65-1.275-3.148-2.226-6.143-2.226c-3.486 0-6.312 1.289-6.312 2.879s2.825 2.879 6.312 2.879c2.994 0 5.492-.951 6.143-2.226c.659.841 1.953 2.191 3.437 2.191V5.156c-1.484 0-2.778 1.351-3.437 2.192zm-10.494.777H1V6.917h1.043v1.208z"/>`),
		g.Group(children),
	)
}