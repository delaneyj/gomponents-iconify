package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M12 12.166h48V60H12z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.25 11.916h48v48h-48z"/><path d="M31.419 41.115A4.776 4.776 0 0 0 36.2 44.57h0a4.629 4.629 0 0 0 4.88-4.327a4.629 4.629 0 0 0-4.88-4.327a4.629 4.629 0 0 0 4.88-4.327a4.629 4.629 0 0 0-4.88-4.327h0a4.776 4.776 0 0 0-4.781 3.455"/></g>`),
		g.Group(children),
	)
}