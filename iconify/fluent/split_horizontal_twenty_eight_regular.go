package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 12.998a.75.75 0 0 0 0 1.5h22.5a.75.75 0 0 0 0-1.5H2.75ZM24 4.75A2.75 2.75 0 0 0 21.25 2H6.75A2.75 2.75 0 0 0 4 4.75v7.248h1.5V4.75c0-.691.56-1.25 1.25-1.25h14.5c.69 0 1.25.559 1.25 1.25v7.247H24V4.75ZM5.5 22.749V15.5H4v7.25a2.75 2.75 0 0 0 2.75 2.75h14.5A2.75 2.75 0 0 0 24 22.75v-7.252h-1.5v7.252c0 .69-.56 1.25-1.25 1.25H6.75c-.69 0-1.25-.56-1.25-1.25Z"/>`),
		g.Group(children),
	)
}