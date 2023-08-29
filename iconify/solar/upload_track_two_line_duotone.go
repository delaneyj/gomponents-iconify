package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTrackTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M13 15V7"/><circle cx="11" cy="15" r="2"/><path stroke-linecap="round" d="M16 10a3 3 0 0 1-3-3"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 22v-7m0 0l2.5 2.5M18 15l-2.5 2.5"/><circle cx="12" cy="12" r="10" opacity=".5"/></g>`),
		g.Group(children),
	)
}