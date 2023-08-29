package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23.27l-9-7l1.62-1.26l7.37 5.73l7.38-5.739L21 16.27l-9 7ZM12 19l-9-7l1.62-1.26l7.37 5.73l7.38-5.74L21 12l-9 7Zm0-4.27L4.63 9L3 7.73l9-7l9 7L19.36 9L12 14.73Z"/>`),
		g.Group(children),
	)
}