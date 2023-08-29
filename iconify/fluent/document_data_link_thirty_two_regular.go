package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDataLinkThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a3 3 0 0 0-3 3v14.689a6.01 6.01 0 0 1 1.5-.189H7V5a1 1 0 0 1 1-1h9v5a3 3 0 0 0 3 3h5v15a1 1 0 0 1-1 1h-6.044a6.018 6.018 0 0 1-1.487 2H24a3 3 0 0 0 3-3V10.828a3 3 0 0 0-.879-2.12l-5.828-5.83A3 3 0 0 0 18.172 2H8Zm2 11v6.5h2V13a1 1 0 1 0-2 0Zm5 7.044a6.018 6.018 0 0 1 2 1.487V19a1 1 0 1 0-2 0v1.044ZM24.586 10H20a1 1 0 0 1-1-1V4.414L24.586 10ZM21 15a1 1 0 0 1 1 1v9a1 1 0 1 1-2 0v-9a1 1 0 0 1 1-1ZM8 22a1 1 0 0 0-1-1h-.5a4.5 4.5 0 1 0 0 9H7a1 1 0 1 0 0-2h-.5a2.5 2.5 0 0 1 0-5H7a1 1 0 0 0 1-1Zm4-1a1 1 0 1 0 0 2h.5a2.5 2.5 0 0 1 0 5H12a1 1 0 1 0 0 2h.5a4.5 4.5 0 1 0 0-9H12Zm-5 3.5a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2H7Z"/>`),
		g.Group(children),
	)
}