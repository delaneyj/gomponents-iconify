package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.064 9h13.872C18.45 5.255 15.25 3 12 3S5.55 5.255 5.064 9ZM19 11h-2v9h2v-9Zm2 9h1v2h-8v-2h1v-9H9v9h1v2H2v-2h1V10c0-5.6 4.529-9 9-9s9 3.4 9 9v10ZM7 11H5v9h2v-9Zm3-6h4v2h-4V5Z"/>`),
		g.Group(children),
	)
}