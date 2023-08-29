package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TennisTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.34 17c2.761 4.783 8.877 6.421 13.66 3.66a9.956 9.956 0 0 0 4.197-4.731a9.985 9.985 0 0 0-.537-8.93a9.985 9.985 0 0 0-7.464-4.928A9.956 9.956 0 0 0 7 3.339C2.217 6.101.58 12.217 3.34 17Z"/><path stroke-dasharray="1.5 2" d="M13.196 2.071s-1.098 4.099 1.402 8.43c2.5 4.33 6.599 5.428 6.599 5.428M2.804 8.071s4.098 1.099 6.598 5.43c2.5 4.33 1.402 8.428 1.402 8.428"/></g>`),
		g.Group(children),
	)
}