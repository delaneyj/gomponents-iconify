package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 19H8c-1.654 0-3-1.346-3-3V8c0-1.654 1.346-3 3-3h5a1 1 0 1 1 0 2H8c-.552 0-1 .449-1 1v8c0 .551.448 1 1 1h8c.552 0 1-.449 1-1v-3a1 1 0 1 1 2 0v3c0 1.654-1.346 3-3 3zm-2.834-4.167c-.35 0-.689-.139-.941-.391l-2.668-2.667a1.334 1.334 0 0 1 1.887-1.885l1.416 1.417l3.475-5.455a1.334 1.334 0 1 1 2.332 1.294l-4.334 7a1.332 1.332 0 0 1-.98.673l-.187.014z"/>`),
		g.Group(children),
	)
}