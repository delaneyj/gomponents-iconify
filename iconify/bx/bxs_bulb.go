package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsBulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M9 20h6v2H9zm7.906-6.288C17.936 12.506 19 11.259 19 9c0-3.859-3.141-7-7-7S5 5.141 5 9c0 2.285 1.067 3.528 2.101 4.73c.358.418.729.851 1.084 1.349c.144.206.38.996.591 1.921h-.792v2h8.032v-2h-.79c.213-.927.45-1.719.593-1.925c.352-.503.726-.94 1.087-1.363z" fill="currentColor"/>`),
		g.Group(children),
	)
}