package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pligg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M320 351L0 464Q255 279 384 31zm0 288v320Q232 737 33 543zm320 128l320 128q-139 0-189.5 5.5T640 927q-70 19-95 26t-70.5 25.5T384 1023zm357 7L852 515l172-228q-64 241-27 487zM479 0q121 77 237 124.5T960 191l-288 64z"/>`),
		g.Group(children),
	)
}