package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M75.846 176a4 4 0 0 1 4.002 4.002V203.6a4.002 4.002 0 0 1-4.002 4.003H52.001A4 4 0 0 1 48 203.599v-23.597A4.002 4.002 0 0 1 52.001 176h23.845z"/>`),
		g.Group(children),
	)
}