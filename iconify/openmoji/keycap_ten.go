package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.875 12.166h48V60h-48z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.125 11.916h48v48h-48z"/><path d="M39.822 44.442h0a4.643 4.643 0 0 1-4.643-4.643v-7.681a4.643 4.643 0 0 1 4.643-4.643h0a4.643 4.643 0 0 1 4.643 4.643v7.681a4.643 4.643 0 0 1-4.643 4.643ZM25.535 31.05l4.644-3.48v16.776"/></g>`),
		g.Group(children),
	)
}