package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CharacterWholeNumber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 9h-6v2h6v4h-4v2h4v4h-6v2h6a2.003 2.003 0 0 0 2-2V11a2.002 2.002 0 0 0-2-2zm-8 14h-8v-6a2.002 2.002 0 0 1 2-2h4v-4h-6V9h6a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2h-4v4h6zm-17.5-.5v-1h3v-11h-3v-1h4v12h3v1h-7z"/><path fill="currentColor" d="M6 10v12v-12m1-1H2v2h3v10H2v2h8v-2H7V9Z"/>`),
		g.Group(children),
	)
}