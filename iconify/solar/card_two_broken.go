package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M10 16.5H6m2-3H6"/><path d="M14 15c0-.943 0-1.414.293-1.707C14.586 13 15.057 13 16 13c.943 0 1.414 0 1.707.293c.293.293.293.764.293 1.707c0 .943 0 1.414-.293 1.707C17.414 17 16.943 17 16 17c-.943 0-1.414 0-1.707-.293C14 16.414 14 15.943 14 15Z"/><path stroke-linecap="round" d="M22 12c0-3.771 0-5.657-1.172-6.828C19.657 4 17.771 4 14 4h-4C6.229 4 4.343 4 3.172 5.172C2 6.343 2 8.229 2 12c0 3.771 0 5.657 1.172 6.828C4.343 20 6.229 20 10 20h4c3.771 0 5.657 0 6.828-1.172c.654-.653.943-1.528 1.07-2.828M2 10h5m15 0H11"/></g>`),
		g.Group(children),
	)
}