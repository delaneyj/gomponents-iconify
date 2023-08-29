package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletLandscapeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 82v348a18 18 0 0 0 18 18h476a18 18 0 0 0 18-18V82a18 18 0 0 0-18-18H18A18 18 0 0 0 0 82Zm448 330H64V100h384Z"/>`),
		g.Group(children),
	)
}