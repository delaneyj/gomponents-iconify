package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CinemaEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10 5.5v2a.5.5 0 0 1-1 0a.66.66 0 0 0-.51-.5H8v1.63a.37.37 0 0 1-.37.37H1.37A.37.37 0 0 1 1 8.63V5.37A.37.37 0 0 1 1.37 5h6.26a.37.37 0 0 1 .37.37V6h.49A.66.66 0 0 0 9 5.5a.5.5 0 0 1 1 0zM2.5 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM6 1a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}