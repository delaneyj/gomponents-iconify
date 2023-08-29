package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.875 12.166h48V60h-48z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.125 11.916h48v48h-48z"/><path d="M35.875 44.78h0a4.76 4.76 0 0 1-4.76-4.76v-7.874a4.76 4.76 0 0 1 4.76-4.76h0a4.76 4.76 0 0 1 4.76 4.76v7.874a4.76 4.76 0 0 1-4.76 4.76Z"/></g>`),
		g.Group(children),
	)
}