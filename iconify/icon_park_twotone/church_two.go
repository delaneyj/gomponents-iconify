package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChurchTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="m13 24l11-10l11 10v20H13V24Z"/><path stroke-linejoin="round" d="m7 8l3-4l3 4v36H7V8Zm28 0l3-4l3 4v36h-6V8Z"/><path d="M24 25v10m-4-6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChurchTwo0)"/>`),
		g.Group(children),
	)
}