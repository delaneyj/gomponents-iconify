package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSExpandRight0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" d="M18 6v36"/><path stroke="#fff" stroke-linecap="round" d="M11 6h25M11 42h25"/><path stroke="#000" stroke-linecap="round" d="m32 20l-4 4l4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSExpandRight0)"/>`),
		g.Group(children),
	)
}