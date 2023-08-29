package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Permissions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPermissions0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M20 10H6a2 2 0 0 0-2 2v26a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2v-2.5"/><path d="M10 23h8m-8 8h24"/><circle cx="34" cy="16" r="6" fill="#555" stroke-linejoin="round"/><path stroke-linejoin="round" d="M44 28.419C42.047 24.602 38 22 34 22s-5.993 1.133-8.05 3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPermissions0)"/>`),
		g.Group(children),
	)
}