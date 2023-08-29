package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossMarkButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#B1CC33" d="M59.04 60.166H12.96a.96.96 0 0 1-.96-.96v-46.08c0-.53.43-.96.96-.96h46.08c.53 0 .96.43.96.96v46.08c0 .53-.43.96-.96.96z"/><path fill="#D0CFCE" d="m42.837 35.575l2.481-2.521l8.661-8.811l-6.68-6.796l-8.666 8.816l-2.358 2.397l-.305.306l-11.261-11.457l-6.68 6.796l11.066 11.258l.193.195l-2.649 2.691l-9.407 9.57l6.68 6.796l9.409-9.572l2.419-2.458l.226-.232l12.122 12.332l6.68-6.796l-12.117-12.327z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M59.04 60.166H12.96a.96.96 0 0 1-.96-.96v-46.08c0-.53.43-.96.96-.96h46.08c.53 0 .96.43.96.96v46.08c0 .53-.43.96-.96.96z"/><path d="m42.837 35.575l2.481-2.521l8.661-8.811l-6.68-6.796l-8.666 8.816l-2.358 2.397l-.305.306l-11.261-11.457l-6.68 6.796l11.066 11.258l.193.195l-2.649 2.691l-9.407 9.57l6.68 6.796l9.409-9.572l2.419-2.458l.226-.232l12.122 12.332l6.68-6.796l-12.117-12.327z"/></g>`),
		g.Group(children),
	)
}