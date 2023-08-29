package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardTextRtlTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.98 3.945A2.25 2.25 0 0 0 13.75 2h-3.5a2.25 2.25 0 0 0-2.23 1.945l-.014.135l.008-.08H6.25A2.25 2.25 0 0 0 4 6.25v13.5A2.25 2.25 0 0 0 6.25 22h11.5A2.25 2.25 0 0 0 20 19.75V6.25A2.25 2.25 0 0 0 17.75 4h-1.764l.007.08l-.014-.135ZM10.25 3.5h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1 0-1.5ZM8 9h8a.75.75 0 0 1 0 1.5H8A.75.75 0 0 1 8 9Zm3.25 4.75A.75.75 0 0 1 12 13h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75ZM10 17h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}