package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archives(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 2H5v4h6V2zM9 4H7V3h2v1z"/><path fill="currentColor" d="M3 0v16h2v-1h6v1h2V0H3zm9 14H4V8h8v6zm0-7H4V1h8v6z"/><path fill="currentColor" d="M11 9H5v4h6V9zm-2 2H7v-1h2v1z"/>`),
		g.Group(children),
	)
}