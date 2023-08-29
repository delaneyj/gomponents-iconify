package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineCharacterUTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 5a1 1 0 0 0-2 0v6.5a5 5 0 0 0 10 0V5a1 1 0 1 0-2 0v6.5a3 3 0 1 1-6 0V5ZM7 18a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7Z"/>`),
		g.Group(children),
	)
}