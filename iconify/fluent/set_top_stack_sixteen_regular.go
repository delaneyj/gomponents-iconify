package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetTopStackSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M6.42 3.729a.5.5 0 0 0-.84 0l-.634.983l-1.04.2a.5.5 0 0 0-.25.853l.776.734l-.41 1.356a.5.5 0 0 0 .745.568L6 7.645l1.233.778a.5.5 0 0 0 .746-.568L7.569 6.5l.775-.734a.5.5 0 0 0-.25-.854l-1.04-.199l-.634-.983zm-.75 1.705L6 4.923l.33.511a.5.5 0 0 0 .326.22l.291.056l-.29.275a.5.5 0 0 0-.136.508l.112.37l-.366-.232a.5.5 0 0 0-.534 0l-.366.232l.112-.37a.5.5 0 0 0-.135-.508l-.291-.275l.291-.056a.5.5 0 0 0 .326-.22z" fill="currentColor"/><path d="M2.5 2A1.5 1.5 0 0 0 1 3.5v5A1.5 1.5 0 0 0 2.5 10h7A1.5 1.5 0 0 0 11 8.5v-5A1.5 1.5 0 0 0 9.5 2h-7zm7 1a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h7z" fill="currentColor"/><path d="M4.5 12a1.5 1.5 0 0 1-1.415-1H10a2 2 0 0 0 2-2V4.085A1.5 1.5 0 0 1 13 5.5V9a3 3 0 0 1-3 3H4.5z" fill="currentColor"/><path d="M6.5 14a1.5 1.5 0 0 1-1.415-1H10.5A3.5 3.5 0 0 0 14 9.5V6.085A1.5 1.5 0 0 1 15 7.5v2a4.5 4.5 0 0 1-4.5 4.5h-4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}