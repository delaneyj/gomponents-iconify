package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.73 16.795a.75.75 0 0 0 1.031-1.09L4.522 10.75h12.724a.75.75 0 0 0 0-1.5H4.52l5.241-4.957a.75.75 0 1 0-1.03-1.09l-6.417 6.07a.995.995 0 0 0-.3.566a.752.752 0 0 0 0 .329a.995.995 0 0 0 .3.558l6.417 6.07Z"/>`),
		g.Group(children),
	)
}