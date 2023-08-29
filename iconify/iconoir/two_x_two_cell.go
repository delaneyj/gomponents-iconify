package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoXTwoCell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M21 3.6V12h-9V3h8.4a.6.6 0 0 1 .6.6Zm0 16.8V12h-9v9h8.4a.6.6 0 0 0 .6-.6ZM3 12V3.6a.6.6 0 0 1 .6-.6H12v9H3Zm0 0v8.4a.6.6 0 0 0 .6.6H12v-9H3Z"/>`),
		g.Group(children),
	)
}