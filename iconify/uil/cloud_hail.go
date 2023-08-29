package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudHail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 2a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 2a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm2.42-4.78A7 7 0 0 0 5.06 8.11A4 4 0 0 0 2 12a4 4 0 0 0 1.34 3a1 1 0 1 0 1.32-1.5A2 2 0 0 1 4 12a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 11a2.91 2.91 0 0 1-.74 2a1 1 0 0 0 1.48 1.34a5 5 0 0 0-2.32-8.08ZM16 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}