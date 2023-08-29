package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glassdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M8 2a3 3 0 0 0-3 3v11.5h3V5h11a3 3 0 0 0-3-3H8m8 5.5V19H5a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V7.5h-3z" fill="currentColor"/>`),
		g.Group(children),
	)
}