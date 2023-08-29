package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.332 7A1.75 1.75 0 1 1 5.75 6h4.5a1.75 1.75 0 1 1-1.582 1H7.332ZM5 7.75a.75.75 0 1 0 1.5 0a.75.75 0 0 0-1.5 0ZM10.25 7a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5ZM3.5 3A1.5 1.5 0 0 0 2 4.5V11a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 14 11V4.5A1.5 1.5 0 0 0 12.5 3h-9ZM3 4.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5V11a.5.5 0 0 1-.5.5h-9A.5.5 0 0 1 3 11V4.5Z"/>`),
		g.Group(children),
	)
}