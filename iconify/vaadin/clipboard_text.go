package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 6h8v1H4V6zm0 2h8v1H4V8zm0 2h5v1H4v-1z"/><path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/>`),
		g.Group(children),
	)
}