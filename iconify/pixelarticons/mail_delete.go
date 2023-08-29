package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4H2v16h12v-2H4V6h16v8h2V4h-2zM6 8h2v2H6V8zm4 4H8v-2h2v2zm4 0v2h-4v-2h4zm2-2v2h-2v-2h2zm0 0V8h2v2h-2zm2 6h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2h2v-2h-2v2h-2v-2z"/>`),
		g.Group(children),
	)
}