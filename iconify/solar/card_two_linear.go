package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 12c0-3.771 0-5.657 1.172-6.828C4.343 4 6.229 4 10 4h4c3.771 0 5.657 0 6.828 1.172C22 6.343 22 8.229 22 12c0 3.771 0 5.657-1.172 6.828C19.657 20 17.771 20 14 20h-4c-3.771 0-5.657 0-6.828-1.172C2 17.657 2 15.771 2 12Z"/><path stroke-linecap="round" d="M10 16.5H6m2-3H6M2 10h20"/><path d="M14 15c0-.943 0-1.414.293-1.707C14.586 13 15.057 13 16 13c.943 0 1.414 0 1.707.293c.293.293.293.764.293 1.707c0 .943 0 1.414-.293 1.707C17.414 17 16.943 17 16 17c-.943 0-1.414 0-1.707-.293C14 16.414 14 15.943 14 15Z"/></g>`),
		g.Group(children),
	)
}