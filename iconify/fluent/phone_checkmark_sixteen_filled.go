package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneCheckmarkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 11a5.5 5.5 0 0 0 1.5-.207v2.457A1.75 1.75 0 0 1 10.25 15h-4.5A1.75 1.75 0 0 1 4 13.25V2.75C4 1.784 4.784 1 5.75 1h1.587A5.5 5.5 0 0 0 10.5 11ZM7 12a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1H7Zm8-6.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-2.146-1.854a.5.5 0 0 0-.708 0L9.5 6.293l-.646-.647a.5.5 0 1 0-.708.708l1 1a.5.5 0 0 0 .708 0l3-3a.5.5 0 0 0 0-.708Z"/>`),
		g.Group(children),
	)
}