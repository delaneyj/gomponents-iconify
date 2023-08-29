package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 6h3.9A5.006 5.006 0 0 0 10 2.1V6Zm0-4.917A6.005 6.005 0 0 1 15 7H9V1c.34 0 .675.028 1 .083ZM7 8l1 1h4.9A5.002 5.002 0 0 1 3 8a5.002 5.002 0 0 1 4-4.9V8Zm1 6a6.002 6.002 0 0 0 6-6H8V2a6.002 6.002 0 0 0 0 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}