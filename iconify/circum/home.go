package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.37 10.22l-6.2-7.6a1.5 1.5 0 0 0-2.33-.01l-6.21 7.61a2.5 2.5 0 0 0-.57 1.59v7.63a2.507 2.507 0 0 0 2.5 2.5h10.88a2.507 2.507 0 0 0 2.5-2.5v-7.63a2.5 2.5 0 0 0-.57-1.59ZM10 20.94v-5.5a1.5 1.5 0 0 1 1.5-1.5h1a1.5 1.5 0 0 1 1.5 1.5v5.5Zm8.94-1.5a1.511 1.511 0 0 1-1.5 1.5H15v-5.5a2.5 2.5 0 0 0-2.5-2.5h-1a2.5 2.5 0 0 0-2.5 2.5v5.5H6.56a1.511 1.511 0 0 1-1.5-1.5v-7.63a1.474 1.474 0 0 1 .34-.95l6.22-7.61a.474.474 0 0 1 .38-.19a.479.479 0 0 1 .39.19l6.21 7.61a1.474 1.474 0 0 1 .34.95Z"/>`),
		g.Group(children),
	)
}