package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideSearchTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.75 4A2.75 2.75 0 0 0 2 6.75v4.507a5.496 5.496 0 0 1 1.5-.882V6.75c0-.69.56-1.25 1.25-1.25h14.5c.69 0 1.25.56 1.25 1.25v10.5c0 .69-.56 1.25-1.25 1.25h-8.275l1.5 1.5h6.775A2.75 2.75 0 0 0 22 17.25V6.75A2.75 2.75 0 0 0 19.25 4H4.75Zm3.913 7a5.529 5.529 0 0 1 1.447 1.5h7.14a.75.75 0 0 0 0-1.5H8.663Zm6.587 4.5H11a5.5 5.5 0 0 0-.207-1.5h4.457a.75.75 0 0 1 0 1.5ZM6.75 8a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5h-6.5ZM5.5 20a4.48 4.48 0 0 0 2.607-.832l2.613 2.612a.75.75 0 1 0 1.06-1.06l-2.612-2.613A4.5 4.5 0 1 0 5.5 20Zm0-1.5a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		g.Group(children),
	)
}