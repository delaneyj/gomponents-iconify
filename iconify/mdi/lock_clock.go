package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 2C6 2 4 4 4 6.5V7c-1.11 0-2 .89-2 2v9c0 1.11.89 2 2 2h4.72c1.46 1.29 3.34 2 5.28 2a8 8 0 0 0 8-8a8 8 0 0 0-8-8c-.34 0-.68.03-1 .08C12.76 3.77 10.82 2 8.5 2m0 2A2.5 2.5 0 0 1 11 6.5V7H6v-.5A2.5 2.5 0 0 1 8.5 4M14 8a6 6 0 0 1 6 6a6 6 0 0 1-6 6a6 6 0 0 1-6-6a6 6 0 0 1 6-6m-1 2v5l3.64 2.19l.78-1.29l-2.92-1.75V10H13Z"/>`),
		g.Group(children),
	)
}