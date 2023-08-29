package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroservicesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 22v-6h-6v-6H2v20h20v-8h-6zm-2-4v4h-4v-4h4zM4 12h4v4H4v-4zm4 6v4H4v-4h4zM4 28v-4h4v4H4zm10 0h-4v-4h4v4zm6 0h-4v-4h4v4zm9.6-14.4L27 16.2V10h3V2h-8v8h3v6.2l-2.6-2.6L21 15l5 5l5-5l-1.4-1.4zM24 4h4v4h-4V4z"/>`),
		g.Group(children),
	)
}