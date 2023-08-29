package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmericanFootball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.53 3C3.09 3 1 7.5 1 7.5S3.09 12 7.53 12S14 7.5 14 7.5S12 3 7.53 3zM11 7v1.5a.5.5 0 0 1-1 0V8H8v.5a.5.5 0 0 1-1 0V8H5v.5a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 1 0V7h2v-.5a.5.5 0 0 1 1 0V7h2v-.5a.5.5 0 0 1 1 0V7z"/>`),
		g.Group(children),
	)
}