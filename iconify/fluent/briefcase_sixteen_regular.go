package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 5V3.5A1.5 1.5 0 0 1 6.5 2h3A1.5 1.5 0 0 1 11 3.5V5h1a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h1Zm1-1.5V5h4V3.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5Zm-3 6V12a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V9.5c-.418.314-.937.5-1.5.5H9v.5a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5V10H4.5A2.489 2.489 0 0 1 3 9.5ZM7 9v-.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5V9h2.5A1.5 1.5 0 0 0 13 7.5V7a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v.5A1.5 1.5 0 0 0 4.5 9H7Z"/>`),
		g.Group(children),
	)
}