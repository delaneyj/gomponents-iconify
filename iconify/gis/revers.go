package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Revers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M81.885 31.885c-9.716 0-17.653 7.77-18.043 17.394a3.623 3.623 0 0 0-.074.721a3.623 3.623 0 0 0 .074.72a3.623 3.623 0 0 0 0 .005c.392 9.622 8.329 17.39 18.043 17.39C91.847 68.115 100 59.962 100 50c0-9.962-8.153-18.115-18.115-18.115zm0 7.246A10.814 10.814 0 0 1 92.754 50c0 6.046-4.823 10.87-10.87 10.87A10.816 10.816 0 0 1 71.015 50c0-6.046 4.825-10.87 10.87-10.87zM25.469 24.637V41.34h26.705v17.32H25.469v16.703L0 50Z" color="currentColor"/>`),
		g.Group(children),
	)
}