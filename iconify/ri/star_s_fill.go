package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarSFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17l-5.877 3.59l1.598-6.7l-5.23-4.48l6.865-.55L12 2.5l2.645 6.36l6.865.55l-5.23 4.48l1.598 6.7L12 17Z"/>`),
		g.Group(children),
	)
}