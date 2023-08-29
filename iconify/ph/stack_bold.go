package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234.36 170a12 12 0 0 1-4.36 16.37l-96 56a12 12 0 0 1-12.1 0l-96-56a12 12 0 0 1 12.09-20.74l90 52.48L218 165.63a12 12 0 0 1 16.36 4.37ZM218 117.63l-90 52.48l-89.95-52.48A12 12 0 0 0 26 138.37l96 56a12 12 0 0 0 12.1 0l96-56a12 12 0 0 0-12.1-20.74ZM20 80a12 12 0 0 1 6-10.37l96-56a12.06 12.06 0 0 1 12.1 0l96 56a12 12 0 0 1 0 20.74l-96 56a12 12 0 0 1-12.1 0l-96-56A12 12 0 0 1 20 80Zm35.82 0L128 122.11L200.18 80L128 37.89Z"/>`),
		g.Group(children),
	)
}