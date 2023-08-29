package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StairsUpPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5.5 14V6.5H4.443a3 3 0 0 0-2.938 2.392L.5 13.75V14l5 2.5v.5c0 1.5 0 3.5.75 5c0 0 .75 1.5 1.75 1.5m-4.5-6c-.21 0-.466.133-.737.344C.958 19.24.5 21.718.5 24m9.5-.5h5.5v-.008a7.5 7.5 0 0 0-.988-3.721a.09.09 0 0 1-.012-.045V19.5h5v-.008a7.5 7.5 0 0 0-.988-3.721a.09.09 0 0 1-.012-.045V15.5H24M5.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.246 3.5 5.65 4.5 5.65 4.5h-.3Z"/>`),
		g.Group(children),
	)
}