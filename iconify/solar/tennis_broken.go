package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TennisBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M13.196 2.071s-.232 3.599 2.268 7.93c2.5 4.33 5.732 5.928 5.732 5.928M2.803 8.071s3.233 1.599 5.733 5.93c2.5 4.33 2.268 7.928 2.268 7.928"/><path stroke-linecap="round" d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}