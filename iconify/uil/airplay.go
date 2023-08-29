package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.83 13.45a1 1 0 0 0-1.66 0l-4 6a1 1 0 0 0 0 1A1 1 0 0 0 8 21h8a1 1 0 0 0 .88-.53a1 1 0 0 0-.05-1ZM9.87 19L12 15.8l2.13 3.2ZM19 3H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h.85a1 1 0 1 0 0-2H5a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-.85a1 1 0 0 0 0 2H19a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}