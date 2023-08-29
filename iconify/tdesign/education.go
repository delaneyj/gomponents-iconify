package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Education(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 .807L23.835 8.5L22 9.693V16h-2v-5.008l-1 .65V17.5c0 1.47-1.014 2.615-2.253 3.338C15.482 21.576 13.802 22 12 22s-3.483-.424-4.747-1.162C6.013 20.116 5 18.97 5 17.5v-5.857L.165 8.5L12 .807ZM7 12.943V17.5c0 .463.33 1.067 1.261 1.61c.908.53 2.227.89 3.739.89s2.831-.36 3.739-.89c.932-.543 1.26-1.147 1.26-1.61v-4.557l-5 3.25l-5-3.25ZM20.165 8.5L12 3.192L3.835 8.5L12 13.807L20.165 8.5Z"/>`),
		g.Group(children),
	)
}