package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.875 12.291h48v47.834h-48z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.125 12.041h48v48h-48z"/><path d="m31.664 30.895l4.8-3.59l.122 17.472"/></g>`),
		g.Group(children),
	)
}