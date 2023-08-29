package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#d0cfce"><path d="M59 41H13V31h46"/><circle cx="36" cy="51" r="5"/><circle cx="36" cy="21" r="5"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M13 31h46v10H13z"/><circle cx="36" cy="51" r="5"/><circle cx="36" cy="21" r="5"/></g>`),
		g.Group(children),
	)
}