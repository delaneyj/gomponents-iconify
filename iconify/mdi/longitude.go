package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Longitude(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.03 10.03 0 0 0 12 2M9.4 19.6a8.05 8.05 0 0 1 0-15.2A16.45 16.45 0 0 0 7.5 12a16.45 16.45 0 0 0 1.9 7.6m2.6.4a13.81 13.81 0 0 1-2.5-8A13.81 13.81 0 0 1 12 4a13.81 13.81 0 0 1 2.5 8a13.81 13.81 0 0 1-2.5 8m2.6-.4a16.15 16.15 0 0 0 0-15.2A8.03 8.03 0 0 1 20 12a7.9 7.9 0 0 1-5.4 7.6Z"/>`),
		g.Group(children),
	)
}