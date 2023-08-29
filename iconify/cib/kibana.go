package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kibana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.052 31.984H5.573l13.172-15.812c5.64 3.683 9.307 9.391 9.307 15.812zm0-31.932H4.104v28.735z"/>`),
		g.Group(children),
	)
}