package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toothbrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M999.46 998.5q-25 24.5-60 24.5t-59-25l-868-863q-12-12-12-29.5t12.5-30t30-12.5t29.5 12l29 29q4-7 7-10l83-83q11-11 27-11t27.5 11.5t11.5 27.5t-12 27l-82 83q-3 3-10 7l12 12q4-7 7-10l83-83q11-11 27-11t27.5 11.5t11.5 27.5t-12 27l-82 83q-3 3-10 6l12 13q4-7 7-10l83-83q11-11 27-11t27.5 11.5t11.5 27.5t-12 27l-82 83q-3 3-10 6l60 60q56 56 106.5 97.5t86.5 64t76.5 47.5t76.5 47.5t86.5 64t106.5 96.5l119 119q25 25 25 60t-25 59.5z"/>`),
		g.Group(children),
	)
}