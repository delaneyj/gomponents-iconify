package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalColon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M76.846 80a4 4 0 0 1 4.002 4.002V107.6a4.002 4.002 0 0 1-4.002 4.003H53.001A4 4 0 0 1 49 107.599V84.002A4.002 4.002 0 0 1 53.001 80h23.845zm0 64a4 4 0 0 1 4.002 4.002V171.6a4.002 4.002 0 0 1-4.002 4.003H53.001A4 4 0 0 1 49 171.599v-23.597A4.002 4.002 0 0 1 53.001 144h23.845z"/>`),
		g.Group(children),
	)
}