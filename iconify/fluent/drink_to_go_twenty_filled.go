package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrinkToGoTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.853 2.146a.5.5 0 0 1 0 .708L12.706 4h1.624a.75.75 0 0 1 .704.492l.917 2.5A.75.75 0 0 1 15.247 8h-.339l-1.425 8.337A2 2 0 0 1 11.512 18H8.487a2 2 0 0 1-1.972-1.663L5.09 8h-.34a.75.75 0 0 1-.703-1.01l.923-2.5A.75.75 0 0 1 5.674 4h5.642a.5.5 0 0 1 .08-.104l1.75-1.75a.5.5 0 0 1 .707 0ZM5.109 7h9.78l-.734-2H5.848l-.739 2Z"/>`),
		g.Group(children),
	)
}