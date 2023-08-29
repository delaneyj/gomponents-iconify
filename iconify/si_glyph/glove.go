package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.807 9.708c1.863-1.863 3.68-4.907.307-8.28a4.771 4.771 0 0 0-6.747 0L6.342 2.453v-.719c0-.958-.777-1.735-1.735-1.735C3.648 0 3 .777 3 1.735v4.188l-.847.719l-.613-.613L.313 7.256l7.974 7.974l1.227-1.227l-.549-.549l4.842-3.746z"/>`),
		g.Group(children),
	)
}