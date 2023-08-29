package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylRecordLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="3" opacity=".5"/><path stroke-linecap="round" d="M4.929 19.071c3.905 3.905 10.237 3.905 14.142 0c3.905-3.905 3.905-10.237 0-14.142c-3.905-3.905-10.237-3.905-14.142 0c-3.905 3.905-3.905 10.237 0 14.142Z" opacity=".5"/><path stroke-linecap="round" d="M7.404 16.597a6.5 6.5 0 0 1 0-9.193m9.192 0a6.5 6.5 0 0 1 0 9.193"/></g>`),
		g.Group(children),
	)
}