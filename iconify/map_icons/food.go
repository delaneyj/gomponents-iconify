package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Food(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M48.894 15.154L44.959 46H31.668l-3.919-31h16.226l3.207-11.077L49 4.471l-3.077 10.66l2.971.023zM25.87 33s.497-4-6.395-4H8.499c-6.882 0-6.395 4-6.395 4H25.87zM2.104 42s-.487 4 6.395 4h10.977c6.892 0 6.395-4 6.395-4H2.104zm22.735-2c1.128 0 2.039-1.114 2.039-2.499c0-1.393-.911-2.501-2.039-2.501H3.04C1.917 35 1 36.108 1 37.501C1 38.886 1.917 40 3.04 40h21.799z"/>`),
		g.Group(children),
	)
}