package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm14-4h-7a1 1 0 0 0 0 2h7a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 0 0 0-6Zm-5-3a1 1 0 0 0 1 1h4a3 3 0 0 0 3-3a1 1 0 0 0-2 0a1 1 0 0 1-1 1h-4a1 1 0 0 0-1 1Zm-4 4a1 1 0 0 0-1-1H6a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 1 0 1.9-.64A7 7 0 0 0 5.06 8.11A4 4 0 0 0 6 16h3a1 1 0 0 0 1-1Zm0-4a1 1 0 1 0 1-1a1 1 0 0 0-1 1Zm4 7H9a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}