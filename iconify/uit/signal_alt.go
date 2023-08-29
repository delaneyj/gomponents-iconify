package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 18a.5.5 0 0 0-.5.5v3a.5.5 0 1 0 1 0v-3a.5.5 0 0 0-.5-.5zm5-4a.5.5 0 0 0-.5.5v7a.5.5 0 1 0 1 0v-7a.5.5 0 0 0-.5-.5zm10-12a.5.5 0 0 0-.5.5v19a.5.5 0 1 0 1 0v-19a.5.5 0 0 0-.5-.5zm-5 7a.5.5 0 0 0-.5.5v12a.5.5 0 1 0 1 0v-12a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}