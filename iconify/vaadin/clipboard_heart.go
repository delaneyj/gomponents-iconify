package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.5 7c-.6 0-1.1.6-1.5 1c-.4-.4-.9-1-1.5-1c-1.5 0-2.1 1.9-1 2.9L8 12l2.5-2.1c1.1-1 .5-2.9-1-2.9z"/><path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/>`),
		g.Group(children),
	)
}