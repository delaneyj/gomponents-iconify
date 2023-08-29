package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 10.19V18h2v-6h6v6h2v-7.81l-5-4.5l-5 4.5zm7-.19h-4c0-1.1.9-2 2-2s2 .9 2 2z" opacity=".3"/><path fill="currentColor" d="M19 9.3V4h-3v2.6L12 3L2 12h3v8h6v-6h2v6h6v-8h3l-3-2.7zM17 18h-2v-6H9v6H7v-7.81l5-4.5l5 4.5V18z"/><path fill="currentColor" d="M10 10h4c0-1.1-.9-2-2-2s-2 .9-2 2z"/>`),
		g.Group(children),
	)
}