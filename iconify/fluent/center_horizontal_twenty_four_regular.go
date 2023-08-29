package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterHorizontalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 3.75a.75.75 0 0 0-1.5 0v16.5a.75.75 0 0 0 1.5 0V3.75Zm15 0a.75.75 0 0 1 1.5 0v16.5a.75.75 0 0 1-1.5 0V3.75ZM10.25 5A2.25 2.25 0 0 0 8 7.25v9.5A2.25 2.25 0 0 0 10.25 19h3.5A2.25 2.25 0 0 0 16 16.75v-9.5A2.25 2.25 0 0 0 13.75 5h-3.5ZM9.5 7.25a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 .75.75v9.5a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1-.75-.75v-9.5Z"/>`),
		g.Group(children),
	)
}