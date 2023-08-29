package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rollerblade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 18a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 0-5m14 0a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 0-5m1.5-1c0-1.37.09-3.19-.5-4.05c-.95-2.05-3.68-2.39-5.59-2.9C13 10 12 9 11.82 8H9a.49.49 0 0 1-.5-.5c0-.28.21-.5.5-.5h2.5V6H9a.49.49 0 0 1-.5-.5c0-.28.21-.5.5-.5h2.5V2H3v15h17m-8.5 1a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 0-5Z"/>`),
		g.Group(children),
	)
}