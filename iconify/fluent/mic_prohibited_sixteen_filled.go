package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicProhibitedSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4.5a2.5 2.5 0 0 1 5 0v.707a5.504 5.504 0 0 0-3.979 4.809A2.496 2.496 0 0 1 4 8V4.5Zm1.042 6.683A3.5 3.5 0 0 1 3 8a.5.5 0 0 0-1 0a4.502 4.502 0 0 0 3.316 4.343a5.466 5.466 0 0 1-.274-1.16ZM6 10.5a4.5 4.5 0 1 0 9 0a4.5 4.5 0 0 0-9 0Zm2.404 2.803l4.9-4.9a3.5 3.5 0 0 1-4.9 4.9Zm-.707-.707a3.5 3.5 0 0 1 4.9-4.9l-4.9 4.9Z"/>`),
		g.Group(children),
	)
}