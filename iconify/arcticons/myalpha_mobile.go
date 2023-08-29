package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyalphaMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.141 32.191L27.45 22.5h11.974a15.397 15.397 0 0 0-2.47-7H22.5V8.576a15.397 15.397 0 0 0-7 2.47V15.5h-4.454a15.397 15.397 0 0 0-2.47 7H15.5v14.454a15.397 15.397 0 0 0 7 2.47V27.45l9.691 9.691a15.553 15.553 0 0 0 4.95-4.95Z"/>`),
		g.Group(children),
	)
}