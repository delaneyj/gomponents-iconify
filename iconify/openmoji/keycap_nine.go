package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.875 12.38h48v47.834h-48z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.125 12.13h48v48h-48z"/><circle cx="36.125" cy="32.092" r="4.659"/><path d="m36.125 44.828l3.816-10.066"/></g>`),
		g.Group(children),
	)
}