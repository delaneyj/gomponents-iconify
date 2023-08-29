package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GppBadRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 13.4l1.4 1.4q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7L13.4 12l1.4-1.4q.275-.275.275-.7t-.275-.7q-.275-.275-.7-.275t-.7.275L12 10.6l-1.4-1.4q-.275-.275-.7-.275t-.7.275q-.275.275-.275.7t.275.7l1.4 1.4l-1.4 1.4q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275l1.4-1.4Zm0 8.5q-.175 0-.325-.025t-.3-.075Q8 20.675 6 17.637T4 11.1V6.375q0-.625.363-1.125t.937-.725l6-2.25q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 3.5-2 6.538T12.625 21.8q-.15.05-.3.075T12 21.9Z"/>`),
		g.Group(children),
	)
}