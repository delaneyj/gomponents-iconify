package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomInternal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSZoomInternal0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M44 4H4v40h40V4Z"/><path stroke="#000" stroke-linecap="round" d="M16 4v12H4m32 8v12H24m12 0L24 24"/><path stroke="#fff" stroke-linecap="round" d="M4 6v20M7 4h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSZoomInternal0)"/>`),
		g.Group(children),
	)
}