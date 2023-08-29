package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouchThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 108.7V72a12 12 0 0 0-12-12H32a12 12 0 0 0-12 12v36.7a12 12 0 0 0-8 11.3v48a12 12 0 0 0 12 12h12v20a4 4 0 0 0 8 0v-20h168v20a4 4 0 0 0 8 0v-20h12a12 12 0 0 0 12-12v-48a12 12 0 0 0-8-11.3ZM228 72v36h-12a12 12 0 0 0-12 12v16a4 4 0 0 1-4 4h-68V68h92a4 4 0 0 1 4 4ZM32 68h92v72H56a4 4 0 0 1-4-4v-16a12 12 0 0 0-12-12H28V72a4 4 0 0 1 4-4Zm204 100a4 4 0 0 1-4 4H24a4 4 0 0 1-4-4v-48a4 4 0 0 1 4-4h16a4 4 0 0 1 4 4v16a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12v-16a4 4 0 0 1 4-4h16a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}