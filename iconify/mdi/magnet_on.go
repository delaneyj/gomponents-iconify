package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7v6a9 9 0 0 0 9 9a9 9 0 0 0 9-9V7h-4v6a5 5 0 0 1-5 5a5 5 0 0 1-5-5V7m10-2h4V2h-4M3 5h4V2H3m10-.5L9 9h2v5.5L15 7h-2V1.5Z"/>`),
		g.Group(children),
	)
}