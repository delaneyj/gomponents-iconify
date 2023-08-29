package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoordinateSystem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCoordinateSystem0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="m24 12l14 8v16l-14 8l-14-8V20l14-8Z"/><path stroke="#fff" stroke-linecap="round" d="M24 6v6"/><path stroke="#000" stroke-linecap="round" d="m10 20l14 8l14-8"/><path stroke="#fff" stroke-linecap="round" d="m38 36l6 3M4 39l6-3"/><path stroke="#000" stroke-linecap="round" d="M24 28v16"/><path stroke="#fff" d="m31 16l7 4v8M17 16l-7 4v8m7 12l7 4l7-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCoordinateSystem0)"/>`),
		g.Group(children),
	)
}