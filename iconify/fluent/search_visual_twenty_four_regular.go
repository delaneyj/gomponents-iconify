package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchVisualTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.75 3A3.25 3.25 0 0 1 21 6.25v3a.75.75 0 0 1-1.5 0v-3a1.75 1.75 0 0 0-1.75-1.75h-3a.75.75 0 0 1 0-1.5h3ZM6.25 3A3.25 3.25 0 0 0 3 6.25v3a.75.75 0 0 0 1.5 0v-3c0-.966.784-1.75 1.75-1.75h3a.75.75 0 0 0 0-1.5h-3Zm11.5 18A3.25 3.25 0 0 0 21 17.75v-3a.75.75 0 0 0-1.5 0v3a1.75 1.75 0 0 1-1.75 1.75h-3a.75.75 0 0 0 0 1.5h3ZM3 17.75A3.25 3.25 0 0 0 6.25 21h3a.75.75 0 0 0 0-1.5h-3a1.75 1.75 0 0 1-1.75-1.75v-3a.75.75 0 0 0-1.5 0v3ZM12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-1.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-4.5-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}