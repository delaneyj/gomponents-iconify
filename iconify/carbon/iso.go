package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 23h-4a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2zm-4-12v10h4V11zm-6 12h-6v-2h6v-4h-4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h6v2h-6v4h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2zM2 11h3v10H2v2h8v-2H7V11h3V9H2v2z"/>`),
		g.Group(children),
	)
}