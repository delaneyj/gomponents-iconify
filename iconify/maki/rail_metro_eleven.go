package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RailMetroEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M3.5 0C3 0 3 .5 3 .5L2 5v1c0 1.024 1 1 1 1h5s1 0 1-1V5L8 .5S8 0 7.5 0h-4zM4 1h3l.5 3h-4L4 1zm-.5 4a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zm1.75 0h.5a.25.25 0 1 1 0 .5h-.5a.25.25 0 1 1 0-.5zM7.5 5a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zM3 8l-1 3h1.5l.334-1h3.332l.334 1H9L8 8H6.5l.334 1H4.166L4.5 8H3z" fill="currentColor"/>`),
		g.Group(children),
	)
}