package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Id(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><circle cx="8" cy="7.5" r="2.25"/><path d="M4.75 12.75c0-1 .75-3 3.25-3s3.25 2 3.25 3"/></g>`),
		g.Group(children),
	)
}