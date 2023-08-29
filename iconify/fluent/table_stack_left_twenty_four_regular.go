package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 3.75a.75.75 0 0 0-1.5 0v16.5a.75.75 0 0 0 1.5 0V3.75Zm4 0A.75.75 0 0 1 9.25 3h8.5A3.25 3.25 0 0 1 21 6.25v11.5A3.25 3.25 0 0 1 17.75 21h-8.5a.75.75 0 0 1-.75-.75V3.75ZM10 10v4h4v-4h-4Zm0 5.5v4h4v-4h-4Zm5.5 0v4h2.25a1.75 1.75 0 0 0 1.75-1.75V15.5h-4Zm0-5.5v4h4v-4h-4Zm4-1.5V6.25a1.75 1.75 0 0 0-1.75-1.75H15.5v4h4Zm-9.5-4v4h4v-4h-4Z"/>`),
		g.Group(children),
	)
}