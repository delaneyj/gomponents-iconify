package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FpsSixtySixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2.75a.5.5 0 0 1 1 0V3A.75.75 0 0 0 7 3v-.25a2 2 0 1 0-4 0V6a2 2 0 1 0 1.5-1.937V2.75ZM4.5 6a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM10 1a2 2 0 0 0-2 2v3a2 2 0 1 0 4 0V3a2 2 0 0 0-2-2Zm.5 5a.5.5 0 0 1-1 0V3a.5.5 0 0 1 1 0v3Zm-9 4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 1 0V13h1.5a.5.5 0 0 0 0-1H2v-1h2.5a.5.5 0 0 0 0-1h-3Zm5 0a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 1 0v-1h1.25a1.75 1.75 0 1 0 0-3.5H6.5Zm1.75 2.5H7V11h1.25a.75.75 0 0 1 0 1.5Zm2.75-1a1.5 1.5 0 0 1 1.5-1.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 0 0 1h1a1.5 1.5 0 0 1 0 3h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 0 0-1h-1a1.5 1.5 0 0 1-1.5-1.5Z"/>`),
		g.Group(children),
	)
}