package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M32 160L166.394 26.643a4.001 4.001 0 0 1 5.654.026l57.837 58.237a4.034 4.034 0 0 1-.007 5.676L97.348 223.59L32 224v-64zm16.797 5.594V208h40.488l121.92-121.396L180.57 56.56L64.656 175.772a3.937 3.937 0 0 1-5.624.037l-10.235-10.215z"/>`),
		g.Group(children),
	)
}