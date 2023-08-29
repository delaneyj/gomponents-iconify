package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaitai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m300.04 183.836l152.664-36.76c-25.558 133.466-12.907 205.28-100.349 301.28l-97.116-32.373L321.1 512C487.426 426.047 471.257 299.332 512 96.17l-189.06 33.97L337 99c11.42-28.234 4.472-57.08-23-70L224 0c38.84 44.126 48.38 94.43 28.804 142.558l-93.564 17.24l-96-32.93l63.51 98.232l106.462-25.484L192 319L2.318 509.67L0 512l225.425-143.738l74.615-184.426z"/>`),
		g.Group(children),
	)
}