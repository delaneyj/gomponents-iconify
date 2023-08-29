package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MistLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h4v2H4V4Zm12 15h4v2h-4v-2ZM2 9h5v2H2V9Zm7 0h3v2H9V9Zm5 0h6v2h-6V9ZM4 14h6v2H4v-2Zm8 0h3v2h-3v-2Zm5 0h5v2h-5v-2ZM10 4h12v2H10V4ZM2 19h12v2H2v-2Z"/>`),
		g.Group(children),
	)
}