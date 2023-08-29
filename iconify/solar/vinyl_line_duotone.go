package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="3"/><path stroke-linecap="round" d="M21.95 13c-.501 5.054-4.765 9-9.95 9c-5.523 0-10-4.477-10-10c0-5.185 3.947-9.448 9-9.95" opacity=".5"/><path stroke-linecap="round" d="M15 2.457a10.024 10.024 0 0 1 6.542 6.542M15 12V2.456"/></g>`),
		g.Group(children),
	)
}