package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineProtectionSyringe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.123 9.03l-4.816-4.816L4.121 12.4a1.362 1.362 0 0 0 0 1.926L6.689 16.9a1.5 1.5 0 0 0 2.124 0v0"/><path d="m4.602 14.808l-2.408 2.408l1.926 1.926l2.408-2.408l-1.926-1.926ZM.75 20.587l2.408-2.408M10.863 2.77l1.444 1.444m1.445 1.445l3.371-3.371m1.926 1.926l-3.371 3.371m0-6.742l4.816 4.816m2.756 9.998a7.67 7.67 0 0 1-6 7.5a7.67 7.67 0 0 1-6-7.5v-2.25a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5v2.25Zm-6-.75v4.5M15 17.157h4.5"/></g>`),
		g.Group(children),
	)
}