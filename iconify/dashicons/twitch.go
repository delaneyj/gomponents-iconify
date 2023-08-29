package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.7 2L2 4.6v11.8h3.2V18H7l1.8-1.6h2.9l5.7-5.2V2H2.7zM16 10.5l-2.5 2.3h-4l-2.2 2v-2H4.2V3.3H16v7.2zm-2.5-4.6h-1.4v3.9h1.4V5.9zm-4 0H8.1v3.9h1.4V5.9z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}