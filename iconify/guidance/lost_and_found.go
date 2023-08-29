package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LostAndFound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 11v-1a2.5 2.5 0 0 1 5 0v.039a2 2 0 0 1-.75 1.562l-1 .798a2 2 0 0 0-.75 1.562V15.5m0 1.5v2M5.5 4.5v17m13 0v-17h-2.027m-8.946 0H.5v.25l.055.31a45.684 45.684 0 0 1 0 15.88l-.055.31v.25h23v-.25l-.055-.31a45.686 45.686 0 0 1 0-15.88l.055-.31V4.5h-7.027m-8.946 0a4.5 4.5 0 0 1 8.945 0m-8.945 0h8.945"/>`),
		g.Group(children),
	)
}