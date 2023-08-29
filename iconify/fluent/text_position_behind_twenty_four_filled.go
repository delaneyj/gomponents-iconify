package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionBehindTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 3.75a1 1 0 0 0 0 2h16.5a1 1 0 1 0 0-2H3.75Zm7 7.25c0-.086.009-.17.025-.25h2.45c.016.08.025.164.025.25v1.75h-2.5V11Zm3.486-.25c.01.082.014.165.014.25v4.75a1 1 0 1 0 2 0V11a4.25 4.25 0 0 0-8.5 0v4.75a1 1 0 1 0 2 0V11a2.25 2.25 0 0 1 4.486-.25Zm6.014 2h-3V11c0-.084-.002-.167-.006-.25h3.006a1 1 0 1 1 0 2ZM6.75 11c0-.084.002-.167.006-.25H3.75a1 1 0 1 0 0 2h3V11Zm-4 7.75a1 1 0 0 1 1-1h16.5a1 1 0 1 1 0 2H3.75a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}