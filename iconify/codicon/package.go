package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.61 3l5.74 1.53L15 5v6.74l-.37.48l-6.13 1.69l-6.14-1.69l-.36-.48V5l.61-.47L8.34 3h.27zm-.09 1l-4 1l.55.2l3.43.9l3-.81l.95-.29l-3.93-1zM3 11.36l5 1.37V7L3 5.66v5.7zM9 7v5.73l5-1.37V5.63l-2.02.553V8.75l-1 .26V6.457L9 7z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}