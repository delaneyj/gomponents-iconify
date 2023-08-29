package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.875 12.124h48v47.834h-48z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.125 11.874h48v48h-48z"/><circle cx="36.125" cy="39.913" r="4.659"/><path d="m36.125 27.177l-3.816 10.066"/></g>`),
		g.Group(children),
	)
}