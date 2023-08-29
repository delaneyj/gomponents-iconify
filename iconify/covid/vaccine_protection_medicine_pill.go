package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineProtectionMedicinePill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.6 16.4a4.6 4.6 0 0 1-6.5-6.5l7.8-7.8a4.6 4.6 0 0 1 6.5 6.5M5.999 5.999l4.25 4.25"/><path d="M23.25 15.75a7.669 7.669 0 0 1-6 7.5a7.669 7.669 0 0 1-6-7.5V13.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5v2.25Zm-6-.75v4.5M15 17.25h4.5"/></g>`),
		g.Group(children),
	)
}