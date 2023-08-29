package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoordinateSystem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCoordinateSystem0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m24 12l14 8v16l-14 8l-14-8V20l14-8Z"/><path stroke-linecap="round" d="M24 6v6m-14 8l14 8l14-8m0 16l6 3M4 39l6-3m14-8v16"/><path d="m31 16l7 4v8M17 16l-7 4v8m7 12l7 4l7-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCoordinateSystem0)"/>`),
		g.Group(children),
	)
}