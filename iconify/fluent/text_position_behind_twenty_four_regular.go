package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionBehindTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4.75A.75.75 0 0 1 3.75 4h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 4.75ZM9.5 11a2.5 2.5 0 0 1 5 0v4.75a.75.75 0 0 0 1.5 0V11a4 4 0 0 0-8 0v4.75a.75.75 0 0 0 1.5 0V11Zm10.75 1.5H17V11h3.25a.75.75 0 0 1 0 1.5Zm-16.5 0H7V11H3.75a.75.75 0 0 0 0 1.5Zm0 5.5a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 0-1.5H3.75Zm6.75-7h3v1.5h-3V11Z"/>`),
		g.Group(children),
	)
}