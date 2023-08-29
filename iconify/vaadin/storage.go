package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 4L7.94 0L0 4v1h1v11h2V7h10v9h2V5h1V4zM4 6V5h2v1H4zm3 0V5h2v1H7zm3 0V5h2v1h-2z"/><path fill="currentColor" d="M6 9H5V8H4v3h3V8H6v1zm0 4H5v-1H4v3h3v-3H6v1zm4 0H9v-1H8v3h3v-3h-1v1z"/>`),
		g.Group(children),
	)
}