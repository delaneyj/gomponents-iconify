package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDataThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a3 3 0 0 0-3 3v22a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V10.828a3 3 0 0 0-.879-2.12l-5.828-5.83A3 3 0 0 0 18.172 2H8ZM7 5a1 1 0 0 1 1-1h9v5a3 3 0 0 0 3 3h5v15a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V5Zm17.586 5H20a1 1 0 0 1-1-1V4.414L24.586 10ZM12 13a1 1 0 1 0-2 0v12a1 1 0 1 0 2 0V13Zm5 8a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4Zm4-5a1 1 0 0 1 1 1v8a1 1 0 1 1-2 0v-8a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}