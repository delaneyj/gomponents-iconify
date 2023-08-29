package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M3 0C2 0 1 .531 1 2v7.5s0 .5.5.5l.5.016v.484s0 .5.5.5H3s.5 0 .5-.5v-.484l4-.016v.5s0 .5.5.5h.5c.5 0 .5-.5.5-.5v-.484L9.5 10s.5 0 .5-.5V2c0-1.5-1-2-2-2H3zm.75 1h3.5a.25.25 0 1 1 0 .5h-3.5a.25.25 0 1 1 0-.5zM3 2h5c1 0 1 1 1 1v2s0 1-1 1H3C2 6 2 5 2 5V3s0-1 1-1zm-.25 5.5a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zm5.5 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}