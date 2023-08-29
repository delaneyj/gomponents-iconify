package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhosphorLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 96a76.08 76.08 0 0 0-76-76H64a12 12 0 0 0-12 12v128a84.09 84.09 0 0 0 84 84a12 12 0 0 0 12-12v-60.11A76.09 76.09 0 0 0 220 96ZM76 77.81L115.48 148H76Zm48 36.38L84.52 44H124ZM77.22 172H124v46.79A60.18 60.18 0 0 1 77.22 172ZM148 147.83V44.17a52 52 0 0 1 0 103.66Z"/>`),
		g.Group(children),
	)
}