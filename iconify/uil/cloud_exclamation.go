package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.71 18.79a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-1.09.21a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 .29.71a1.15 1.15 0 0 0 .33.21a1 1 0 0 0 1.3-1.3a1 1 0 0 0-.21-.33ZM12 12.5a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1Zm6.42-4.79a7 7 0 0 0-13.36 1.9A4 4 0 0 0 6 17.5h2a1 1 0 0 0 0-2H6a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.66A3 3 0 0 1 17 15.5h-1a1 1 0 0 0 0 2h1a5 5 0 0 0 1.42-9.79Z"/>`),
		g.Group(children),
	)
}