package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrabberSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm0-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-4 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm5-9a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM7 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM6 5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}