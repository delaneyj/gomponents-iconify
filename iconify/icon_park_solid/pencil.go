package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPencil0"><g fill="none"><g stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-path="url(#ipSPencil1)"><path d="m31 8.999l8 8m-31 15L36 4l8 7.999l-28 28l-10 2l2-10Zm23-23l8 8m-30 15l7 7m-3-4l22-22"/></g><defs><clipPath id="ipSPencil1"><path fill="#000" d="M0 0h48v48H0z"/></clipPath></defs></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPencil0)"/>`),
		g.Group(children),
	)
}