package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailUnread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2h-6v6h6V2zM4 4h10v2H4v12h16v-8h2v10H2V4h2zm4 4H6v2h2v2h2v2h4v-2h2v-2h-2v2h-4v-2H8V8z"/>`),
		g.Group(children),
	)
}