package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 14c0-3.771 0-5.657 1.172-6.828C4.343 6 6.229 6 10 6h4c3.771 0 5.657 0 6.828 1.172C22 8.343 22 10.229 22 14c0 3.771 0 5.657-1.172 6.828C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 17.771 2 14Z"/><circle cx="8" cy="14" r="3" opacity=".5"/><path stroke-linecap="round" d="M13.5 11H19"/><path stroke-linecap="round" d="M13.5 14H19" opacity=".5"/><path stroke-linecap="round" d="M13.5 17H19M6.5 6L15 2"/></g>`),
		g.Group(children),
	)
}