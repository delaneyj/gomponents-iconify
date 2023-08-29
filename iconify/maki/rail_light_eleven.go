package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RailLightEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4 0c-.5 0-.5.5-.5.5s0 .5.5.5h1v.97L4 2S3 2 3 3v3c0 2 2 2 2 2h1s2 0 2-2V3c0-1-1-1-1-1H6V1h1c.5 0 .5-.5.5-.5S7.5 0 7 0H4zm1.5 3l1.5.5V5H4V3.5L5.5 3zm0 3a.499.499 0 1 1 0 1a.499.499 0 1 1 0-1zM2.834 8.5L2 11h1.5l.334-1h3.332l.334 1H9l-.834-2.5H6.668l.166.5H4.166l.166-.5H2.834z" fill="currentColor"/>`),
		g.Group(children),
	)
}