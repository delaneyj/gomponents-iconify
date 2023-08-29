package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 140h-36v-40h20a12 12 0 0 0 12-12V40a12 12 0 0 0-12-12H64a12 12 0 0 0-12 12v48a12 12 0 0 0 12 12h20v40H48a12 12 0 0 0-12 12v16a12 12 0 0 0 12 12h12v44a4 4 0 0 0 8 0v-44h120v44a4 4 0 0 0 8 0v-44h12a12 12 0 0 0 12-12v-16a12 12 0 0 0-12-12ZM60 88V40a4 4 0 0 1 4-4h128a4 4 0 0 1 4 4v48a4 4 0 0 1-4 4H64a4 4 0 0 1-4-4Zm32 12h72v40H92Zm120 68a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4v-16a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}