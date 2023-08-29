package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 2a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4Zm7.656 1.929a.5.5 0 1 0-.707.707l1.414 1.414a.5.5 0 1 0 .707-.707L15.156 3.93ZM9.5 6a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 1 0v-5a.5.5 0 0 0-.5-.5Zm0 12a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0-1a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`),
		g.Group(children),
	)
}