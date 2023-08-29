package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M12 12.166h48V60H12z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.25 11.916h48v48h-48z"/><path d="M31.03 31.498a5.33 5.33 0 0 1 5.22-4.255h0a5.312 5.312 0 0 1 3.768 1.561a4.115 4.115 0 0 1-.046 5.58l-9.05 10.205h10.656"/></g>`),
		g.Group(children),
	)
}