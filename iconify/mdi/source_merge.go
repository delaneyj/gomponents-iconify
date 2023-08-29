package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SourceMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3a3 3 0 0 1 3 3c0 1.29-.81 2.39-1.96 2.81c.54 5 5.04 5.96 7.15 6.15A2.985 2.985 0 0 1 18 13a3 3 0 0 1 3 3a3 3 0 0 1-3 3c-1.31 0-2.43-.84-2.84-2c-4.25-.2-5.72-1.81-7.16-3.61v1.78c1.17.41 2 1.52 2 2.83a3 3 0 0 1-3 3a3 3 0 0 1-3-3c0-1.31.83-2.42 2-2.83V8.83A2.99 2.99 0 0 1 4 6a3 3 0 0 1 3-3m0 2a1 1 0 0 0-1 1a1 1 0 0 0 1 1a1 1 0 0 0 1-1a1 1 0 0 0-1-1m0 12a1 1 0 0 0-1 1a1 1 0 0 0 1 1a1 1 0 0 0 1-1a1 1 0 0 0-1-1m11-2a1 1 0 0 0-1 1a1 1 0 0 0 1 1a1 1 0 0 0 1-1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}