package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.745 5a2.25 2.25 0 0 1 2.25 2.25v9.505a2.25 2.25 0 0 1-2.25 2.25H4.25A2.25 2.25 0 0 1 2 16.755V7.25A2.25 2.25 0 0 1 4.25 5h15.495Zm-2.495 9.5H6.75l-.102.007a.75.75 0 0 0 0 1.486L6.75 16h10.5l.102-.007a.75.75 0 0 0 0-1.486l-.102-.007ZM16.5 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-2.995 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM6 8a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm2.995 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}